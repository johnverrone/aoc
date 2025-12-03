package util

import (
	"log"
	"os"
	"runtime"
	"strings"
)

func ParseInput(sample string) string {
	if sample != "" {
		return strings.TrimSpace(sample)
	}

	_, file, _, _ := runtime.Caller(1)
	fp := strings.Split(file, "/")
	fp = fp[:len(fp)-1]

	data, err := os.ReadFile(strings.Join(fp, "/") + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(data))
}
