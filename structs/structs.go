package structs

import "time"

type Fixture struct {
	Id        string     `json:"id"`
	Home      string     `json:"home"`
	HomeScore int        `json:"home_score"`
	Away      string     `json:"away"`
	AwayScore int        `json:"away_score"`
	KickOff   time.Time  `json:"kick_off"`
	Matchday  int        `json:"-"`
	Referee   string     `json:"referee"`
	Questions []Question `json:"questions,omitempty"`
	Televised bool       `json:"televised"`
	Live      bool       `json:"live"`
}

type Club struct {
	Crest            string `json:"crestUrl"`
	SquadMarketValue string `json:"squadMarketValue"`
	ShortName        string `json:"shortName"`
	Name             string `json:"name"`
}

type Name struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Display string `json:"display"`
}

type Referee struct {
	ID             string            `json:"id"`
	Name           Name              `json:"name"`
	Appearances    int               `json:"appearances"`
	YellowCards    int               `json:"yellowCards"`
	RedCards       int               `json:"redCards"`
	Debut          time.Time         `json:"debut"`
	AlternativeIDs map[string]string `json:"alternativeIds"`
}

type Question struct {
	ID            string    `json:"id"`
	Controversial bool      `json:"controversial"`
	Asked         time.Time `json:"asked"`
	Decision      string    `json:"decision"`
	Description   string    `json:"description"`
	Match         string    `json:"match"`
	Question      string    `json:"question"`
	Time          string    `json:"time"`
}

type QuestionResults struct {
	ID        string           `json:"id"`
	Simple    SimpleStatistics `json:"simple"`
	Breakdown *Breakdown       `json:"breakdown,omitempty"`
}

type SimpleStatistics struct {
	Yes int `json:"yes"`
	No  int `json:"no"`
}

type Breakdown struct {
	Club ClubBreakdown `json:"club"`
	Age  AgeBreakdown  `json:"age"`
}

type ClubBreakdown map[string]SimpleStatistics
type AgeBreakdown map[string]SimpleStatistics
