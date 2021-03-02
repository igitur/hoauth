package config

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/igitur/hoauth/pkg/db"
)

func ChooseClient(database *db.CredentialStore) (string, error) {
	allClients, err := database.GetClients()

	if err != nil {
		panic(err)
	}

	if len(allClients) == 0 {
		log.Fatalf("Please create a connection using `%s`",
			color.Green.Sprintf("hoauth setup [connectionName]"))
	}

	var connections []string

	for _, value := range allClients {
		connections = append(connections, value.Alias)
	}

	connectionPicker := &survey.Select{
		Message: "Choose a client",
		Options: connections,
	}

	var chosenConnection string

	askErr := survey.AskOne(connectionPicker, &chosenConnection, survey.WithValidator(survey.Required))

	if askErr != nil {
		return "", askErr
	}

	return chosenConnection, nil
}
