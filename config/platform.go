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
		Charset      int    `yaml:"Charset"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"DB"`

	// Github
	Github struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"Github"`

	// QQ登录
	QQConnect struct {
		AppId  string `yaml:"AppId"`
		AppKey string `yaml:"AppKey"`
	} `yaml:"QQConnect"`

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

	// 百度ai
	BaiduAi struct {
		ApiKey    string `yaml:"ApiKey"`
		SecretKey string `yaml:"SecretKey"`
	} `yaml:"BaiduAi"`

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
