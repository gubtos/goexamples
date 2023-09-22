package main

import (
	"fmt"

	"github.com/gubtos/pattern/structtags"
)

func main() {
	def := structtags.DefaultClientConfig()
	fmt.Printf("%+v", def)
}

func StructTagsComparison() {
	// mode 1
	structtags.NewClient(structtags.ConfigFromEnv())

	// mode 2
	cfg := structtags.DefaultClientConfig()
	cfg.URL = "https://gopherconbr.org/"
	structtags.NewClient(structtags.Config(cfg))

	//mode 3
	serviceConfig := FromExternalConfigLib()
	structtags.NewClient(structtags.Config(serviceConfig.ClientConfig))
}

type ServiceConfig struct {
	structtags.ClientConfig `key:"clientConfig"`
}

func FromExternalConfigLib() ServiceConfig {
	return ServiceConfig{}
}
