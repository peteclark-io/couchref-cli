package televised

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/peteclark-io/couchref-cli/structs"
)

func (a TelevisedMatches) enrichFixtures(fixtures []structs.Fixture) (*map[string]structs.Fixture, error) {
	results := make(map[string]structs.Fixture)
	for _, f := range fixtures {
		results[f.Id] = f
	}

	doc, err := goquery.NewDocument(a.TelevisedMatchesURL)
	if err != nil {
		logrus.WithError(err).Error("Failed to scrape!")
		return nil, nil
	}

	doc.Find("tr").Each(func(i int, fixture *goquery.Selection) {
		var clubs []string
		fixture.Find(".fixture em").Each(func(i1 int, em *goquery.Selection) {
			if em.Text() == "Live Stream" {
				return
			}

			clubs = append(clubs, a.mapName(em.Text()))
		})

		length := fixture.Find(`.channel-details img[src="../images/newchannels/not-televised.gif"]`).Length()

		for _, f := range fixtures {
			if len(clubs) == 2 && strings.Contains(clubs[0], f.Home) && strings.Contains(clubs[1], f.Away) {
				if length > 0 {
					f.Televised = false
				} else {
					f.Televised = true
				}
				results[f.Id] = f
			}
		}
	})

	return &results, nil
}

func (a TelevisedMatches) mapName(name string) string {
	switch name {
	case "Manchester City":
		return "Man City"
	case "West Bromwich Albion":
		return "West Brom"
	case "Manchester United":
		return "Man Utd"
	}
	return name
}
