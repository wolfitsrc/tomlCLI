package flags

import (
	"flag"
)

type CLIOptions struct {
	NewAPI bool
}

func ParseFlags() *CLIOptions {
	newAPI := flag.Bool("new-api", false, "Use the new API")
	flag.Parse()

	return &CLIOptions{
		NewAPI: *newAPI,
	}
}
