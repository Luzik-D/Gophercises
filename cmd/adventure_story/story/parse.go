package story

import (
	"encoding/json"
	"io"
)

func ParseJson(r io.Reader) (StoryMap, error) {
	var story StoryMap
	jsonReader := json.NewDecoder(r)
	err := jsonReader.Decode(&story)

	if err != nil {
		return nil, err
	}

	return story, nil
}
