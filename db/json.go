package db

import (
	"encoding/json"

	"github.com/peteclark-io/couchref-cli/structs"
)

func AsFixture(data []byte, err error) ([]structs.Fixture, error) {
	var fixtures []structs.Fixture
	if err != nil {
		return fixtures, err
	}

	var mapped map[string]structs.Fixture
	err = json.Unmarshal(data, &mapped)

	for _, v := range mapped {
		fixtures = append(fixtures, v)
	}

	return fixtures, err
}

func AsClub(data []byte, err error) ([]structs.Club, error) {
	var clubs []structs.Club
	if err != nil {
		return clubs, err
	}

	err = json.Unmarshal(data, &clubs)
	return clubs, err
}
