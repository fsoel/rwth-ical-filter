package ical

import "strings"

// Parameters:
//   ical    the full iCalendar data as a string
//   filters a slice of substrings; any event containing one of these substrings
//           will be omitted
//
// Returns:
//   A new iCalendar string with all matching VEVENT components removed.
func RemoveAllMatchingEvents(ical string, filters []string) string {
	const eventStart = "BEGIN:VEVENT"
	parts := strings.Split(ical, eventStart)
	newIcal := parts[0]

	for _, part := range parts[1:] {
		event := eventStart + part
		skip := false
		for _, f := range filters {
			if strings.Contains(event, f) {
				skip = true
				break
			}
		}
		if !skip {
			newIcal += event
		}
	}
	return newIcal
}
