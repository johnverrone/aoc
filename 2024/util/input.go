package util

import (
	"log"
	"os"
	"runtime"
	"strings"
)

type Opts struct {
	UseSample string
}

func ParseInput(opts *Opts) string {
	if opts != nil && opts.UseSample != "" {
		return opts.UseSample
	}

	_, file, _, _ := runtime.Caller(1)
	fp := strings.Split(file, "/")
	fp = fp[:len(fp)-1]

	data, err := os.ReadFile(strings.Join(fp, "/") + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
