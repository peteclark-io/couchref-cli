package resources

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/peteclark-io/couchref-cli/db"
	"github.com/peteclark-io/couchref-cli/structs"
	uuid "github.com/satori/go.uuid"
)

func zeroStat() structs.SimpleStatistics {
	return structs.SimpleStatistics{
		Yes: 0,
		No:  0,
	}
}

func SaveNewQuestion(firebase db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		question := structs.Question{}
		err := dec.Decode(&question)
		if err != nil {
			write(w, 500, "Failed to decode body, make sure it's json.")
			return
		}

		id := uuid.NewV4()
		question.ID = id.String()
		question.Asked = time.Now()

		statistics := structs.QuestionResults{
			ID:        question.ID,
			Simple:    zeroStat(),
			Breakdown: nil,
		}

		time, err := validateTime(question.Time)
		if err != nil {
			write(w, 400, err.Error())
			return
		}

		question.Time = time

		firebase.Write(`/v0/live-questions/`+question.ID, question)
		firebase.Write(`/v0/live-statistics/`+statistics.ID, statistics)
		w.WriteHeader(200)
	}
}
