package linux

import (
	"encoding/json"
)

const (
	oneOf = "oneof"
)

type Schema struct {
	OneOf    map[string]Subjects
	Subjects Subjects
}

type Subjects map[string]interface{}

func (s *Schema) UnmarshalJSON(data []byte) error {
	md := make(map[string]interface{})
	err := json.Unmarshal(data, &md)
	if err != nil {
		return err
	}
	if _, ok := md[oneOf]; ok {
		s.OneOf = make(map[string]Subjects)
		for group, id := range md[oneOf].(map[string]interface{}) {
			if _, ok = s.OneOf[group]; !ok {
				s.OneOf[group] = make(Subjects)
			}
			for subject, input := range id.(map[string]interface{}) {
				s.OneOf[group][subject] = input
			}
		}
	} else {
		s.Subjects = md
	}
	return nil
}
