package config

type ConfigRoot struct {
	DbConfig     DbConfig     `yaml:"dbConfig"`
	ApiConfig    ApiConfig    `yaml:"apiConfig"`
	ServerConfig ServerConfig `yaml:"serverConfig"`
	RedisConfig  RedisConfig  `yaml:"redisConfig"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Replica  string `yaml:"replica"`
}

type ApiConfig struct {
	Token   string `yaml:"token"`
	BaseUri string `yaml:"base_uri"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}
