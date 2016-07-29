package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config is the global configuration for an instance.
type Config struct {
	WaitNextCheck int `yaml:"wait_time_next_check"`
	Slack         SlackConfiguration
	Hosts         []HostConfiguration
}

// File represents a YAML configuration file that namespaces all Kraken
// configuration under the top-level "kraken" key.
type File struct {
	PortScan Config `yaml:"sherlock"`
}

// HostConfiguration has all hosts that should be checked
type HostConfiguration struct {
	Hostname string
	Port     []string
}

type SlackConfiguration struct {
	Token     string
	Username  string
	IconEmoji string
	Channel   string
}

// Load the configuration
func Load(path string) (config *Config, err error) {
	var configFile File

	f, err := os.Open(os.ExpandEnv(path))
	if err != nil {
		return
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(d, &configFile)
	if err != nil {
		return
	}

	config = &configFile.PortScan

	return
}
