package envloader

import "fmt"

//SampleConfig demonstrates the struct tags and options for your environment
type SampleConfig struct {
	NotRequired      string `env:"NOT_REQUIRED,required=false"`
	RequiredVariable string `env:"REQUIRED_VARIABLE,required=true"`
	CommentVariable  string `env:"INLINE_COMMENT,required=true"`
	NotSetVariable   string `env:"NOT_SET,required=false"`
}

//Example shows the output from parsing an env file into the config struct
func Example() {
	fmt.Println("testing parsing sample.env into SampleConfig")
	sampleConfig := SampleConfig{}
	ParseEnv(&sampleConfig, "sample.env")
	fmt.Println(sampleConfig)
	// Output:
	// testing parsing sample.env into SampleConfig
	// NOT_SET is empty var in map and is not required 'NotSetVariable' not set
	// {not_required_but_we_set_it_anyway required_env_variable inline_comment }
}
