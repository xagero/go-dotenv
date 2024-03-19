package dotenv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/xagero/go-helper/helper"
)

const (
	commentSeparator = "#"
	kvSeparator      = "="
)

func ReadFromFile(filename string) {
	file, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
	}
	defer func() {
		if e = file.Close(); e != nil {
			log.Fatal(e)
		}
	}()

	scanner := bufio.NewScanner(file)

	var text string
	var result []string

	for scanner.Scan() {
		text = scanner.Text()
		if helper.IsBlank(text) {
			continue
		}
		if -1 == strings.Index(text, commentSeparator) {
			defineEnv(text)
		} else {
			result = strings.Split(text, commentSeparator)
			defineEnv(result[0])
		}
	}
}

func parseQuotedValue(kv string) (string, error) {
	quoteRegex := regexp.MustCompile(`^"(.*?)"$`)
	parts := quoteRegex.FindStringSubmatch(kv)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid quoted value: %s", kv)
	}
	return parts[1], nil
}

func defineEnv(s string) {
	if -1 != strings.Index(s, kvSeparator) {
		parts := strings.Split(s, kvSeparator)
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid key-value pair: %s", s))
		}

		k := parts[0]
		var v string

		if strings.HasPrefix(parts[1], `"`) {
			v, _ = parseQuotedValue(parts[1])
		} else {
			v = parts[1]
		}

		// fmt.Printf("Key = %s, value = %s\n", k, v)
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
}
