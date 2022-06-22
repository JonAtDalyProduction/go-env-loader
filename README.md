# go-env-loader

### download the package
```shell
go get github.com/JonAtDalyProduction/go-env-loader
```
### run the demo

```shell
 go run examples/main.go --env=examples/sample.env
```

## load a .env file into a config struct with tags

```go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	envloader "github.com/JonAtDalyProduction/go-env-loader"
)

type SampleConfig struct {
	SampleVariable string `env:"SAMPLE_VARIABLE,required=true"`
	NotRequired    string `env:"NOT_REQUIRED,required=false"`
}

func main() {
	var config SampleConfig
	var envFile string
	flag.StringVar(&envFile, "env", ".env", "sets the .env filename")
	flag.Parse()
	envloader.ParseEnv(config, envFile)
	bs, _ := json.MarshalIndent(config, "", "    ")
	fmt.Printf("CONFIG SET:\n%+v\n", string(bs))
}


```
