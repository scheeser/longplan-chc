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
	// Get tenant mnemonic used to build the URL.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Tenant Mnemonic: ")
	tenant, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error reading the tenant mnemonic: %s\n", err)
		return
	}
	tenant = strings.Replace(tenant, "\n", "", -1)
	baseURL := fmt.Sprintf("https://%s.api.us.healtheintent.com", tenant)

	// The auth header could be either a bearer token (which is much easier) or generated token.
	fmt.Print("Enter Authorization Header: ")
	authHeader, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("There was an error reading the authorization header value: %s\n", err)
		return
	}
	authHeader = strings.Replace(authHeader, "\n", "", -1)

	httpClient := http.Client{
		//TODO: Need to set resonable defaults.
	}

	// Create the Health Concern Definition struct.
	healthConcernDefn := models.HealthConcernDefinition{
		Text:   "CHC 2019: Diabetes Mellitus Type 1",
		Status: "ACTIVE",
		Concept: models.Concept{
			ContextID: "53EF3068AE8F4EDE9951DC170CBBE6DA",
			Alias:     "DIABETES_MELLITUS_TYPE_1_CLIN",
		},
	}

	createdHealthConcernDefn, err := client.CreateHealthConcernDefinition(healthConcernDefn, httpClient, baseURL, authHeader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Created Health Concern Definition '%s' wit ID: %s\n", createdHealthConcernDefn.Text, createdHealthConcernDefn.ID)

	// Create goals. The intial set in the repo is based on https://www.heart.org/en/health-topics/diabetes/prevention--treatment-of-diabetes/healthy-living-tips-for-people-with-diabetes
	goalDefinitions, err := models.ReadDefinitionsFromCSV("goal_definitions.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Using the Goal Definitions parsed from the CSV POST each one to the API.
	var createdGoalDefns []models.GoalDefinition
	for _, goalDefn := range goalDefinitions {
		createdGoalDefn, err := client.CreateGoalDefinition(goalDefn, httpClient, baseURL, authHeader)
		if err != nil {
			fmt.Println(err)
			return
		}

		createdGoalDefns = append(createdGoalDefns, createdGoalDefn)

		fmt.Printf("Created Goal Definition '%s' with ID: %s\n", createdGoalDefn.Text, createdHealthConcernDefn.ID)
	}

	// Create the Template using the Health Concern and Goals that were just created.
	var goalTemplates []models.GoalDefinitionTemplate
	for _, goalDefn := range createdGoalDefns {
		goalTemplates = append(goalTemplates, models.GoalDefinitionTemplate{
			ID: goalDefn.ID,
		})
	}

	planTemplate := models.PlanTemplate{
		Description: "CHC 2019: Diabetes Management",
		Status:      "ACTIVE",
		HealthConcernDefinitions: []models.HealthConcernDefinitionTemplate{
			models.HealthConcernDefinitionTemplate{
				ID:              createdHealthConcernDefn.ID,
				GoalDefinitions: goalTemplates,
			},
		},
	}

	createdPlanTemplate, err := client.CreatePlanTemplate(planTemplate, httpClient, baseURL, authHeader)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print out the ID of the template that was created.
	fmt.Printf("Created Plan Template '%s' with ID: %s\n", createdPlanTemplate.Description, createdPlanTemplate.ID)
}
