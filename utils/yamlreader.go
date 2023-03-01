package utils

import (
	"os"

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
	Host     string `yaml:"redishost"`     // redis主机地址以及端口
	Pwd      string `yaml:"redispwd"`      // redis密码
	PoolSize int    `yaml:"redisPoolSize"` // redis连接池大小
}

type SecurityConfig struct {
	RSAUpdateLifecycle int `yaml:"rsaUpdateLifecycle"` // RSA密钥更新周期（单位：分钟）
}

type CosConfig struct {
	Bucket    string `yaml:"cosBucket"`
	Region    string `yaml:"cosRegion"`
	Appid     string `yaml:"cosAppid"`
	SecretId  string `yaml:"cosSecretId"`
	SecretKey string `yaml:"cosSecretKey"`
	Domain    string `yaml:"cosDomain"`
}

func GetSqlConnConfigStr() string {
	var sqlConfig SqlConfig
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &sqlConfig)
	dns := sqlConfig.User + ":" + sqlConfig.Password + "@tcp(" + sqlConfig.Host + ")/" + sqlConfig.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	return dns
}

func GetSqlConfig() *SqlConfig {
	var sqlConfig SqlConfig
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &sqlConfig)
	return &sqlConfig
}

func GetRedisConfig() *RedisConfig {
	var redisConfig RedisConfig
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &redisConfig)
	return &redisConfig
}

func GetSecurityConfig() *SecurityConfig {
	var securityConfig SecurityConfig
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &securityConfig)
	return &securityConfig
}

func GetLogConfig() *Logger {
	var logConfig Logger
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &logConfig)
	return &logConfig
}

func GetCosConfig() *CosConfig {
	var cosConfig CosConfig
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &cosConfig)
	return &cosConfig
}
