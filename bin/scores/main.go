package main

import (
	"encoding/json"
	"net/url"
	"os"

	"github.com/peteclark-io/couchref-cli/gamification"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "scores"
	app.Usage = "Score generator!"

	flags := []cli.Flag{
		cli.Float64Flag{
			Name:  "result",
			Usage: "The overall verdict of the question in percent.",
		},
		cli.Float64Flag{
			Name:  "confidence",
			Usage: "The overall confidence of the verdict in percent.",
		},
		cli.BoolFlag{
			Name:  "vote",
			Usage: "The users vote!",
		},
	}

	app.Version = version()
	app.Flags = flags
	app.Action = func(c *cli.Context) error {
		score := gamification.Score(c.Bool("vote"), c.Float64("result"), c.Float64("confidence"))
		enc := json.NewEncoder(os.Stdout)

		return enc.Encode(struct {
			Score float64 `json:"score"`
		}{score})
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
