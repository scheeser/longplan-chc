package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.cerner.com/JS016083/longplan-chc/models"
)

// CreateHealthConcern Creates a health concern for a consumer.
func CreateHealthConcern(consumerID string, healthConcern models.HealthConcern, client http.Client, baseURL string, authHeader string) (createdHealthConcern models.HealthConcern, funcErr error) {
	URL := fmt.Sprintf("%s/health-concern/v1/consumers/%s/health-concerns", baseURL, consumerID)

	body, err := postJSON(healthConcern, client, URL, authHeader)
	if err != nil {
		return createdHealthConcern, err
	}

	err = json.Unmarshal(body, &createdHealthConcern)
	if err != nil {
		funcErr = fmt.Errorf("error marshaling health concern POST response body: %s", err.Error())
		return
	}

	return createdHealthConcern, nil
}

// GetHealthConcernDefinition Get a health concern definition by id.
func GetHealthConcernDefinition(definitionID string, client http.Client, baseURL string, authHeader string) (retrievedDefinition models.HealthConcernDefinition, funcErr error) {
	URL := fmt.Sprintf("%s/health-concern/v1/health-concern-definitions/%s", baseURL, definitionID)

	body, err := getJSON(client, URL, authHeader)
	if err != nil {
		return retrievedDefinition, err
	}

	err = json.Unmarshal(body, &retrievedDefinition)
	if err != nil {
		funcErr = fmt.Errorf("issue marshaling Health Concern Definition GET response body: %s", err.Error())
		return
	}

	return retrievedDefinition, nil
}

// CreateHealthConcernDefinition Creates a health concern definiton.
func CreateHealthConcernDefinition(healthConcernDefn models.HealthConcernDefinition, client http.Client, baseURL string, authHeader string) (createdHealthConcernDefn models.HealthConcernDefinition, funcErr error) {
	URL := fmt.Sprintf("%s/health-concern/v1/health-concern-definitions", baseURL)

	body, err := postJSON(healthConcernDefn, client, URL, authHeader)
	if err != nil {
		return createdHealthConcernDefn, err
	}

	err = json.Unmarshal(body, &createdHealthConcernDefn)
	if err != nil {
		funcErr = fmt.Errorf("error marshaling health concern definition POST response body: %s", err.Error())
		return
	}

	return createdHealthConcernDefn, nil
}
