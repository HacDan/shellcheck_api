package storage

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func FindAllCodes(s string) []string {
	re := regexp.MustCompile("SC\\d\\d\\d\\d")
	return re.FindAllString(s, -1)
}

func ParseSCFile() map[string]string {
	f, err := os.Open("sc.md")
	if err != nil {
		panic(err)
	}

	shellCheckCodes := make(map[string]string)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := parseSCLine(scanner.Text())
		shellCheckCodes[line[0]] = line[1]
	}

	return shellCheckCodes
}

func parseSCLine(s string) []string {
	split := strings.Split(s, " ")

	code := split[0]
	code = strings.Replace(code, "[", "", -1)
	code = strings.Replace(code, "]", "", -1)

	return []string{code, strings.Join(split[1:], " ")}
}
