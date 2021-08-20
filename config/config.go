package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host    string  `yaml:"host"`
	Port    int32   `yaml:"port"`
	Test    string  `yaml:"test"`
	Package Package `yaml:"package"`
}

type Package struct {
	OpenCheck     bool  `yaml:"open-check"`
	BodyLenOffset int32 `yaml:"body-len-offset"`
	PackageMax    int32 `yaml:"package-max"`
	HeaderLen     int32 `yaml:"header-len"`
}

type Mysql struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	Dbname        string `yaml:"dbname"`
	Charset       string `yaml:"charset"`
	ActiveMax     int    `yaml:"active-max"`
	ConnectionMax int    `yaml:"connection-max"`
	Dev           string `yaml:"dev"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Dev      string `yaml:"dev"`
}

type MysqlConf struct {
	Master []Mysql
	Slave  []Mysql
}

type RedisConf struct {
	Master []Redis
	Slave  []Redis
}

type Logger struct {
	Level  string `yaml:"level"`
	LogDir string `yaml:"log-dir"`
}

type Config struct {
	Server     Server     `yaml:"server"`
	Mysql      MysqlConf  `yaml:"mysql"`
	Logger     Logger     `yaml:"logger"`
	Redis      RedisConf  `yaml:"redis"`
	ClickHouse ClickHouse `yaml:"clickhouse"`
}

func LoadConfigBy(path string, t interface{}) error {
	if !IsFile(path) {
		return fmt.Errorf("path[%s] is not exists", path)
	}

	conf := reflect.ValueOf(t)
	if conf.Kind() != reflect.Ptr {
		return fmt.Errorf("params is not ptr")
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, t)
}

func LoadConfig(path string) (*Config, error) {
	conf := &Config{}
	err := LoadConfigBy(path, conf)
	if err != nil {
		return nil, err
	}

	return conf, err
}

func IsFile(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !p.IsDir()
}
