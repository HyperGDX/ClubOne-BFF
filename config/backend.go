package config

type Backend struct {
	BaseApi string `mapstructure:"base-api-url" json:"base-api-url" yaml:"base-api-url"`
}
