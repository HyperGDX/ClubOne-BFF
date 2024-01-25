package config

type Backend struct {
	BaseApi string `mapstructure:"base-api-url" json:"base-api-url" yaml:"base-api-url"`
	ForumApi string `mapstructure:"forum-api-url" json:"forum-api-url" yaml:"forum-api-url"`
}
