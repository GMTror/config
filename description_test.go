package config

import (
	"testing"
)

type ConfigDescription struct {
	String string `env:"STRING"default:"test_string"description:"this is description"`
	Struct struct {
		Struct struct {
			String string `env:"STRING"default:"test_string"description:"this is description Struct>Struct>String"`
		}
		PasStruct struct {
			String string `env:"PAS_STRING"default:"test_string"description:"this is description Struct>PasStruct>String"`
		} `env:"-"description:"this is description Struct>PasStruct"`
	} `env:"STRUCT"description:"This is description Struct"`
}

func TestDescription(t *testing.T) {
	c := &ConfigDescription{}

	des, err := Description(c)
	if err != nil {
		t.Error(err)
	}

	if des != "STRING - this is description (defauilt: test_string)\nThis is description Struct\nSTRUCT_STRING - this is description Struct>Struct>String (defauilt: test_string)" {
		t.Error("description")
	}
}
