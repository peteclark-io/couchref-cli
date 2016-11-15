package fixtures

import (
	"encoding/json"

	"github.com/peteclark-io/couch-ref/tools/cli/structs"
	uuid "github.com/satori/go.uuid"
)

type fixturesResponse struct {
	Fixtures []fixture `json:"fixtures"`
}

type fixture struct {
	Date     string `json:"date"`
	Matchday int    `json:"matchday"`
	Home     string `json:"homeTeamName"`
	Away     string `json:"awayTeamName"`
}

func (f Fixtures) readFixtures() (*[]structs.Fixture, error) {
	resp, err := f.Client.Get(f.FixturesURL)
	if err != nil {
		return nil, err
	}

	var results fixturesResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&results)

	if err != nil {
		return nil, err
	}

	fixtures := make([]structs.Fixture, 0)
	for _, f := range results.Fixtures {
		fixtures = append(fixtures, mapFixture(f))
	}

	return &fixtures, nil
}

func mapFixture(f fixture) structs.Fixture {
	return structs.Fixture{
		Id:        uuid.NewV4().String(),
		Home:      f.Home,
		Away:      f.Away,
		HomeScore: 0,
		AwayScore: 0,
		KickOff:   f.Date,
		Matchday:  f.Matchday,
		Referee:   "",
	}
}
