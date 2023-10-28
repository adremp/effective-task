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

func ParseMinMaxMaybeQuery(startArgIdx int, key, value string) (string, []string) {
	argIdx := startArgIdx
	var queryRet []string
	var valuesRet []string
	if strings.Contains(value, "-") {
		valueArr := strings.Split(value, "-")

		if valueArr[0] != "" {
			argIdx++
			queryRet = append(queryRet, fmt.Sprintf("%s <= $%v", key, argIdx))
			valuesRet = append(valuesRet, valueArr[0])
		}
		if valueArr[1] != "" {
			queryRet = append(queryRet, fmt.Sprintf("%s >= '%v'", key, valueArr[0]))
			valueArr = append(valueArr, valueArr[1])
		}
		if len(valueArr) == 0 {
			return "", []string{}
		}
		return strings.Join(queryRet, " AND "), valuesRet
	}
	return fmt.Sprintf("%s = $%v", key, argIdx), []string{value}
}
