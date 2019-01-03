package alienvault

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLicense(t *testing.T) {

	license, err := testClient.GetLicense()
	if err != nil {
		t.Fatalf("Error retrieving license: %s", err)
	}

	assert.True(t, license.ControlNodeLimit > 0)
	assert.True(t, license.SensorNodeLimit > 0)
}
