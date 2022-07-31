package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type SqlConfig struct {
	Host            string `yaml:"sqlhost"`            // 数据库主机地址以及端口
	User            string `yaml:"sqluser"`            // 数据库用户名
	Password        string `yaml:"sqlpwd"`             // 数据库密码
	Dbname          string `yaml:"sqldbname"`          // 数据库名
	MaxOpenConns    int    `yaml:"sqlMaxOpenConns"`    // 最大打开连接数
	MaxIdleConns    int    `yaml:"sqlMaxIdleConns"`    // 最大空闲连接数
	ConnMaxLifetime int    `yaml:"sqlConnMaxLifetime"` // 连接最大存活时间
}

type RedisConfig struct {
	Host string `yaml:"redishost"` // redis主机地址
	Port int    `yaml:"redisport"` // redis端口
	Pwd  string `yaml:"redispwd"`  // redis密码
}

func GetSqlConnConfig() string {
	var sqlConfig SqlConfig
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &sqlConfig)
	dns := sqlConfig.User + ":" + sqlConfig.Password + "@tcp(" + sqlConfig.Host + ")/" + sqlConfig.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	return dns
}

func GetSqlConfig() *SqlConfig {
	var sqlConfig SqlConfig
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &sqlConfig)
	return &sqlConfig
}
