package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/lights-T/rbac/config"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type pg struct {
	DbAdapter
}

func (s *pg) NewDriver(address string) (*goqu.Database, error) {
	var self *goqu.Database
	if len(address) == 0 {
		return self, errors.New("Database url is empty ")
	}
	dialect := goqu.Dialect("postgres")
	pgDb, err := sql.Open("postgres", address)
	if err != nil {
		return self, err
	}
	self = dialect.DB(pgDb)

	l := zerolog.New(os.Stderr)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	self.Logger(&l)
	return self, nil
}

func (s *pg) CreateDatabase(name string) error {
	var err error
	if _, err = s.Engine.Exec(fmt.Sprintf("CREATE DATABASE %s", name)); err != nil {
		return err
	}
	return nil
}

func (s *pg) CreateTable(tableName ...string) error {
	var err error
	if len(tableName) != 3 {
		return errors.New("invalid parameter: tableName")
	}
	for k, name := range tableName {
		switch k {
		case 0:
			if _, err = s.Engine.Exec(fmt.Sprintf(config.PgRbacAuthRule, name, name, name, name, name, name, name, name, name, name, name, name, name)); err != nil {
				return err
			}
		case 1:
			if _, err = s.Engine.Exec(fmt.Sprintf(config.PgRbacAuthGroup, name, name, name, name, name, name, name, name)); err != nil {
				return err
			}
		case 2:
			if _, err = s.Engine.Exec(fmt.Sprintf(config.PgRbacAuthGroupAccess, name, name, name, name, name, name, name, name)); err != nil {
				return err
			}
		}
	}
	return nil
}
