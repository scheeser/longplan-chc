package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Common static strings used in HTTP service calls.
const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
	ApplicationJSON     = "application/json"
)

func getJSON(client http.Client, URL string, authHeader string) (respBody []byte, funcErr error) {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		funcErr = fmt.Errorf("issue creating resource GET for %s: %s", URL, err.Error())
		return
	}

	req.Header.Add(HeaderAuthorization, authHeader)

	resp, err := client.Do(req)
	if err != nil {
		funcErr = fmt.Errorf("problem performing GET request for %s: %s", URL, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		funcErr = fmt.Errorf("GET to %s returned with status %d", URL, resp.StatusCode)
		return
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		funcErr = fmt.Errorf("issue reading GET response body for %s: %s", URL, err.Error())
		return
	}

	return respBody, nil
}

func postJSON(body interface{}, client http.Client, URL string, authHeader string) (respBody []byte, funcErr error) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest(http.MethodPost, URL, buf)
	if err != nil {
		funcErr = fmt.Errorf("issue creating resource POST for %s: %s", URL, err.Error())
		return
	}

	req.Header.Add(HeaderAuthorization, authHeader)
	req.Header.Add(HeaderContentType, ApplicationJSON)

	resp, err := client.Do(req)
	if err != nil {
		funcErr = fmt.Errorf("problem performing POST request for %s: %s", URL, err.Error())
		return
	}

	if resp.StatusCode != http.StatusCreated {
		funcErr = fmt.Errorf("POST to %s returned with status %d", URL, resp.StatusCode)
		return
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		funcErr = fmt.Errorf("issue reading POST response body for %s: %s", URL, err.Error())
		return
	}

	return respBody, nil
}

func putJSON(body interface{}, client http.Client, URL string, authHeader string) (respBody []byte, funcErr error) {
	buf := new(bytes.Buffer)
	if body != nil {
		json.NewEncoder(buf).Encode(body)
	}

	req, err := http.NewRequest(http.MethodPut, URL, buf)
	if err != nil {
		funcErr = fmt.Errorf("issue creating resource PUT for %s: %s", URL, err.Error())
		return
	}

	req.Header.Add(HeaderAuthorization, authHeader)
	req.Header.Add(HeaderContentType, ApplicationJSON)

	resp, err := client.Do(req)
	if err != nil {
		funcErr = fmt.Errorf("problem performing PUT request for %s: %s", URL, err.Error())
		return
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) {
		funcErr = fmt.Errorf("PUT to %s returned with status %d", URL, resp.StatusCode)
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()

		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			funcErr = fmt.Errorf("issue reading PUT response body for %s: %s", URL, err.Error())
			return
		}
	}

	return respBody, nil
}
