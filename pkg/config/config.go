package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Slack Slack `yaml:"slack"`
	Grafana Grafana `yaml:"grafana"`
	Dashboards []Dashboard `yaml:"dashboards"`
}

type Slack struct {
	Token   string `yaml:"token"`
	Channel string `yaml:"channel"`
}

type Grafana struct {
	UseClientAuth bool   `yaml:"use_client_auth"`
	ClientAuthP12 string `yaml:"client_auth_p12"`
	Endpoint      string `yaml:"endpoint"`
	ApiKey		  string `yaml:"apikey"`
}

type Dashboard struct {
	Name          string `yaml:"name"`
	DashboardName string `yaml:"dashboardName"`
	OrgID         string `yaml:"orgId"`
	PanelID       string `yaml:"panelId"`
}

func Load(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := Config{}
	if err := yaml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
