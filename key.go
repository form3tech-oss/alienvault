package alienvault

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SensorKey is a key used to activate a sensor. The ID is traditionally used as an auth code to activate a sensor using the web UI.
type SensorKey struct {
	ID        string `json:"id"`
	Consumed  bool
	CreatedAt int     `json:"createdAt"`
	ExpiresAt int     `json:"expires"`
	NodeID    *string `json:"nodeId"`
}

// CreateSensorKey will create a new key used to activate a sensor. However, if the useExisting option is used, and an unused key already exists, this will be returned instead.
func (client *Client) CreateSensorKey() (*SensorKey, error) {

	req, err := client.createRequest("POST", "/sensors/key", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var key SensorKey
	if err := json.NewDecoder(resp.Body).Decode(&key); err != nil {
		return nil, err
	}

	return &key, nil
}

// GetSensorKeys returns a list of all sensor keys on the account
func (client *Client) GetSensorKeys() ([]SensorKey, error) {

	req, err := client.createRequest("GET", "/sensors/key", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var keys []SensorKey
	if err := json.NewDecoder(resp.Body).Decode(&keys); err != nil {
		return nil, err
	}

	return keys, nil
}

// GetSensorKey returns a particular sensor key identified by the supplied id
func (client *Client) GetSensorKey(id string) (*SensorKey, error) {

	// There is no GET for a singular key in the AV API atm

	keys, err := client.GetSensorKeys()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		if key.ID == id {
			return &key, nil
		}
	}

	// if the key is not found mark it as consumed in the returned value, as keys are only available temporarily
	return &SensorKey{
		ID:       id,
		Consumed: true,
	}, nil
}

// DeleteSensorKey deletes a particular sensor key as identified by the supplied id
func (client *Client) DeleteSensorKey(key *SensorKey) error {

	req, err := client.createRequest("DELETE", fmt.Sprintf("/sensors/key/%s", key.ID), nil)
	if err != nil {
		return err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response code when deleting key: %d", resp.StatusCode)
	}

	return nil
}
