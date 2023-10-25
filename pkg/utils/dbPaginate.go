package utils

import "fmt"

type PageFilter struct {
	Page   int `query:"page"`
	Limit  int `query:"limit"`
}

func (p *PageFilter) WithQuery(query string) string {
	if p.Limit <= 0 || p.Page < 1 {
		return query
	}

	offset := (p.Page - 1) * p.Limit

	return fmt.Sprintf("%s OFFSET %d LIMIT %d", query, offset, p.Limit)
}
