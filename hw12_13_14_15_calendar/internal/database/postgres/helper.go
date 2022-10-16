package postgres

import (
	"strings"
)

func clearSQL(sql string) string {
	sql = strings.ReplaceAll(sql, "\n", "")
	sql = strings.ReplaceAll(sql, "\t", "")
	return sql
}
