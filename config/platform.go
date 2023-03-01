package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Instance *Config

type Config struct {
	AppName    string `yaml:"AppName"`    // 应用名
	Env        string `yaml:"Env"`        // 环境：prod、dev
	BaseUrl    string `yaml:"BaseUrl"`    // base url
	Port       string `yaml:"Port"`       // 端口
	LogFile    string `yaml:"LogFile"`    // 日志文件
	ShowSql    bool   `yaml:"ShowSql"`    // 是否显示日志
	StaticPath string `yaml:"StaticPath"` // 静态文件目录

	// 数据库配置
	DB struct {
		Dns          string `yaml:"Dns"`
		Host         string `yaml:"Host"`
		DbName       string `yaml:"DbName"`
		UserName     string `yaml:"UserName"`
		Password     string `yaml:"Password"`
		Prefix       string `yaml:"Prefix"`
		DbPort       int    `yaml:"DbPort"`
		Charset      string `yaml:"Charset"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"DB"`

	// 数据库配置
	MSSQLDB struct {
		Sqlserver          string `yaml:"Sqlserver"`
		Host               string `yaml:"Host"`
		DbName             string `yaml:"DbName"`
		UserName           string `yaml:"UserName"`
		Password           string `yaml:"Password"`
		Prefix             string `yaml:"Prefix"`
		DbPort             int    `yaml:"DbPort"`
		SetMaxIdleConns    int    `yaml:"SetMaxIdleConns"`
		SetMaxOpenConns    int    `yaml:"SetMaxOpenConns"`
		SetConnMaxLifetime int    `yaml:"SetConnMaxLifetime"`
	} `yaml:"MSSQLDB"`

	// 数据库类型
	DatabaseType struct {
		UseDbType string `yaml:"UseDbType"`
	} `yaml:"DatabaseType"`

	// 数据库日志文件
	DbLogFile struct {
		LogFile string `yaml:"LogFile"`
	} `yaml:"DbLogFile"`

	// 阿里云oss配置
	Uploader struct {
		Enable    string `yaml:"Enable"`
		AliyunOss struct {
			Host          string `yaml:"Host"`
			Bucket        string `yaml:"Bucket"`
			Endpoint      string `yaml:"Endpoint"`
			AccessId      string `yaml:"AccessId"`
			AccessSecret  string `yaml:"AccessSecret"`
			StyleSplitter string `yaml:"StyleSplitter"`
			StyleAvatar   string `yaml:"StyleAvatar"`
			StylePreview  string `yaml:"StylePreview"`
			StyleSmall    string `yaml:"StyleSmall"`
			StyleDetail   string `yaml:"StyleDetail"`
		} `yaml:"AliyunOss"`
		Local struct {
			Host string `yaml:"Host"`
			Path string `yaml:"Path"`
		} `yaml:"Local"`
	} `yaml:"Uploader"`

	// smtp
	Smtp struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		SSL      bool   `yaml:"SSL"`
	} `yaml:"Smtp"`

	// es
	Es struct {
		Url   string `yaml:"Url"`
		Index string `yaml:"Index"`
	} `yaml:"Es"`
}

func Init(filename string) *Config {
	Instance = &Config{}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		logrus.Error(err)
	}
	return Instance
}
