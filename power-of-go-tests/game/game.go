package game

import (
	"fmt"
	"strings"
)

func ListItems(items []string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("You can see %s here.", items[0])
	case 2:
		return fmt.Sprintf("You can see here %s and %s.", items[0], items[1])
	default:
		sb := strings.Builder{}
		sb.WriteString("You can see here ")
		for i, item := range items {
			isNotLastItem := len(items) != i+1
			if isNotLastItem {
				sb.WriteString(fmt.Sprintf("%s, ", item))
			} else {
				sb.WriteString(fmt.Sprintf("and %s.", item))
			}
		}
		return sb.String()
	}
}
