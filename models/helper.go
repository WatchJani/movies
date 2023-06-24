package models

import (
	"fmt"
	"strings"
)

func GenerateCustomInput(query string, list []int, id int64) string {
	var builder strings.Builder
	builder.Grow(len(list) * 6) //6 -> number of length word -> "(?, ?), "

	for index := 0; index < len(list); index++ {
		builder.WriteString(fmt.Sprintf("(%d, %d), ", id, list[index]))
	}

	placeholders := builder.String()
	placeholders = placeholders[:len(placeholders)-2]

	insert := fmt.Sprintf("%s VALUES %s;", query, placeholders)

	return insert
}
