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
```env
NOT_REQUIRED=not_required_but_we_set_it_anyway
REQUIRED_VARIABLE=required_env_variable
INLINE_COMMENT=inline_comment # ths is a comment
# normal comment
```
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
	InlineComment  string `env:"INLINE_COMMENT,required=true"`
}

func main() {
	var config SampleConfig
	var envFile string
	flag.StringVar(&envFile, "env", ".env", "sets the .env filename")
	flag.Parse()
	# pass a pointer to your config struct
	envloader.ParseEnv(&config, envFile)
	bs, _ := json.MarshalIndent(config, "", "    ")
	fmt.Printf("CONFIG SET:\n%+v\n", string(bs))
}
```
