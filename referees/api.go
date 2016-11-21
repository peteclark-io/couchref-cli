package referees

import (
	"net/http"

	"github.com/peteclark-io/couchref-cli/structs"
)

type Api interface {
	ReadReferees() (*[]structs.Referee, error)
}

type Referees struct {
	Client  *http.Client
	RefsURL string
}

func NewRefereesReader(client *http.Client) Api {
	return Referees{Client: client, RefsURL: "https://footballapi.pulselive.com/football/matchofficials?pageSize=500&comps=1&compSeasons=54&altIds=true&type=M&page=0"}
}

func (a Referees) ReadReferees() (*[]structs.Referee, error) {
	return a.readReferees()
}
