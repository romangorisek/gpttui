package text

import "strings"

func ToMaxWidth(text string, width int) string {
	if width <= 0 {
		return text
	}
	var resultLines []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if len(line) <= width {
			resultLines = append(resultLines, line)
			continue
		}

		tempLine := line
		for len(tempLine) > width {
			// Find the last space within the 'width' limit
			breakPoint := width
			segment := tempLine[:width]
			lastSpace := strings.LastIndex(segment, " ")

			if lastSpace > 0 { // Ensure space is not the very first character and exists
				breakPoint = lastSpace
			}

			resultLines = append(resultLines, tempLine[:breakPoint])
			// If we broke at a space, skip the space in the next line segment
			if lastSpace > 0 {
				tempLine = tempLine[breakPoint+1:]
			} else {
				tempLine = tempLine[breakPoint:]
			}
		}
		if len(tempLine) > 0 {
			resultLines = append(resultLines, tempLine)
		}
	}
	return strings.Join(resultLines, "\n")
}
