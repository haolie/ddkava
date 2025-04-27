package Config

import (
	"strings"

	"github.com/spf13/viper"
	"lyh/ddkava/fileSharker/Common/Def"
)

var (
	configMap      = make(map[string]interface{}, 8)
	isLoadFinished = false
)

func UpdateConfig(k string, v interface{}) {
	if isLoadFinished {
		panic("")
	}

	configMap[k] = v
}

/*
返回指定类型的配置
*/
func GetTypeValue[T int32 | int64 | string | bool | []string](k string) (v T, exists bool) {
	temp, exists := configMap[k]
	if exists {
		v = temp.(T)
	}

	return
}

/*
返回指定配置
*/
func GetValue(k string) (v interface{}, exists bool) {
	v, exists = configMap[k]
	return v, exists
}

/*
加载配置文件
*/
func Load(path string) error {
	viper.SetConfigName(Def.Config_FileName)
	viper.SetConfigType("toml")
	if len(path) == 0 {
		path = "."
	}

	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	for k, v := range viper.AllSettings() {
		configMap[k] = v
	}

	err = convert()

	return err
}

/*
转换配置
*/
func convert() error {

	typeStr, exists := GetTypeValue[string](Def.Config_File_Types)
	if exists {
		tempList := strings.Split(typeStr, ",")
		tempMap := make(map[string]struct{})
		for _, v := range tempList {
			tempMap[v] = struct{}{}
		}

		UpdateConfig(Def.Config_File_TypeMap, tempMap)
	}

	return nil
}
