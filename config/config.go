package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config is the global configuration for an instance.
type Config struct {
	Hosts []HostConfiguration
}

// File represents a YAML configuration file that namespaces all Kraken
// configuration under the top-level "kraken" key.
type File struct {
	PortScan Config `yaml:"port_scan"`
}

// HostConfiguration has all hosts that should be checked
type HostConfiguration struct {
	Hostname string
	Port     []string
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
