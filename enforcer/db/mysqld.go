package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/lights-T/rbac/config"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lights-T/lib-go/database/db"
	"github.com/rs/zerolog"
)

type mysql struct {
	DbAdapter
}

func (s *mysql) NewDriver(address string) (*goqu.Database, error) {
	var self *goqu.Database
	if len(address) == 0 {
		return self, errors.New("Database url is empty ")
	}

	var mcs []*db.Conf
	mcs = append(mcs, &db.Conf{
		InstanceName: "rbac",
		DriverName:   "mysql",
		DataSource:   address,
	})

	nDb, err := db.New(mcs)
	if err != nil {
		return self, err
	}

	if self = nDb.GetInstance(""); self == nil {
		return self, errors.New("database init fail ")
	}
	l := zerolog.New(os.Stderr)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	self.Logger(&l)
	return self, nil
}

func (s *mysql) CreateDatabase(name string) error {
	var err error
	if _, err = s.Engine.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)); err != nil {
		return err
	}
	return nil
}

func (s *mysql) CreateTable(tableName ...string) error {
	var err error
	if len(tableName) != 3 {
		return errors.New("invalid parameter: tableName")
	}
	for k, name := range tableName {
		switch k {
		case 0:
			if _, err = s.Engine.Exec(fmt.Sprintf(config.MysqlRbacAuthRule, name)); err != nil {
				return err
			}
		case 1:
			if _, err = s.Engine.Exec(fmt.Sprintf(config.MysqlRbacAuthGroup, name)); err != nil {
				return err
			}
		case 2:
			if _, err = s.Engine.Exec(fmt.Sprintf(config.MysqlRbacAuthGroupAccess, name)); err != nil {
				return err
			}
		}
	}
	return nil
}
