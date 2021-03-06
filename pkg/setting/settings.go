package setting

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Organ      string `yaml:"organ"`
	HttpPort   int    `yaml:"httpPort"`
	ServerMode string `yaml:"serverMode"`
	JwtSecret  string `yaml:"jwtSecret"`
	XxteaKey   string `yaml:"xxteaKey"`

	MysqlHost     string `yaml:"mysqlHost"`
	MysqlPort     int    `yaml:"mysqlPort"`
	MysqlUser     string `yaml:"mysqlUser"`
	MysqlPassword string `yaml:"mysqlPassword"`
	MysqlDbName   string `yaml:"mysqlDbName"`

	RedisUrl         string `yaml:"redisUrl"`
	RedisPassword    string `yaml:"redisPassword"`
	RedisPoolSize    int    `yaml:"redisPoolSize"`
	RedisPoolMinIdel int    `yaml:"redisPoolMinIdel"`

	LoggerMaxSize    int `yaml:"loggerMaxSize"`
	LoggerMaxAge     int `yaml:"loggerMaxAge"`
	LoggerMaxBackups int `yaml:"loggerMaxBackups"`

	GormLoggerFileName    string `yaml:"gormLoggerFileName"`
	ParamsLoggerFileName  string `yaml:"paramsLoggerFileName"`
	MonitorLoggerFileName string `yaml:"monitorLoggerFileName"`
}

var instance *Settings

var once sync.Once

func GetInstance() *Settings {
	once.Do(func() {
		instance = loadSettings()
	})

	return instance
}

// Setup initialize the configuration instance
func loadSettings() *Settings {
	settings := &Settings{}
	yamlFile, err := ioutil.ReadFile("conf/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, settings)
	if err != nil {
		fmt.Println(err.Error())
	}

	return settings
}
