package dbModel

type ITable interface {
	GetTableName() string
	OneWhere(where string) (interface{}, bool)
	ListWhere(where string) ([]interface{}, bool)
}
