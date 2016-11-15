package fixtures

import (
	"net/http"

	"github.com/peteclark-io/couch-ref/tools/cli/clubs"
	"github.com/peteclark-io/couch-ref/tools/cli/structs"
)

type Api interface {
	ReadFixtures(matchday int) (*[]structs.Fixture, error)
}

type Fixtures struct {
	Client      *http.Client
	FixturesURL string
}

func NewFixturesReader(client *http.Client) Api {
	return Fixtures{Client: client, FixturesURL: "http://api.football-data.org/v1/competitions/426/fixtures"}
}

func (f Fixtures) ReadFixtures(matchday int) (*[]structs.Fixture, error) {
	all, err := f.readFixtures()
	if err != nil {
		return nil, err
	}

	clubReader := clubs.NewClubReader(f.Client)
	cs, err := clubReader.ReadClubs()

	filtered := make([]structs.Fixture, 0)
	for _, fixture := range *all {
		if fixture.Matchday == matchday {
			updateClubName(*cs, &fixture)
			filtered = append(filtered, fixture)
		}
	}

	return &filtered, nil
}

func updateClubName(clubs []structs.Club, fixture *structs.Fixture) {
	for _, c := range clubs {
		if fixture.Home == c.Name {
			fixture.Home = c.ShortName
		}

		if fixture.Away == c.Name {
			fixture.Away = c.ShortName
		}
	}
}

func ToMap(fixtures []structs.Fixture) map[string]structs.Fixture {
	m := make(map[string]structs.Fixture)
	for _, f := range fixtures {
		m[f.Id] = f
	}
	return m
}
