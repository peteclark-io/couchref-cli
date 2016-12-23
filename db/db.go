package db

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/Sirupsen/logrus"
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

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	logrus.WithField("path", path).WithField("data", string(j)).Info("Writing to database!")

	cmd.Start()

	pipe.Write(j)
	pipe.Close()

	cmd.Wait()

	logrus.Info(outbuf.String())
	logrus.Error(errbuf.String())

	return nil
}

func (f *Firebase) Get(path string) ([]byte, error) {
	cmd := exec.Command("/usr/local/bin/firebase", "--non-interactive", "--project", f.Project, "--token", f.AuthToken, "database:get", path)
	data, err := cmd.CombinedOutput()
	return data, err
}
