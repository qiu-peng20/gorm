package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql struct {}

func init()  {
	RegisterDialect("mysql",&mysql{})
}

func (m mysql)DataTypeOf(value reflect.Value) string  {
	switch value.Kind() {
	case reflect.Int, reflect.Int8,reflect.Int16, reflect.Int32,
	reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32:
		return "integer"
	case reflect.Int64,reflect.Uint64:
		return "bigint"
	case reflect.Float32,reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := value.Interface().(time.Time);ok {
			return "datatime"
		}
	}
	panic(fmt.Sprintf("sql type %s(%s)",value.Type().Name(), value.Kind()))
}

func (m mysql)TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT name FROM gorm WHERE type=`table` and name=?",args
}