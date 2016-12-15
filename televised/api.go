package televised

import (
	"net/http"

	"github.com/peteclark-io/couchref-cli/structs"
)

type Api interface {
	CheckTelevisedMatches(fixtures []structs.Fixture) (*map[string]structs.Fixture, error)
}

type TelevisedMatches struct {
	Client              *http.Client
	TelevisedMatchesURL string
}

func NewTelevisedMatchReader(client *http.Client) Api {
	return TelevisedMatches{Client: client, TelevisedMatchesURL: "http://www.wheresthematch.com/barclays-premier-league/"}
}

func (a TelevisedMatches) CheckTelevisedMatches(fixtures []structs.Fixture) (*map[string]structs.Fixture, error) {
	return a.enrichFixtures(fixtures)
}
