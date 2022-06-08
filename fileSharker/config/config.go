package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	configObj   *Config
	fileTypeMap = make(map[string]struct{})
	isLoad      bool
)

func Load(path string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	if len(path) == 0 {
		path = "."
	}

	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	configObj = &Config{}
	err = viper.Unmarshal(configObj)

	nitFileTypeMap()

	isLoad = true
	return err
}

func nitFileTypeMap() {
	strList := strings.Split(configObj.FileTypes, ",")
	fileTypeMap = make(map[string]struct{}, len(strList))
	for _, str := range strList {
		if len(str) == 0 {
			continue
		}

		fileTypeMap[str] = struct{}{}
	}
}

func CheckFileType(fileType string) bool {
	_, exists := fileTypeMap[fileType]
	return exists
}

func GetConfig() Config {
	var temp = *configObj
	return temp
}

func GetPaths() []string {
	return configObj.FilePathList
}
