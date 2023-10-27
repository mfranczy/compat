package schema

import (
	_ "embed"
	"encoding/json"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v5"

	v1 "compat/pkg/types/v1"
)

const (
	CompatibilityURL = "https://opencontainers.org/schema/image-compatibility"
)

//go:embed schema.json
var rawSchema string
var Schema *jsonschema.Schema

func init() {
	var err error
	Schema, err = jsonschema.CompileString(CompatibilityURL, rawSchema)
	if err != nil {
		panic(err)
	}
}

func OpenUserSchema(userSchemaPath string) ([]byte, error) {
	d, err := os.ReadFile(userSchemaPath)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func Validate(userSchema []byte) error {
	var v interface{}
	if err := json.Unmarshal(userSchema, &v); err != nil {
		return err
	}

	if err := Schema.Validate(v); err != nil {
		return err
	}

	return nil
}

func Load(userSchemaPath string) (*v1.ImageCompatibilitySchema, error) {
	d, err := OpenUserSchema(userSchemaPath)

	if err = Validate(d); err != nil {
		return nil, err
	}

	var i v1.ImageCompatibilitySchema
	if err = json.Unmarshal(d, &i); err != nil {
		return nil, err
	}

	return &i, nil
}
