package config



type Config struct {
	DbName string
	DbIp   string
	DbPort int
	DbUser string
	DbPw   string

	FileTypes    string
	FilePathList []string
}


