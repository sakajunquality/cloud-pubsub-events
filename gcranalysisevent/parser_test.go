package gcranalysisevent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var VulnerabilityEvent = []byte(`{
	"name": "projects/sakajunquality-public/occurrences/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
	"kind":"VULNERABILITY"
}`)

func TestParseEvent(t *testing.T) {
	event, err := ParseMessage(VulnerabilityEvent)
	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, event.Kind, EventVulnerbility)
}
