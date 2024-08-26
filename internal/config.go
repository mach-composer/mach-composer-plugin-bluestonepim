package internal

type BluestonePimGlobalConfig struct {
	ClientID     string  `mapstructure:"client_id"`
	ClientSecret string  `mapstructure:"client_secret"`
	AuthUrl      *string `mapstructure:"auth_url"`
	ApiUrl       *string `mapstructure:"api_url"`
}

type BluestonePimSiteConfig struct {
	ClientID     string  `mapstructure:"client_id"`
	ClientSecret string  `mapstructure:"client_secret"`
	AuthUrl      *string `mapstructure:"auth_url"`
	ApiUrl       *string `mapstructure:"api_url"`
}
