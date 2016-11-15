package clubs

import (
	"encoding/json"

	"github.com/peteclark-io/couch-ref/tools/cli/structs"
)

type clubResponse struct {
	Teams []structs.Club `json:"teams"`
}

func (c Clubs) readClubs() (*[]structs.Club, error) {
	resp, err := c.Client.Get(c.ClubsURL)
	if err != nil {
		return nil, err
	}

	var results clubResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&results)

	var teams []structs.Club
	for _, team := range results.Teams {
		team.ShortName = c.mapShortname(team.ShortName)
		teams = append(teams, team)
	}

	if err != nil {
		return nil, err
	}

	return &teams, nil
}

func (c Clubs) mapShortname(short string) string {
	switch short {
	case "ManCity":
		return "Man City"
	case "Foxes":
		return "Leicester"
	case "West Bromwich":
		return "West Brom"
	case "ManU":
		return "Man Utd"
	case "Swans":
		return "Swansea"
	case "Spurs":
		return "Tottenham"
	case "Crystal":
		return "Crystal Palace"
	}
	return short
}
