package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nzai/stockrecorder/io"
)

const (
	configFile     = "project.json"
	defaultDataDir = "data"
)

type Config struct {
	RootDir  string
	MongoUrl string
}

//	当前系统配置
var configValue *Config = nil

//	设置配置文件
func SetRootDir(root string) error {

	//	构造配置文件路径
	filePath := filepath.Join(root, configFile)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return err
	}

	//	读取文件
	buffer, err := io.ReadAllBytes(filePath)
	if err != nil {
		return err
	}

	//	解析配置项
	configValue = &Config{}
	err = json.Unmarshal(buffer, configValue)
	if err != nil {
		return err
	}

	if configValue == nil {
		return fmt.Errorf("配置文件错误")
	}

	return nil
}

//	获取当前系统配置
func Get() *Config {
	return configValue
}
