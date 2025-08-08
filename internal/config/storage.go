package config

type Credentials struct {
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
}

type Storage struct {
	Provider       string      `mapstructure:"provider"`
	Region         string      `mapstructure:"region"`
	Bucket         string      `mapstructure:"bucket"`
	Endpoint       string      `mapstructure:"endpoint"`
	ForcePathStyle bool        `mapstructure:"force_path_style"`
	Credentials    Credentials `mapstructure:"credentials"`
}
