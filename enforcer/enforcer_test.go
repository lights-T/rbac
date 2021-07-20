package enforcer

import (
	"context"
	"testing"
	"time"
)

func TestRule_mysql(t *testing.T) {
	_, cance := context.WithTimeout(context.TODO(), time.Second)
	defer cance()
	_, err := NewEnforcer(
		WithDbConfig("mysql", "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8", []string{"auth_rule", "auth_group", "auth_group_access"}...),
		WithRedisConfig("elysium:admin:roleRule", []string{"10.4.61.61:9001", "10.4.61.61:9002", "10.4.61.61:9003"}...),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("mysql：", func(t *testing.T) {

	})
}

//func TestRule_pg(t *testing.T) {
//	driverName := "postgres"
//	dataSourceName := "postgres://postgres:pgsql123@10.4.61.84:5432/d_user?sslmode=disable"
//	adapterContext, err := NewEnforcer(driverName, dataSourceName)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Run("pg：", func(t *testing.T) {
//
//	})
//}
