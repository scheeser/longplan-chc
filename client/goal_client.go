package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.cerner.com/JS016083/longplan-chc/models"
)

// CreateGoal Creates a Goal for a consumer.
func CreateGoal(consumerID string, goal models.Goal, client http.Client, baseURL string, authHeader string) (createdGoal models.Goal, funcErr error) {
	URL := fmt.Sprintf("%s/longitudinal-plan/v1/consumers/%s/goals", baseURL, consumerID)

	body, err := postJSON(goal, client, URL, authHeader)
	if err != nil {
		return createdGoal, err
	}

	err = json.Unmarshal(body, &createdGoal)
	if err != nil {
		funcErr = fmt.Errorf("error marshaling goal POST response body: %s", err.Error())
		return
	}

	return createdGoal, nil
}

// RelateHealthConcernToGoal Relate the provided goal to the health concern.
func RelateHealthConcernToGoal(consumerID string, goalID string, healthConcernID string, client http.Client, baseURL string, authHeader string) (funcErr error) {
	URL := fmt.Sprintf("%s/longitudinal-plan/v1/consumers/%s/goals/%s/related-health-concerns/%s", baseURL, consumerID, goalID, healthConcernID)

	_, err := putJSON(nil, client, URL, authHeader)
	if err != nil {
		return err
	}

	return nil
}

// GetGoalDefinition Get a Goal definition by id.
func GetGoalDefinition(definitionID string, client http.Client, baseURL string, authHeader string) (retrievedDefinition models.GoalDefinition, funcErr error) {
	URL := fmt.Sprintf("%s/longitudinal-plan/v1/goal-definitions/%s", baseURL, definitionID)

	body, err := getJSON(client, URL, authHeader)
	if err != nil {
		return retrievedDefinition, err
	}

	err = json.Unmarshal(body, &retrievedDefinition)
	if err != nil {
		funcErr = fmt.Errorf("issue marshaling Goal Definition GET response body: %s", err.Error())
		return
	}

	return retrievedDefinition, nil
}

// CreateGoalDefinition Creates a goal definition.
func CreateGoalDefinition(goalDefn models.GoalDefinition, client http.Client, baseURL string, authHeader string) (createdGoalDefn models.GoalDefinition, funcErr error) {
	URL := fmt.Sprintf("%s/longitudinal-plan/v1/goal-definitions", baseURL)

	body, err := postJSON(goalDefn, client, URL, authHeader)
	if err != nil {
		return createdGoalDefn, err
	}

	err = json.Unmarshal(body, &createdGoalDefn)
	if err != nil {
		funcErr = fmt.Errorf("issue marshaling goal definition POST response body: %s", err.Error())
		return
	}

	return createdGoalDefn, nil
}
