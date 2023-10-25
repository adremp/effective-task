package utils

import (
	"fmt"
	"strings"
)

type SqlFilter struct {
	Data     map[string]string
	FindKeys []string
}

func (s *SqlFilter) CreateQuery() string {
	var query []string
	builder := strings.Builder{}
	for _, key := range s.FindKeys {
		if value, ok := s.Data[key]; ok {
			builder.WriteString(key)
			builder.WriteString(" = ")
			builder.WriteString(value)
			query = append(query, builder.String())
			builder.Reset()
		}
	}
	return strings.Join(query, " AND ")
}

func FilterMap(data map[string]string, keys []string) map[string]string {
	var filtered map[string]string
	for _, key := range keys {
		if value, ok := data[key]; ok {
			filtered[key] = value
		}
	}
	return filtered
}

func ParseMinMaxMaybeQuery(key, value string) string {
	var ret []string
	if strings.Contains(value, "-") {
		age := strings.Split(value, "-")

		if age[0] == "" {
			ret = append(ret, fmt.Sprintf("%s <= '%v'", key, age[1]))
		}
		if age[1] == "" {
			ret = append(ret, fmt.Sprintf("%s >= '%v'", key, age[0]))
		}
		return strings.Join(ret, " AND ")
	}
	return fmt.Sprintf("%s = '%v'", key, value)
}
