// Go contains rich function for grab web contents. _net/http_ is the major
// library.
// Ref: [golang.org](http://golang.org/pkg/net/http/#pkg-examples).
package main

import (
	"flag"
	"fmt"

	"github.com/gausim/hello/core"
	"github.com/gausim/hello/dto"
)

var account string
var password string
var username string
var secret string
var identifier string

func main() {

	flag.StringVar(&account, "account", "", "account")
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&secret, "secret", "", "secret")
	flag.StringVar(&identifier, "identifier", "", "identifier")
	flag.Parse()

	client := core.ClientConfig{
		Identifier:    identifier,
		Version:       "0.0.1-beta",
		Secret:        secret,
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
