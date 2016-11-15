package simulation

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"

	"github.com/peteclark-io/couch-ref/tools/cli/db"
	"github.com/peteclark-io/couch-ref/tools/cli/structs"
)

var genders = []string{"Male", "Female"}
var countries = []string{"Bangladesh", "Belgium", "Burkina Faso", "Bulgaria", "Bosnia and Herzegovina", "Barbados", "Wallis and Futuna", "Saint Barthelemy", "Bermuda", "Brunei", "Bolivia", "Bahrain", "Burundi", "Benin", "Bhutan", "Jamaica", "Bouvet Island", "Botswana", "Samoa", "Bonaire, Saint Eustatius and Saba", "Brazil", "Bahamas", "Jersey", "Belarus", "Belize", "Russia", "Rwanda", "Serbia", "East Timor", "Reunion", "Turkmenistan", "Tajikistan", "Romania", "Tokelau", "Guinea-Bissau", "Guam", "Guatemala", "South Georgia and the South Sandwich Islands", "Greece", "Equatorial Guinea", "Guadeloupe", "Japan", "Guyana", "Guernsey", "French Guiana", "Georgia", "Grenada", "United Kingdom", "Gabon", "El Salvador", "Guinea", "Gambia", "Greenland", "Gibraltar", "Ghana", "Oman", "Tunisia", "Jordan", "Croatia", "Haiti", "Hungary", "Hong Kong", "Honduras", "Heard Island and McDonald Islands", "Venezuela", "Puerto Rico", "Palestinian Territory", "Palau", "Portugal", "Svalbard and Jan Mayen", "Paraguay", "Iraq", "Panama", "French Polynesia", "Papua New Guinea", "Peru", "Pakistan", "Philippines", "Pitcairn", "Poland", "Saint Pierre and Miquelon", "Zambia", "Western Sahara", "Estonia", "Egypt", "South Africa", "Ecuador", "Italy", "Vietnam", "Solomon Islands", "Ethiopia", "Somalia", "Zimbabwe", "Saudi Arabia", "Spain", "Eritrea", "Montenegro", "Moldova", "Madagascar", "Saint Martin", "Morocco", "Monaco", "Uzbekistan", "Myanmar", "Mali", "Macao", "Mongolia", "Marshall Islands", "Macedonia", "Mauritius", "Malta", "Malawi", "Maldives", "Martinique", "Northern Mariana Islands", "Montserrat", "Mauritania", "Isle of Man", "Uganda", "Tanzania", "Malaysia", "Mexico", "Israel", "France", "British Indian Ocean Territory", "Saint Helena", "Finland", "Fiji", "Falkland Islands", "Micronesia", "Faroe Islands", "Nicaragua", "Netherlands", "Norway", "Namibia", "Vanuatu", "New Caledonia", "Niger", "Norfolk Island", "Nigeria", "New Zealand", "Nepal", "Nauru", "Niue", "Cook Islands", "Kosovo", "Ivory Coast", "Switzerland", "Colombia", "China", "Cameroon", "Chile", "Cocos Islands", "Canada", "Republic of the Congo", "Central African Republic", "Democratic Republic of the Congo", "Czech Republic", "Cyprus", "Christmas Island", "Costa Rica", "Curacao", "Cape Verde", "Cuba", "Swaziland", "Syria", "Sint Maarten", "Kyrgyzstan", "Kenya", "South Sudan", "Suriname", "Kiribati", "Cambodia", "Saint Kitts and Nevis", "Comoros", "Sao Tome and Principe", "Slovakia", "South Korea", "Slovenia", "North Korea", "Kuwait", "Senegal", "San Marino", "Sierra Leone", "Seychelles", "Kazakhstan", "Cayman Islands", "Singapore", "Sweden", "Sudan", "Dominican Republic", "Dominica", "Djibouti", "Denmark", "British Virgin Islands", "Germany", "Yemen", "Algeria", "United States", "Uruguay", "Mayotte", "United States Minor Outlying Islands", "Lebanon", "Saint Lucia", "Laos", "Tuvalu", "Taiwan", "Trinidad and Tobago", "Turkey", "Sri Lanka", "Liechtenstein", "Latvia", "Tonga", "Lithuania", "Luxembourg", "Liberia", "Lesotho", "Thailand", "French Southern Territories", "Togo", "Chad", "Turks and Caicos Islands", "Libya", "Vatican", "Saint Vincent and the Grenadines", "United Arab Emirates", "Andorra", "Antigua and Barbuda", "Afghanistan", "Anguilla", "U.S. Virgin Islands", "Iceland", "Iran", "Armenia", "Albania", "Angola", "Antarctica", "American Samoa", "Argentina", "Australia", "Austria", "Aruba", "India", "Aland Islands", "Azerbaijan", "Ireland", "Indonesia", "Ukraine", "Qatar", "Mozambique"}
var ageGroups = []string{"< 20", "21 - 30", "31 - 40", "41 - 50", "51 - 60", "60+"}

type Simulation struct {
	PercentFans float64
	Total       int
	fixtures    []structs.Fixture
	clubs       []structs.Club
}

type userSimulator struct {
	Fan      bool   `json:"fan"`
	Sex      string `json:"sex"`
	Location string `json:"location"`
	Age      string `json:"age"`
	Club     string `json:"club"`
}

func NewSimulation(total int, percentFans float64) *Simulation {
	return &Simulation{Total: total, PercentFans: percentFans}
}

func (s *Simulation) Simulate(firebase db.DB) error {
	fixtures, err := db.AsFixture(firebase.Get("/v0/live-matches"))
	if err != nil {
		return err
	}
	s.fixtures = fixtures

	clubs, err := db.AsClub(firebase.Get("/v0/clubs"))
	if err != nil {
		return err
	}
	s.clubs = clubs

	if len(clubs) <= 0 {
		return errors.New("Failed to load clubs :/")
	}

	for _, fixture := range fixtures {
		users, err := s.createUsers(fixture)
		if err != nil {
			return err
		}

		enc := json.NewEncoder(os.Stdout)
		enc.Encode(users)
	}

	return nil
}

/*func performSimulation(user userSimulator, fixture structs.Fixture, firebase db.DB) {
	for _, q := range fixture.Questions {
		path := "/v0/live-statistics/" + q.ID

		firebase.Write(path, data)
	}
}*/

func (s *Simulation) createUsers(fixture structs.Fixture) ([]userSimulator, error) {
	numFans := int(float64(s.Total) * s.PercentFans)
	numNeutrals := s.Total - numFans

	var users []userSimulator
	fans := fanClubs(fixture, s.clubs)
	if len(fans) <= 0 {
		return users, errors.New("Failed to find fan clubs!")
	}

	for i := 0; i < numFans; i++ {
		users = append(users, userSimulator{Fan: true, Sex: random(genders), Location: random(countries), Age: random(ageGroups), Club: randomClub(fans).ShortName})
	}

	for i := 0; i < numNeutrals; i++ {
		users = append(users, userSimulator{Fan: false, Sex: random(genders), Location: random(countries), Age: random(ageGroups), Club: randomClub(s.clubs).ShortName})
	}

	return users, nil
}

func fanClubs(fixture structs.Fixture, clubs []structs.Club) []structs.Club {
	var fanClubs []structs.Club
	for _, c := range clubs {
		if c.ShortName == fixture.Home || c.ShortName == fixture.Away {
			fanClubs = append(fanClubs, c)
		}
	}
	return fanClubs
}

func randomClub(from []structs.Club) structs.Club {
	return from[rand.Intn(len(from))]
}

func random(from []string) string {
	return from[rand.Intn(len(from))]
}
