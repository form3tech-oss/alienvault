package alienvault

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyManagement(t *testing.T) {

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

	assert.NotEmpty(t, key.ID, "Key should have an ID assigned")

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
