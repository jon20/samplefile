package config

import (
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Image string `yaml:"image"`
}

func NewConfigFromBytes(b []byte) (*Config, error) {
	c := &Config{}

	err := yaml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewFromFilename(name string) (*Config, error) {
	d, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	contents, _ := ioutil.ReadAll(d)
	return NewConfigFromBytes(contents)
}
