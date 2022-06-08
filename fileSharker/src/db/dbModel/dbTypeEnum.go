package dbModel

type DbTypeEnum int

const (
	DbTypeEnum_Mysql   DbTypeEnum = 1
	DbTypeEnum_Mongodb DbTypeEnum = 2
)

// 全部属性类型
var dbTypeEnumMap = map[int32]struct{}{
	int32(DbTypeEnum_Mysql):   {},
	int32(DbTypeEnum_Mongodb): {},
}

func Verify(dbType int32) bool {
	_, exists := dbTypeEnumMap[dbType]
	return exists
}
