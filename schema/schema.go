package schema

import (
	_ "embed"
	"encoding/json"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v5"

	v1 "github.com/mfranczy/compat/pkg/types/v1"
)

const (
	CompatibilityURL = "https://opencontainers.org/schema/image-compatibility"
)

//go:embed schema.json
var Schema string

func LoadSchema(artifact string) (*v1.ImageCompatibilitySchema, error) {
	s, err := jsonschema.CompileString(CompatibilityURL, Schema)

	if err != nil {
		return nil, err
	}

	d, err := os.ReadFile(artifact)
	if err != nil {
		return nil, err
	}

	var v interface{}
	if err = json.Unmarshal(d, &v); err != nil {
		return nil, err
	}

	if err = s.Validate(v); err != nil {
		return nil, err
	}

	var i v1.ImageCompatibilitySchema
	if err = json.Unmarshal(d, &i); err != nil {
		return nil, err
	}

	return &i, nil
}
