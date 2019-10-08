package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.cerner.com/JS016083/longplan-chc/client"
	"github.cerner.com/JS016083/longplan-chc/models"
)

func main() {
	// Get tenant mnemonic used to build the required URL.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Tenant Mnemonic: ")
	tenant, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error reading the tenant mnemonic: %s\n", err)
		return
	}
	tenant = strings.Replace(tenant, "\n", "", -1)
	baseURL := fmt.Sprintf("https://%s.api.us.healtheintent.com", tenant)

	// The authorization header could be either a bearer token (easier) or generated token.
	fmt.Print("Enter Authorization Header: ")
	authHeader, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error reading the authorization header value: %s\n", err)
		return
	}
	authHeader = strings.Replace(authHeader, "\n", "", -1)

	// To align with how millennium/one plan resolves consumers with millennium person ids, we search
	// with the millennium person id and the data partition id for millennium within the consumer source identifier.
	fmt.Print("Enter millennium person id: ")
	sourceID, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error reading the source id: %s\n", err)
		return
	}
	sourceID = strings.Replace(sourceID, "\n", "", -1)

	fmt.Print("Enter data partition id for millennium: ")
	dataPartitionID, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error reading the data partition id: %s\n", err)
		return
	}
	dataPartitionID = strings.Replace(dataPartitionID, "\n", "", -1)

	// We need a template id to use as the base for creating all of the plan items.
	fmt.Print("Enter Template Id: ")
	templateID, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error reading the template id: %s\n", err)
	}
	templateID = strings.Replace(templateID, "\n", "", -1)

	httpClient := http.Client{
		//TODO: Need to set resonable defaults.
	}

	// Find or create the consumer based of the provided source identifier.
	consumer, err := getConsumer(sourceID, dataPartitionID, httpClient, baseURL, authHeader)
	if err != nil {
		fmt.Println(err)
		return
	}
	consumerID := consumer.ID

	// Add the long plan items defined in the template to the given consumer's plan.
	err = addToLongPlan(consumerID, templateID, httpClient, baseURL, authHeader)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getConsumer(sourceID string, dataPartitionID string, httpClient http.Client, baseURL string, authHeader string) (funcConsumer models.ConsumerEntity, funcError error) {
	// Using the source identifier, try and find an existing consumer.
	consumerList, err := client.GetConsumersBySourceIdentifier(sourceID, dataPartitionID, httpClient, baseURL, authHeader)
	if err != nil {
		funcError = fmt.Errorf("there was an error getting the list of consumers: %s", err)
		return
	}

	// If we have found consumers, ensure there aren't two or more
	if consumerList.TotalResults > 1 {
		funcError = fmt.Errorf("%d consumers exist with the  source id %s and data partitionid %s", consumerList.TotalResults, sourceID, dataPartitionID)
		return
	}

	// If we have only one consuemr, return it.
	if consumerList.TotalResults == 1 {
		funcConsumer = consumerList.Items[0]
		fmt.Printf("Found consumer %s tied to source id %s data partition id %s.\n", funcConsumer.ID, sourceID, dataPartitionID)
		return
	}

	// If we haven't found a consumer, create it. We'll need a few more details. In a more realistic example we should have the patient details.
	fmt.Printf("No consumer found with the source id %s data partition id %s. A new consumer must be created.\n", sourceID, dataPartitionID)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the person's Given name: ")
	givenName, err := reader.ReadString('\n')

	if err != nil {
		funcError = fmt.Errorf("there was an error reading the given name: %s", err)
		return
	}
	givenName = strings.Replace(givenName, "\n", "", -1)

	fmt.Print("Enter the person's Family name: ")
	familyName, err := reader.ReadString('\n')

	if err != nil {
		funcError = fmt.Errorf("there was an error reading the family name: %s", err)
		return
	}
	familyName = strings.Replace(familyName, "\n", "", -1)

	consumer := &models.ConsumerEntity{
		Name: models.ConsumerName{
			Given:     givenName,
			Family:    familyName,
			Formatted: fmt.Sprintf("%s %s", givenName, familyName),
		},
		SourceIdentifiers: []models.SourceIdentifier{
			models.SourceIdentifier{
				ID:              sourceID,
				DataPartitionID: dataPartitionID,
			},
		},
	}

	createdConsumer, err := client.CreateConsumer(*consumer, httpClient, baseURL, authHeader)
	if err != nil {
		funcError = fmt.Errorf("There was an error creating the consumer: %s", err)
		return
	}
	fmt.Printf("Created consumer with id %s\n", createdConsumer.ID)
	return createdConsumer, nil
}

