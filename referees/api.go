package referees

import (
	"net/http"

	"github.com/peteclark-io/couchref-cli/structs"
)

type Api interface {
	ReadReferees() (map[string]structs.Referee, error)
}

type Referees struct {
	Client  *http.Client
	RefsURL string
}

func NewRefereesReader(client *http.Client) Api {
	return Referees{Client: client, RefsURL: "https://footballapi.pulselive.com/football/matchofficials?pageSize=500&comps=1&compSeasons=54&altIds=true&type=M&page=0"}
}

func (a Referees) ReadReferees() (map[string]structs.Referee, error) {
	return toMap(a.readReferees())
}

func toMap(refs *[]structs.Referee, err error) (map[string]structs.Referee, error) {
	res := make(map[string]structs.Referee)
	if err != nil {
		return res, err
	}

	for _, val := range *refs {
		res[val.ID] = val
	}
	return res, nil
}
