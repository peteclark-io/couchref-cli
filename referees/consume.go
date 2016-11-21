package referees

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/peteclark-io/couchref-cli/structs"
	uuid "github.com/satori/go.uuid"
)

func (a Referees) readReferees() (*[]structs.Referee, error) {
	req, err := http.NewRequest("GET", a.RefsURL, nil)
	headers := req.Header
	headers.Add("origin", "https://www.premierleague.com")
	headers.Add("account", "premierleague")
	req.Header = headers

	resp, err := a.Client.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Failed to query for referee data!")
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, _ := httputil.DumpResponse(resp, true)
		logrus.WithField("resp", string(data)).WithField("status", resp.StatusCode).Error("Error!")
		return nil, errors.New("Shit just got real.")
	}

	dec := json.NewDecoder(resp.Body)
	responseData := struct {
		Content []struct {
			Name        structs.Name `json:"name"`
			Appearances float64      `json:"appearances"`
			YellowCards float64      `json:"yellowCards"`
			RedCards    float64      `json:"redCards"`
			Debut       struct {
				Kickoff struct {
					Millis float64 `json:"millis"`
				} `json:"kickoff"`
			} `json:"debut"`
			AlternativeIDs map[string]string `json:"altIds"`
			ID             float64           `json:"id"`
		} `json:"content"`
	}{}

	err = dec.Decode(&responseData)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse response json!")
	}

	var result []structs.Referee
	for _, ref := range responseData.Content {
		ref.AlternativeIDs["premierLeagueId"] = strconv.Itoa(int(ref.ID))
		result = append(result, structs.Referee{
			ID:             uuid.NewV4().String(),
			Name:           ref.Name,
			Appearances:    int(ref.Appearances),
			YellowCards:    int(ref.YellowCards),
			RedCards:       int(ref.RedCards),
			Debut:          time.Unix(0, int64(ref.Debut.Kickoff.Millis)*int64(time.Millisecond)),
			AlternativeIDs: ref.AlternativeIDs,
		})
	}

	return &result, nil
}
