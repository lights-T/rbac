package enforcer

import (
	"errors"
	"github.com/lights-T/rbac/enforcer/db"
	"runtime"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	redislib "github.com/lights-T/lib-go/cache/redis"
	"github.com/lights-T/lib-go/logger"
	"github.com/lights-T/rbac/constant"
	"xorm.io/xorm"
)

type Option interface {
	apply(*Enforcer)
}

type funcOption struct {
	f func(*Enforcer)
}

func (s *funcOption) apply(do *Enforcer) {
	s.f(do)
}

func newFuncOption(req func(*Enforcer)) *funcOption {
	return &funcOption{
		f: req,
	}
}

// WithDbConfig  set dataBase config
func WithDbConfig(driverName, dataSourceName string, tableNames ...string) Option {
	return newFuncOption(func(o *Enforcer) {
		o.driverName = driverName
		o.dataSourceName = dataSourceName
		o.tableNames = tableNames
	})
}

// WithRedisConfig  set redis config
func WithRedisConfig(roleRuleHashKey string, redisClientAddress ...string) Option {
	return newFuncOption(func(o *Enforcer) {
		o.RoleRuleHashKey = roleRuleHashKey
		o.RedisClientAddress = redisClientAddress
	})
}

// set default params
func setDefaultOptions() *Enforcer {
	return &Enforcer{
		dbAdapter: db.NewDbAdapter(),
	}
}

type clientConn struct {
	timeout int
	dopts   *Enforcer
}

// NewEnforcer entrance
func NewEnforcer(opts ...Option) (*Enforcer, error) {
	cc := &clientConn{
		timeout: 30,
		dopts:   setDefaultOptions(),
	}

	// for opts
	for _, opt := range opts {
		opt.apply(cc.dopts)
	}

	// Open the DB, create it if not existed.
	err := cc.dopts.open()
	if err != nil {
		return nil, err
	}

	// call the destructor when the object is released.
	runtime.SetFinalizer(cc.dopts, finalizer)

	cc.dopts.TableNameGroupAccess = cc.dopts.tableNames[2]
	cc.dopts.TableNameGroup = cc.dopts.tableNames[1]
	cc.dopts.TableNameRule = cc.dopts.tableNames[0]

	return cc.dopts, nil
}

// finalizer is the destructor for Adapter.
func finalizer(a *Enforcer) {
	if a.engine == nil {
		return
	}

	err := a.engine.Close()
	if err != nil {
		logger.Infof("close xorm adapter engine failed, err: %v", err)
	}
}

type Enforcer struct {
	driverName           string
	dataSourceName       string
	dbSpecified          bool
	isFiltered           bool
	engine               *xorm.Engine
	tableNames           []string //tableNames注意顺序 第一个auth_rule 第二个auth_group 第三个auth_group_access
	TableNameRule        string
	TableNameGroup       string
	TableNameGroupAccess string
	RoleRuleHashKey      string //存放groupId_ruleUrlPath = ruleId    用于检查权限
	dbAdapter            *db.DbAdapter
	RedisClientAddress   []string
	MasterDB             *goqu.Database
	RedisClient          *redis.ClusterClient
}

func (a *Enforcer) open() error {
	var err error
	var engineNew *xorm.Engine
	engineNew, err = xorm.NewEngine(a.driverName, a.dataSourceName)
	if err != nil {
		return err
	}
	engineNew.ShowSQL()

	a.dbAdapter.Engine = engineNew
	a.engine = engineNew

	if err = a.createDatabase(); err != nil {
		return err
	}

	if err = a.createTable(); err != nil {
		return err
	}

	if err = a.newDataBase(); err != nil {
		return err
	}

	if err = a.newRedisClient(); err != nil {
		return err
	}

	return nil
}

func (a *Enforcer) newRedisClient() error {
	var err error
	if a.RedisClient, err = redislib.NewCluster(a.RedisClientAddress); err != nil {
		return err
	}
	return nil
}

func (a *Enforcer) newDataBase() error {
	var err error
	switch a.driverName {
	case constant.MysqlDriverName:
		if a.MasterDB, err = a.dbAdapter.Mysql().NewDriver(a.dataSourceName); err != nil {
			return err
		}
	case constant.PostgresDriverName:
		if a.MasterDB, err = a.dbAdapter.Pg().NewDriver(a.dataSourceName); err != nil {
			return err
		}
	default:
		return errors.New(constant.ErrDriverName)
	}
	return err
}

func (a *Enforcer) createTable() error {
	var err error
	switch a.driverName {
	case constant.MysqlDriverName:
		if err = a.dbAdapter.Mysql().CreateTable(a.tableNames...); err != nil {
			return err
		}
	case constant.PostgresDriverName:
		if err = a.dbAdapter.Pg().CreateTable(a.tableNames...); err != nil {
			return err
		}
	default:
		return errors.New(constant.ErrDriverName)
	}
	return err
}

func (a *Enforcer) createDatabase() error {
	var err error
	dataSourceNameArr := strings.Split(a.dataSourceName, "?")
	if len(dataSourceNameArr) != 2 {
		return errors.New("invalid parameter: dataSourceName")
	}
	dataSourceArr := strings.Split(dataSourceNameArr[0], "/")
	dataBaseName := dataSourceArr[len(dataSourceArr)-1]

	switch a.driverName {
	case constant.MysqlDriverName:
		err = a.dbAdapter.Mysql().CreateDatabase(dataBaseName)
	case constant.PostgresDriverName:
		if err = a.dbAdapter.Pg().CreateDatabase(dataBaseName); err != nil {
			// 42P04 is	duplicate_database
			if pqerr, ok := err.(*pq.Error); ok && pqerr.Code == "42P04" {
				return nil
			}
		}
	default:
		return errors.New(constant.ErrDriverName)
	}

	if err != nil {
		return err
	}

	return nil
}
