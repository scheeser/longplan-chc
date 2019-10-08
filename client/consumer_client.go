package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.cerner.com/JS016083/longplan-chc/models"
)

// GetConsumersBySourceIdentifier Get consumers for a tenant by source identifier.
func GetConsumersBySourceIdentifier(sourceID string, dataPartitionID string, client http.Client, baseURL string, authHeader string) (consumerList models.ConsumerList, funcErr error) {
	URL := fmt.Sprintf("%s/consumer/v1/consumers?sourceId=%s&dataPartitionId=%s", baseURL, sourceID, dataPartitionID)

	body, err := getJSON(client, URL, authHeader)
	if err != nil {
		return consumerList, err
	}

	err = json.Unmarshal(body, &consumerList)
	if err != nil {
		funcErr = fmt.Errorf("issue marshaling consumer list GET response body: %s", err.Error())
		return
	}

	return consumerList, nil
}

// CreateConsumer Creates the provided consumer.
func CreateConsumer(consumer models.ConsumerEntity, client http.Client, baseURL string, authHeader string) (createdConsumer models.ConsumerEntity, funcErr error) {
	URL := fmt.Sprintf("%s/consumer/v1/consumers", baseURL)

	body, err := postJSON(consumer, client, URL, authHeader)
	if err != nil {
		return createdConsumer, err
	}

	err = json.Unmarshal(body, &createdConsumer)
	if err != nil {
		funcErr = fmt.Errorf("issue marshaling consumer POST response body: %s", err.Error())
		return
	}

	return createdConsumer, nil
}
