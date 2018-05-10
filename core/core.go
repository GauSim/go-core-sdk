package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gausim/gosdk/dto"
)

type CoreError struct {
	message string
}

func NewCoreError(message string) *CoreError {
	return &CoreError{
		message: message,
	}
}
func (e *CoreError) Error() string {
	return e.message
}

type ClientConfig struct {
	Identifier    string
	Version       string
	Secret        string
	OAuthEndpoint string
	Debug         bool
}

type Company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Companies struct {
	Companies []Company `json:"companies"`
}

type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	ClusterURL  string `json:"cluster_url"`
	AccountID   int    `json:"account_id"`
	Account     string `json:"account"`
	User        string
}

type Credentials struct {
	Auth
	SelectedCompany Company
}

type Query struct {
	SQL  string `json:"query"`
	DTOs []dto.Entity
}

type envelope struct {
	Data interface{} `json:"data"`
}

// GetAuth gets the token
func (r *ClientConfig) GetAuth(account, username, password string) (Auth, error) {
	auth := Auth{User: username}
	loginInfo := url.Values{}
	loginInfo.Set("grant_type", "password")
	loginInfo.Set("username", account)
	loginInfo.Set("password", password)

	b := bytes.NewBufferString(loginInfo.Encode())
	req, err := http.NewRequest("POST", r.OAuthEndpoint+"/token", b)
	if err != nil {
		log.Fatalf("http.NewRequest() error: %v\n", err)
		return auth, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(r.Identifier, r.Secret)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("http.Do() error: %v\n", err)
		return auth, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() error: %v\n", err)
		return auth, err
	}

	json.Unmarshal(data, &auth)

	if len(auth.AccessToken) == 0 {
		return auth, NewCoreError("authentication failed [" + fmt.Sprintf("%s", data) + "]")
	}

	return auth, nil
}

func (r *ClientConfig) GetCredentials(account, username, password string) (Credentials, error) {
	auth, err := r.GetAuth(account, username, password)
	if err != nil {
		return Credentials{}, err
	}

	companies, err := r.GetCompanies(auth)
	if err != nil {
		return Credentials{}, err
	}
	creds := Credentials{Auth: auth, SelectedCompany: companies[0]}

	return creds, nil
}

// GetCompanies todo
func (r ClientConfig) GetCompanies(auth Auth) ([]Company, error) {
	u, _ := url.Parse(auth.ClusterURL)
	u.Scheme = "https"
	u.Path = "/masterCloudInterface/informationService/companies"

	q := u.Query()
	q.Set("includeSubscriptions", "true")
	q.Set("account", auth.Account)
	q.Set("user", auth.User)

	u.RawQuery = q.Encode()
	url := u.String()

	cr := ClientRequest{method: "GET", url: url, auth: auth}
	result := Companies{}
	err := r.Request(cr, &result)

	if len(result.Companies) == 0 {
		return nil, NewCoreError("No Companies found")
	}

	return result.Companies, err
}

func (r ClientConfig) Query(creds Credentials, data Query, result interface{}) {

	u, _ := url.Parse(creds.ClusterURL)
	u.Scheme = "https"
	u.Path = "/api/data/query/v1"

	q := u.Query()

	q.Set("account", creds.Account)
	q.Set("user", creds.User)
	q.Set("company", creds.SelectedCompany.Name)

	var buffer bytes.Buffer
	for _, value := range data.DTOs {
		buffer.WriteString(fmt.Sprintf("%s.%v", value.Name, value.Version))
	}

	q.Set("dtos", buffer.String())

	q.Set("useUdfAs", "unknownAsText")

	u.RawQuery = q.Encode()

	url := u.String()

	cr := ClientRequest{
		method: "POST",
		url:    url,
		auth:   creds.Auth,
		data:   Query{SQL: data.SQL}} // only send SQL

	err := r.Request(cr, &result)
	if err != nil {
		log.Fatalf("query() error: %v\n", err)
	}

}
