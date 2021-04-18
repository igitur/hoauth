package clientCredsFlow

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/igitur/hoauth/pkg/db"
	"github.com/igitur/hoauth/pkg/oidc"
)

type ClientCredsFlowInteractor struct {
	endPoints       oidc.EndPoints
	database        *db.CredentialStore
	operatingSystem string
}

func NewClientCredsFlow(endPoints oidc.EndPoints, database *db.CredentialStore, operatingSystem string) ClientCredsFlowInteractor {
	return ClientCredsFlowInteractor{
		endPoints:       endPoints,
		database:        database,
		operatingSystem: operatingSystem,
	}
}

func (interactor *ClientCredsFlowInteractor) Request(client db.OidcClient, dryRun bool) {
	var scopes = strings.Join(client.Scopes, " ")

	var tokenResult, tokenErr = oidc.RequestWithClientCredentials(interactor.endPoints.TokenEndPoint, client.ClientId, client.ClientSecret, scopes)

	if tokenErr != nil {
		log.Fatalln(tokenErr)
	}

	log.Println("Validating access token")

	//var _, validateErr = oidc.ValidateToken(tokenResult.AccessToken, interactor.endPoints, client.ClientId)

	// if validateErr != nil {
	// 	log.Fatalln(validateErr)
	// }

	jsonData, jsonErr := json.MarshalIndent(tokenResult, "", "    ")

	log.Print("Storing tokens in local keychain")
	_, tokenSaveErr := interactor.database.SaveTokens(client.Alias, oidc.TokenResultSet{
		AccessToken: tokenResult.AccessToken,
		ExpiresAt:   tokenResult.ExpiresAt,
	})

	// Can fail with warning
	if tokenSaveErr != nil {
		log.Printf("%s: %v",
			color.Yellow.Sprintf("failed to save tokens to keychain"),
			tokenSaveErr,
		)
	}

	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}
	_, finalWriteErr := fmt.Fprintln(os.Stdout, string(jsonData))

	if finalWriteErr != nil {
		log.Fatalln(finalWriteErr)
	}
}
