package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	configObj   *Config
	fileTypeMap = make(map[string]bool)
	isLoad      bool
)

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	configObj = &Config{}
	err = viper.Unmarshal(configObj)

	isLoad = true
	return err
}

func CheckFileType(fileType string) bool {
	if !isLoad && len(configObj.FileTypes) > 1 {
		strList := strings.Split(configObj.FileTypes, ",")
		fileTypeMap = make(map[string]bool, len(strList))
		for _, str := range strList {
			if len(str) == 0 {
				continue
			}

			fileTypeMap[str] = true
		}
	}

	return len(fileTypeMap) == 0 || fileTypeMap[fileType]
}

func GetPaths() []string {
	return configObj.FilePathList
}
