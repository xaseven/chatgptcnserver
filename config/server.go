package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Proxy  Proxy  `mapstructure:"proxy" json:"proxy" yaml:"proxy"`
}
