package sonosAPI

import (
	"fmt"
	"strconv"
	"strings"
)

func timeValueFormatter(value int) string {
	seconds := value % 60
	value -= seconds * 60
	value = value / 60
	minutes := value % 60
	value -= minutes * 60
	hours := minutes / 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func timeValueParser(value string) int {
	parts := strings.Split(value, ":")
	parsedValue := 0

	multiplier := 1
	for i := len(parts) - 1; i >= 0 ; i-- {
		parsedPart, err := strconv.Atoi(parts[i])
		if err != nil  {
			fmt.Printf("Could not parse part %d of %s ('%s')", i, value, parts[i])
			continue
		}
		parsedValue += parsedPart * multiplier
		multiplier *= 60
	}

	return parsedValue
}