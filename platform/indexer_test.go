package platform

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/ahstn/defair/domain"
	"gopkg.in/yaml.v2"
)

func Test_YamlIndex_Read(t *testing.T) {
	config := domain.Index{
		Networks: map[string]domain.Network{
			"avax": {
				Endpoint: "test.rpc",
				Tokens: map[string]string{
					"avax": "123",
				},
			},
			"ftm": {
				Endpoint: "test.rpc",
				Markets: []domain.Market{
					{
						Name:  "Market1",
						Token: "Token",
					},
				},
			},
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}

	tmp, err := ioutil.TempFile("/tmp", "*.yaml")
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(tmp.Name(), data, 0)
	if err != nil {
		t.Fatal(err)
	}

	y := YamlIndex{Path: tmp.Name()}
	read, err := y.Read()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(config, *read) {
		t.Errorf("Read() got = %+v, expected = %+v", read, config)
	}
}
