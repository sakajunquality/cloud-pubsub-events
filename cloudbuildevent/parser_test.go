package cloudbuildevent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var successedMessage = []byte(`{
    "id": "xxxxxxxx-yyyy-zzzz-aaa-bbbbbbbbbbbb",
    "projectId": "my-project",
    "status": "SUCCESS",
    "source": {
        "repoSource": {
            "projectId": "my-project",
            "repoName": "github-sakajunquality-test",
            "tagName": "v123"
        }
    },
    "steps": [
        {
            "name": "gcr.io/cloud-builders/docker",
            "args": [
                "build",
                "-t",
                "gcr.io/my-project/test:v123",
                "."
            ],
            "timing": {
                "startTime": "2019-01-16T07:00:06.025539143Z",
                "endTime": "2019-01-16T07:02:43.852546643Z"
            },
            "status": "SUCCESS"
        },
        {
            "name": "gcr.io/cloud-builders/docker",
            "args": [
                "push",
                "gcr.io/my-project/test:v123"
            ],
            "timing": {
                "startTime": "2019-01-16T07:02:43.852561245Z",
                "endTime": "2019-01-16T07:03:18.060770180Z"
            },
            "status": "SUCCESS"
        }
    ],
    "results": {
        "buildStepImages": [
            "sha256:1234567890123456789012345678901234567890123456789012345678901234",
            "sha256:5678901234567890123456789012345678901234567890123456789012345678"
        ],
        "buildStepOutputs": []
    },
    "createTime": "2019-01-16T07:00:00.830888846Z",
    "startTime": "2019-01-16T07:00:01.504078791Z",
    "finishTime": "2019-01-16T07:03:18.494491Z",
    "timeout": "600s",
    "logsBucket": "gs://123123123123.cloudbuild-logs.googleusercontent.com",
    "sourceProvenance": {
        "resolvedRepoSource": {
            "projectId": "my-project",
            "repoName": "github-sakajunquality-test",
            "commitSha": "abcdefgabcdefgabcdefgabcdefgabcdefg12345"
        }
    },
    "buildTriggerId": "11111111-2222-3333-aaaa-bbbcccdddeee",
    "options": {
        "substitutionOption": "ALLOW_LOOSE",
        "logging": "LEGACY"
    },
    "logUrl": "https://console.cloud.google.com/gcr/builds/405fbe89-9c17-4cf0-b0c6-deb0f75a928a?project=123123123123",
    "tags": [
        "event-zzzzzzzz-yyyy-xxxx-wwww-00000000000",
        "trigger-11111111-2222-3333-aaaa-bbbcccdddeee"
    ],
    "timing": {
        "BUILD": {
            "startTime": "2019-01-16T07:00:06.025504626Z",
            "endTime": "2019-01-16T07:03:18.112142710Z"
        },
        "FETCHSOURCE": {
            "startTime": "2019-01-16T07:00:02.809327658Z",
            "endTime": "2019-01-16T07:00:06.010241815Z"
        }
    }
}`)

func TestParseMessage(t *testing.T) {
	e, err := ParseMessage(successedMessage)

	assert.Nil(t, err)
	assert.NotNil(t, e)
	assert.Equal(t, true, e.IsFinished())
	assert.Equal(t, true, e.IsSuuccess())
	assert.Equal(t, true, e.IsTriggerdBuild())
}
