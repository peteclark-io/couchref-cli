package resources

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Msg string `json:"message"`
}

func write(w http.ResponseWriter, status int, msg string) error {
	j, err := json.Marshal(message{msg})
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	w.Write(j)

	return nil
}
