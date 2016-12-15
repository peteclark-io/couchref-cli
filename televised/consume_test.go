package televised

import (
	"net/http"
	"testing"

	"github.com/peteclark-io/couchref-cli/structs"
)

func TestScraping(t *testing.T) {
	reader := NewTelevisedMatchReader(&http.Client{})
	fixtures := []structs.Fixture{
		{
			Away: "Chelsea",
			Home: "Crystal Palace",
		},
		{
			Away: "Watford",
			Home: "Sunderland",
		},
	}
	reader.CheckTelevisedMatches(fixtures)
}
