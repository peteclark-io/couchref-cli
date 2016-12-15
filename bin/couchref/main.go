package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/peteclark-io/couchref-cli/clubs"
	"github.com/peteclark-io/couchref-cli/db"
	"github.com/peteclark-io/couchref-cli/fixtures"
	"github.com/peteclark-io/couchref-cli/referees"
	"github.com/peteclark-io/couchref-cli/simulation"
	"github.com/peteclark-io/couchref-cli/televised"
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

	app.Commands = []cli.Command{
		{
			Name:    "fixtures",
			Aliases: []string{"f"},
			Flags: append(flags, cli.IntFlag{
				Name:  "m, matchday",
				Usage: "The matchday to query for.",
			}),
			Usage: "Query for fixtures for the provided matchday.",
			Action: func(c *cli.Context) error {
				client := &http.Client{}
				api := fixtures.NewFixturesReader(client)

				fs, err := api.ReadFixtures(c.Int("matchday"))
				if err != nil {
					return err
				}

				televisedApi := televised.NewTelevisedMatchReader(client)
				result, err := televisedApi.CheckTelevisedMatches(*fs)
				if err != nil {
					return err
				}

				d, _ := json.Marshal(result)
				os.Stdout.Write(d)

				return err
			},
		},
		{
			Name:    "clubs",
			Aliases: []string{"c"},
			Usage:   "Download club data.",
			Action: func(c *cli.Context) error {
				client := &http.Client{}
				api := clubs.NewClubReader(client)

				clubs, err := api.ReadClubs()
				if err != nil {
					return err
				}

				d, _ := json.Marshal(clubs)
				os.Stdout.Write(d)

				return err
			},
		},
		{
			Name:    "referees",
			Aliases: []string{"r"},
			Usage:   "Download referee data.",
			Action: func(c *cli.Context) error {
				client := &http.Client{}
				api := referees.NewRefereesReader(client)

				refs, err := api.ReadReferees()
				if err != nil {
					return err
				}

				d, _ := json.Marshal(refs)
				os.Stdout.Write(d)

				return err
			},
		},
		{
			Name:    "simulate",
			Aliases: []string{"s"},
			Usage:   "Generate simulated user data.",
			Flags: append(flags, cli.StringFlag{
				Name:  "match",
				Usage: "Match id as a uuid",
			}),
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				sim := simulation.NewSimulation(600, 0.75)
				firebase := db.NewFirebaseDB(c.String("token"), c.String("project"))
				err := sim.Simulate(firebase, c.String("match"))
				return err
			},
		},
	}

	app.Run(os.Args)
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
