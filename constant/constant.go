package constant

const (
	ErrDriverName = "目前仅支持pg与mysql数据库"
)

const (
	MysqlDriverName    = "mysql"
	PostgresDriverName = "postgres"
)

const (
	AuthRuleTypeFirst = iota + 1
	AuthRuleTypeSecond
	AuthRuleTypeThree
)

const (
	UnDelete = iota
	IsDelete
)

const (
	DateLayout     = "2006-01-02"
	DatetimeLayout = "2006-01-02 15:04:05"
)
