package utilities

func FlipLine(lines []string) []string {
	var reversedLines []string
	for i := range lines {
		reversedLines = append(reversedLines, lines[len(lines)-1-i])
	}
	return reversedLines
}
