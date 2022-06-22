package main

import (
	"encoding/json"
	"flag"
	"fmt"
	envloader "github.com/jonatdalyproduction/go-env-loader"
)

type SampleConfig struct {
	SampleVariable string `env:"SAMPLE_VARIABLE,required=true"`
	NotRequired    string `env:"NOT_REQUIRED,required=false"`
}

func main() {
	config := SampleConfig{}
	var envFile string
	flag.StringVar(&envFile, "env", ".env", "sets the .env filename")
	flag.Parse()
	envloader.ParseEnv(&config, envFile)
	bs, _ := json.MarshalIndent(config, "", "    ")
	fmt.Printf("CONFIG SET:\n%+v\n", string(bs))
}
