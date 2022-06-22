// Package go_env_loader
// uses standard .env file format and supports non-inline comments only
//  type DemoConfig struct {
//   notRequiredVar string `env:"NOT_REQUIRED_VAR,required=false"`
//   requiredVar string `env:"REQUIRED_VAR,required=true"`
//  }
// pass this struct to ParseEnv along with the .env filename
package go_env_loader
