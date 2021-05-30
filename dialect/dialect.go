package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(value reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

//注册实例
func RegisterDialect(name string, dialect Dialect)  {
	dialectMap[name] = dialect
}

//获取实例
func GetDialect(name string) (dialect Dialect, ok bool)  {
	dialect, ok = dialectMap[name]
	return
}
