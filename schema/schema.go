package schema

import (
	"go/ast"
	"gorm/dialect"
	"reflect"
)

//定义列
type Field struct {
	Name string //列名
	Type string //列类型
	Tag  string //约束条件
}

// 定义model类别
type Schema struct {
	Model     interface{}
	Name      string
	Fields    []*Field
	FieldName []string          // 记录所有字段名
	FieldMap  map[string]*Field // 方便直接找到该字段
}

func (s *Schema) GetField(name string) *Field {
	return s.FieldMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		FieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v,ok := p.Tag.Lookup("gorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldName = append(schema.FieldName, p.Name)
			schema.FieldMap[p.Name] = field
		}
	}
	return schema
}

func (s *Schema) RecordValues(dest interface{}) []interface{}  {
	destValues := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValues.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
