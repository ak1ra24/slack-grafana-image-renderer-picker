package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ak1ra24/slack-grafana-image-renderer-picker/pkg/config"
	"github.com/ak1ra24/slack-grafana-image-renderer-picker/pkg/grafana"
	"github.com/ak1ra24/slack-grafana-image-renderer-picker/pkg/slack"

	"github.com/urfave/cli/v2"
)

const (
	name        = "gfslack"
	description = "upload grafana graph image to slack"
)

var Version string

func main() {
	if err := newApp().Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Version = Version
	app.Usage = description
	app.Authors = []*cli.Author{
		{
			Name:  "ak1ra24",
			Email: "ak1ra24net@gmail.com",
		},
	}
	flags := []cli.Flag{
		&cli.StringFlag{
			Name: "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "config.yaml",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Specify the dashboard name.",
		},
		&cli.StringFlag{
			Name: "time",
			Aliases: []string{"d"},
			Usage: "Specify from time e.g.) 1h",
		},
		&cli.StringFlag{
			Name: "tz",
			Usage: "Specify timezone",
			Value: "JST",
		},
	}

	app.Flags = flags
	app.Action = cmdAction

	return app
}

func cmdAction(c *cli.Context) error {
	cfg, err := config.Load(c.String("config"))
	if err != nil {
		return err
	}

	s := slack.NewSlack(cfg.Slack.Token, cfg.Slack.Channel)

	client := grafana.NewClient(cfg.Grafana.Endpoint, cfg.Grafana.ApiKey)

	if c.String("name") == "" {
		return errors.New("not set parameter...")
	}

	var from string
	if c.String("time") != "" {
		from, err = grafana.ParseTimeRange(c.String("time"))
		if err != nil {
			return err
		}
	}

	var g *grafana.Graph

	var db config.Dashboard

	for _, dashboard := range cfg.Dashboards {
		if dashboard.Name == c.String("name") {
			db = dashboard
		}
	}

	if from == "" {
		g, err = client.GetDsolo(db.Name, grafana.OrgId(db.OrgID), grafana.PanelId(db.PanelID), grafana.Tz("JST"))
	} else {
		g, err = client.GetDsolo(db.Name, grafana.OrgId(db.OrgID), grafana.PanelId(db.PanelID), grafana.Tz("JST"), grafana.From(from), grafana.To("now"))
	}

	if err != nil {
		return err
	}

	if err := s.PostImage(g); err != nil {
		return err
	}

	return nil
}
