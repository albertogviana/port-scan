package config

import (
	"testing"
)

type testConfig struct {
	hostname string
	port     []string
}

var configLoadTest = []testConfig{
	{"host01.test.com", []string{"80", "443"}},
	{"host02.test.com", []string{"22"}},
}

func TestLoad(t *testing.T) {
	configFile := "../test_data/config.yml"
	config, _ := Load(configFile)

	if len(config.Hosts) != 2 {
		t.Error(
			"It was expected to find 2 hosts on config.Hosts\n ",
			"but it was found ", len(config.Hosts),
		)
	}

	for i, v := range config.Hosts {
		if v.Hostname != configLoadTest[i].hostname {
			t.Error(
				"For: ", configLoadTest[i].hostname, "\n",
				"got: ", v,
			)
		}

		if len(v.Port) != len(configLoadTest[i].port) {
			t.Error(
				"For: ", len(configLoadTest[i].port), "\n",
				"got: ", len(v.Port),
			)
		}

		for index, port := range v.Port {
			if port != configLoadTest[i].port[index] {
				t.Error(
					"For: ", configLoadTest[i].port[index], "\n",
					"got: ", port,
				)
			}
		}
	}
}
