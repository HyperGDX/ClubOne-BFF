package config

type Backend struct {
	BaseApi  string `mapstructure:"base-api-url" json:"base-api-url" yaml:"base-api-url"`
	ForumApi string `mapstructure:"forum-api-url" json:"forum-api-url" yaml:"forum-api-url"`
	AuthApi  string `mapstructure:"auth-api-url" json:"auth-api-url" yaml:"auth-api-url"`
	ThirdApi string `mapstructure:"third-api-url" json:"third-api-url" yaml:"third-api-url"`
}
