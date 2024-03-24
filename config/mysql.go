package config

type MySQLConfig struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	Schema       string `yaml:"schema" json:"schema"`
	DialTimeout  int    `yaml:"dial_timeout" json:"dial_timeout"`
	ReadTimeout  int    `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout" json:"write_timeout"`
	MaxIdleConns int    `yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns" json:"max_open_conns"`
	MaxLifetime  int    `yaml:"max_life_time" json:"max_life_time"`
}
