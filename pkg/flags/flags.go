package flags

import (
	"flag"
	"strings"
)

type FlagOptions struct {
	CreateDir bool
	AddAPI    APIkey
}

type APIkey struct {
	Name string
	Key  string
}

func ParseFlags() *FlagOptions {
	var Flags FlagOptions
	var API string

	CreateFlag := flag.Bool("c", false, "Create all subdirectoried and files")
	flag.StringVar(&API, "api", "", "API name")

	flag.Parse()

	Flags.CreateDir = *CreateFlag

	if API != "" {
		split := strings.Split(API, "/")
		Flags.AddAPI.Name = split[0]
		Flags.AddAPI.Key = split[1]
	}
	return &Flags
}
