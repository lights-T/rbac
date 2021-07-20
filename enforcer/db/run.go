package db

import (
	"github.com/doug-martin/goqu/v9"
	"xorm.io/xorm"
)

type Db interface {
	Mysql() Action
	Pg() Action
}

type Action interface {
	//CreateTable 创建数据表 tableNames注意顺序 第一个auth_rule 第二个auth_group 第三个auth_group_access
	CreateTable(tableNames ...string) error
	CreateDatabase(name string) error
	NewDriver(address string) (*goqu.Database, error)
}

type DbAdapter struct {
	Engine *xorm.Engine
}

func NewDbAdapter() *DbAdapter {
	return &DbAdapter{}
}

func (s *DbAdapter) Mysql() Action {
	return &mysql{DbAdapter{Engine: s.Engine}}
}

func (s *DbAdapter) Pg() Action {
	return &pg{DbAdapter{Engine: s.Engine}}
}
