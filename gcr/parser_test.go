package gcrevent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var InsertEvent = []byte(`{
    "action": "INSERT",
    "digest": "gcr.io/sakajunquality-public/test@sha256:abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcd",
    "tag": "gcr.io/sakajunquality-public/test:xxxx"
}`)

var DeleteTagEvent = []byte(`{
    "action": "DELETE",
    "tag": "gcr.io/sakajunquality-public/test:xxxxx"
}`)

var DeleteDigestEvent = []byte(`{
	"action": "DELETE",
	"digest": "gcr.io/sakajunquality-public/test@sha256:abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcd"
}`)

func TestParseEvent(t *testing.T) {
	insert, err := ParseMessage(InsertEvent)
	assert.Nil(t, err)
	assert.NotNil(t, insert)
	assert.Equal(t, ActionInsert, insert.Action)
	assert.NotNil(t, insert.Digest)
	assert.NotNil(t, insert.Tag)

	deleteTag, err := ParseMessage(DeleteTagEvent)
	assert.Nil(t, err)
	assert.NotNil(t, deleteTag)
	assert.Equal(t, ActionDelete, deleteTag.Action)
	assert.Nil(t, deleteTag.Digest)
	assert.NotNil(t, deleteTag.Tag)

	deleteDigest, err := ParseMessage(DeleteDigestEvent)
	assert.Nil(t, err)
	assert.NotNil(t, deleteDigest)
	assert.Equal(t, ActionDelete, deleteDigest.Action)
	assert.NotNil(t, deleteDigest.Digest)
	assert.Nil(t, deleteDigest.Tag)
}
