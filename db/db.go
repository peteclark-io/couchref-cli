package db

import (
	"encoding/json"
	"os/exec"
)

type Firebase struct {
	AuthToken string
	Project   string
}

type DB interface {
	Write(path string, data interface{}) error
	Get(path string) ([]byte, error)
}

func NewFirebaseDB(authToken string, project string) DB {
	return &Firebase{AuthToken: authToken, Project: project}
}

func (f *Firebase) Write(path string, data interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	cmd := exec.Command("/usr/local/bin/firebase", "--non-interactive", "--project", f.Project, "--token", f.AuthToken, "database:set", "-y", path)
	pipe, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	pipe.Write(j)
	cmd.Start()

	return nil
}

func (f *Firebase) Get(path string) ([]byte, error) {
	cmd := exec.Command("/usr/local/bin/firebase", "--non-interactive", "--project", f.Project, "--token", f.AuthToken, "database:get", path)
	data, err := cmd.CombinedOutput()
	return data, err
}
