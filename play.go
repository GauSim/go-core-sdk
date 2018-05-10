package main

import (
	"flag"
	"fmt"

	"github.com/gausim/gosdk/core"
	"github.com/gausim/gosdk/dto"
)

var account string
var password string
var username string
var clientSecret string
var clientIdentifier string

func validateFlags() {

	if len(account) == 0 {
		panic("missing flag '-account <account>'")
	}

	if len(username) == 0 {
		panic("missing flag '-username <username>'")
	}

	if len(clientSecret) == 0 {
		panic("missing flag '-clientSecret <clientSecret>'")
	}

	if len(clientIdentifier) == 0 {
		panic("missing flag '-clientIdentifier <clientIdentifier>'")
	}
}

func main() {

	flag.StringVar(&account, "account", "", "account")
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&password, "password", "", "password")

	flag.StringVar(&clientSecret, "clientSecret", "", "clientSecret")
	flag.StringVar(&clientIdentifier, "clientIdentifier", "", "clientIdentifier")

	flag.Parse()

	validateFlags()

	client := core.ClientConfig{
		Identifier:    clientIdentifier,
		Version:       "0.0.1-beta",
		Secret:        clientSecret,
		OAuthEndpoint: "https://et.dev.coresuite.com/api/oauth2/v1",
		Debug:         true}

	creds, err := client.GetCredentials(account, username, password)
	if err != nil {
		panic(err)
	}

	q := core.Query{SQL: `
		SELECT 
			equipment.id  
		FROM 
			Equipment equipment  
		WHERE 
			equipment.parentId is null 
			AND ( equipment.inactive IN (false) ) 
		ORDER BY 
			equipment.name ASC, equipment.serialNumber ASC`,
		DTOs: []dto.Entity{dto.Equipment}}

	type Equipment struct {
		ID string `json:"id"`
	}

	type Item struct {
		Equipment Equipment `json:"equipment"`
	}

	list := []Item{}

	client.Query(creds, q, &list)
	for key, value := range list {
		fmt.Println("idx:", key, " val: ", value)
	}

}
