package validate

import (
	"encoding/json"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"compat/schema"
)

const (
	SchemaCmdName = "validate-schema"
)

func RunSchemaCmd(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("image compatibility schema file path is required")
	}

	s, err := schema.OpenUserSchema(args[0])
	if err != nil {
		return err
	}

	if err = schema.Validate(s); err != nil {
		if ve, ok := err.(*jsonschema.ValidationError); ok {
			b, _ := json.MarshalIndent(ve.DetailedOutput(), "", "  ")
			return fmt.Errorf("schema is not valid:\n%s", string(b))
		} else {
			return fmt.Errorf("validation failed: %w", err)
		}
	} else {
		fmt.Printf("Schema %q is valid!\n", args[0])
	}

	return nil
}
