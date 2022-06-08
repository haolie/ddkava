package config

type Config struct {
	DbName       string   // 数据库名称
	DbIp         string   // 数据库Ip
	DbPort       int      // 数据库端口
	DbUser       string   // 数据库用户名
	DbPw         string   // 数据库pw
	SaveRoot     string   // 文件保存根目录
	FileTypes    string   // 文件类型
	FilePathList []string // 扫描文件目录列表
}
