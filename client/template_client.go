package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.cerner.com/JS016083/longplan-chc/models"
)

// GetPlanTemplate Get a template by id.
func GetPlanTemplate(templateID string, client http.Client, baseURL string, authHeader string) (retrievedPlanTemplate models.PlanTemplate, funcErr error) {
	URL := fmt.Sprintf("%s/longitudinal-plan/v1/plan-templates/%s", baseURL, templateID)

	body, err := getJSON(client, URL, authHeader)
	if err != nil {
		return retrievedPlanTemplate, err
	}

	err = json.Unmarshal(body, &retrievedPlanTemplate)
	if err != nil {
		funcErr = fmt.Errorf("issue marshaling plan template GET response body: %s", err.Error())
		return
	}

	return retrievedPlanTemplate, nil
}

// CreatePlanTemplate Creates the provided template.
func CreatePlanTemplate(planTemplate models.PlanTemplate, client http.Client, baseURL string, authHeader string) (createdPlanTemplate models.PlanTemplate, funcErr error) {
	URL := fmt.Sprintf("%s/longitudinal-plan/v1/plan-templates", baseURL)

	body, err := postJSON(planTemplate, client, URL, authHeader)
	if err != nil {
		return createdPlanTemplate, err
	}

	err = json.Unmarshal(body, &createdPlanTemplate)
	if err != nil {
		funcErr = fmt.Errorf("issue marshaling plan template POST response body: %s", err.Error())
		return
	}

	return createdPlanTemplate, nil
}
