package referees

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/peteclark-io/couchref-cli/structs"
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

	dec := json.NewDecoder(resp.Body)
	responseData := struct {
		Name           structs.Name      `json:"name"`
		Appearances    int               `json:"appearances"`
		YellowCards    int               `json:"yellowCards"`
		RedCards       int               `json:"redCards"`
		Debut          string            `json:"debut"`
		AlternativeIDs map[string]string `json:"altIds"`
		ID             int               `json:"id"`
	}{}

	err = dec.Decode(&responseData)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse response json!")
	}

}
