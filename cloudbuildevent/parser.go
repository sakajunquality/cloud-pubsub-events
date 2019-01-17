package cloudbuildevent

import "encoding/json"

func ParseMessage(m []byte) (Event, error) {
	var e Event
	err := json.Unmarshal(m, &e)
	return e, err
}
