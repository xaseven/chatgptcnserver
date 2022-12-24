package config

type Proxy struct {
	Enable       bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Info         string `mapstructure:"info" json:"info" yaml:"info"`
	Chatgptmodel string `mapstructure:"chatgptmodel" json:"chatgptmodel" yaml:"chatgptmodel"`
	Chatgpttoken string `mapstructure:"chatgpttoken" json:"chatgpttoken" yaml:"chatgpttoken"`
}
