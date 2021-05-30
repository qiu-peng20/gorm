package clause

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
)

func (c *Clause) Set(name Type, vars ...interface{}) {

}
