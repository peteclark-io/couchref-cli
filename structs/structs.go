package structs

type Fixture struct {
	Id        string     `json:"id"`
	Home      string     `json:"home"`
	HomeScore int        `json:"home_score"`
	Away      string     `json:"away"`
	AwayScore int        `json:"away_score"`
	KickOff   string     `json:"kick_off"`
	Matchday  int        `json:"-"`
	Referee   string     `json:"referee"`
	Questions []Question `json:"questions,omitempty"`
}

type Club struct {
	Crest            string `json:"crestUrl"`
	SquadMarketValue string `json:"squadMarketValue"`
	ShortName        string `json:"shortName"`
	Name             string `json:"name"`
}

type Question struct {
	ID string `json:"id"`
}

type QuestionResults struct {
	ID        string           `json:"id"`
	Simple    SimpleStatistics `json:"simple"`
	Breakdown Breakdown        `json:"breakdown"`
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
