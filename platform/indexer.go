package platform

import (
	"io/ioutil"

	"github.com/ahstn/defair/domain"
	"gopkg.in/yaml.v2"
)

// Indexer defines the expected behaviour of retrieving a domain.Index for a store.
type Indexer interface {
	Read() (domain.Index, error)
}

// YamlIndex is an implementation of Indexer using Yaml as a store.
type YamlIndex struct {
	Path string
}

// Read opens a file and parses it from YAML into the returned domain.Index.
func (y YamlIndex) Read() (domain.Index, error) {
	yamlFile, err := ioutil.ReadFile(y.Path)
	if err != nil {
		return domain.Index{}, err
	}

	c := &domain.Index{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return domain.Index{}, err
	}

	return *c, nil
}
