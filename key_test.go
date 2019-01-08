package alienvault

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestKeyManagement(t *testing.T) {

	// If you're having problems with the license not allowing more sensor keys, you can delete unclaimed keys caused by failing tests by uncommenting the following:
	// deleteAllKeys()
	// We don't do this all the time as we have no way of telling whether a real produciton key exists in the moment we make the API call. This is why Achi is asking AV for a separate test account.

	if ok, err := testClient.HasSensorAvailability(); err != nil {
		t.Fatalf("Failed to check sensor availability: %s", err)
	} else if !ok {
		t.Skip("Cannot test sensor key management, your license does not have room for more sensors.")
	}

	if ok, err := testClient.HasSensorKeyAvailability(); err != nil {
		t.Fatalf("Failed to check sensor key availability: %s", err)
	} else if !ok {
		t.Skip("Cannot test sensor key management, your license does not have room for more sensor keys.")
	}

	key, err := testClient.CreateSensorKey(false)
	if err != nil {
		t.Fatalf("Failed to create sensor key: %s", err)
	}

	refreshed, err := testClient.GetSensorKey(key.ID)
	if err != nil {
		t.Fatalf("Failed to refresh sensor key: %s", err)
	}

	assert.Equal(t, refreshed.ID, key.ID, "Refreshed key should contain the original ID")

	require.NotEmpty(t, key.ID, "Key should have an ID assigned")

	require.Nil(t, testClient.DeleteSensorKey(key))

	keys, err := testClient.GetSensorKeys()
	if err != nil {
		t.Fatalf("Failed to list sensor keys: %s", err)
	}

	for _, k := range keys {
		if k.ID == key.ID {
			t.Fatalf("Key '%s' still exists after deletion", k.ID)
		}
	}

}

func deleteAllKeys() error {
	keys, err := testClient.GetSensorKeys()
	if err != nil {
		return fmt.Errorf("Failed to list sensor keys: %s", err)
	}

	for _, k := range keys {
		err := testClient.DeleteSensorKey(&k)
		if err != nil {
			return err
		}
	}
	return nil
}
