package main

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/peteclark-io/couchref-cli/db"
	"github.com/peteclark-io/couchref-cli/resources"
	cli "gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

func main() {
	app := cli.NewApp()
	app.Name = "couchref"
	app.Usage = "CLI for CouchRef"

	flags := []cli.Flag{
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "token",
			Usage: "Secret Firebase token.",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "project",
			Usage: "Firebase DB name.",
		}),
		cli.StringFlag{
			Name:  "config",
			Value: "./config.yml",
			Usage: "Path to the YAML config file.",
		},
	}

	app.Version = version()
	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))
	app.Flags = flags

	app.Action = func(c *cli.Context) error {
		firebase := db.NewFirebaseDB(c.String("token"), c.String("project"))
		server(8080, firebase)
		return nil
	}

	app.Run(os.Args)
}

func server(port int, firebase db.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/questions", resources.SaveNewQuestion(firebase)).Methods("PUT")
	r.HandleFunc("/matches", resources.Fixtures(firebase)).Methods("GET")
	r.HandleFunc("/scored", resources.GoalScored(firebase)).Methods("POST")

	addr := "localhost:" + strconv.Itoa(port)
	server := &http.Server{
		Handler: r,
		Addr:    addr,

		WriteTimeout: 60 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Info("Starting server on " + addr)
	server.ListenAndServe()
}

func parseURL(parse string) *url.URL {
	uri, err := url.Parse(parse)
	if err != nil {
		panic(err)
	}
	return uri
}

func version() string {
	v := os.Getenv("app_version")
	if v == "" {
		v = "v0.0.0"
	}
	return v
}
