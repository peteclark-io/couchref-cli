package resources

import (
	"encoding/json"
	"net/http"

	"github.com/peteclark-io/couchref-cli/db"
	"github.com/peteclark-io/couchref-cli/structs"
)

func Fixtures(firebase db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := firebase.Get(`/v0/live-matches/`)
		if err != nil {
			write(w, 500, "Failed to get matches!")
			return
		}

		fixtures := make(map[string]structs.Fixture)
		err = json.Unmarshal(data, &fixtures)
		if err != nil {
			write(w, 500, "Failed to parse matches!")
			return
		}

		var result []structs.Fixture
		for _, fixture := range fixtures {
			result = append(result, fixture)
		}

		data, _ = json.MarshalIndent(result, "", "   ")
		w.Write(data)
	}
}

func GoalScored(firebase db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		request := struct {
			ID             string `json:"id"`
			HomeTeamScored bool   `json:"home_team_scored"`
		}{}
		err := dec.Decode(&request)

		if err != nil {
			write(w, 400, "Failed to parse request!")
			return
		}

		d, err := firebase.Get(`/v0/live-matches/` + request.ID)
		if err != nil {
			write(w, 400, "Can't find match!")
			return
		}

		fixture := structs.Fixture{}
		err = json.Unmarshal(d, &fixture)
		if err != nil {
			write(w, 500, "Can't parse match!")
			return
		}

		if request.HomeTeamScored {
			firebase.Write(`/v0/live-matches/`+fixture.Id+`/home_score`, fixture.HomeScore+1)
		} else {
			firebase.Write(`/v0/live-matches/`+fixture.Id+`/away_score`, fixture.AwayScore+1)
		}

		write(w, 200, "ok!")
	}
}
