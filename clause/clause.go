package clause

import (
	"strings"
)

type Clause struct {
	sql map[Type]string
	sqlVars map[Type][]interface{}
}

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	WHERE
	ORDERBY
	LIMIT
	UPDATE
	DELETE
	COUNT
)

// 根据Type生成对应的sql语句
func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql, lar := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = lar
}

// 根据传入的type顺序，构造出最终的sql
func (c *Clause) Build(orders ...Type) (string, []interface{})   {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql,ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}