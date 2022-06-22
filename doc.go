// Package envloader
// uses standard .env file format and supports non-inline comments only
// currently only string types are supported for the struct
//  type DemoConfig struct {
//   notRequiredVar string `env:"NOT_REQUIRED_VAR,required=false"`
//   requiredVar string `env:"REQUIRED_VAR,required=true"`
//  }
// pass this struct to ParseEnv along with the .env filename
package envloader