func addToLongPlan(consumerID string, templateID string, httpClient http.Client, baseURL string, authHeader string) error {
	// Get the plan template using the provided information.
	template, err := client.GetPlanTemplate(templateID, httpClient, baseURL, authHeader)
	if err != nil {
		return fmt.Errorf("there was an error retrieving the plan template: %s", err)
	}

	// Get any Health Concerns referneced by the template.
	for _, templateHealthConcern := range template.HealthConcernDefinitions {
		healthConcernDefn, err := client.GetHealthConcernDefinition(templateHealthConcern.ID, httpClient, baseURL, authHeader)
		if err != nil {
			return fmt.Errorf("there was an error retrieving the Health Concern Definition with id %s: %s", templateHealthConcern.ID, err)
		}

		// Use the Health Concern Definition id to create an instance of the defined Health Concern for the consumer.
		healthConcern := models.HealthConcern{
			Onset:  "2019",
			Status: "ACTIVE",
			Code: models.CodeableConcept{
				Text: healthConcernDefn.Text,
				Codings: []models.Coding{
					models.Coding{
						// TODO: It would be ideal if we pulled a real code from the concept for the coidng rather than using the concept itself.
						System: healthConcernDefn.Concept.ContextID,
						Code:   healthConcernDefn.Concept.Alias,
					},
				},
			},
			CreatedBy: models.CreatedBy{
				Type: "SUBJECT", // Subject is used here simply because it's the easiest structure to achieve for the sake of simplifying the example.
			},
			HealthConcernDefinition: models.Reference{
				ID: templateHealthConcern.ID,
			},
		}

		createdHealthConcern, err := client.CreateHealthConcern(consumerID, healthConcern, httpClient, baseURL, authHeader)
		if err != nil {
			return fmt.Errorf("there was an error creating the Health Concern from definition %s: %s", templateHealthConcern.ID, err)
		}

		fmt.Printf("Created Health Concern: %s\n", createdHealthConcern.ID)

		// Iterate over every goal in the template and create one for the consumer. In this example we only deal with goals attached directly
		// to the Health Concern. There are more nested options defined in the template that are ignored (Short-Term Goals, Activities, etc.).
		for _, templateGoal := range templateHealthConcern.GoalDefinitions {
			goalDefn, err := client.GetGoalDefinition(templateGoal.ID, httpClient, baseURL, authHeader)
			if err != nil {
				return fmt.Errorf("there was an error retrieving the Goal Definition with id %s: %s", templateGoal.ID, err)
			}

			// Use the Goal Definition id to create an instance of the Goal for the consumer.
			goal := models.Goal{
				Progress:    "NOT_ACHIEVED",
				Description: goalDefn.Text,
				Status:      "PROPOSED",
				CreatedBy: models.CreatedBy{
					Type:    "SYSTEM",
					Display: "CHC Demo",
				},
				GoalDefinition: models.Reference{
					ID: goalDefn.ID,
				},
			}

			createdGoal, err := client.CreateGoal(consumerID, goal, httpClient, baseURL, authHeader)
			if err != nil {
				return fmt.Errorf("there was an error creating the Goal from definition %s: %s", goalDefn.ID, err)
			}

			fmt.Printf("Created Goal: %s\n", createdGoal.ID)

			// Relate the goal with the parent Health Concern
			err = client.RelateHealthConcernToGoal(consumerID, createdGoal.ID, createdHealthConcern.ID, httpClient, baseURL, authHeader)
			if err != nil {
				return fmt.Errorf("there was an error associating Goal %s to Health Concern %s: %s", createdGoal.ID, createdHealthConcern.ID, err)
			}

			fmt.Printf("Associated Goal %s to Health Concern %s.\n", createdGoal.ID, createdHealthConcern.ID)
		}
	}

	return nil
}
