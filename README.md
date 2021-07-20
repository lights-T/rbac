### 后台rbac权限

自构建rbac权限后台

### 特点

- 自动化创建相关数据库
- 支持多角色用户的CRUD管理，更方便灵活。
- 支持后台结合权限生成菜单，更加隐秘。
- 支持后台菜单排序，更加易于操作。
- 结合redis缓存，减少性能上的不必要浪费。

### 使用方法

```
ctx, cance := context.WithTimeout(context.TODO(), time.Second)
	defer cance()
	driverName := "mysql"
	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8"
	e, err := enforcer.NewEnforcer(
		enforcer.WithDbConfig(driverName, dataSourceName, []string{"auth_rule", "auth_group", "auth_group_access"}...),
		enforcer.WithRedisConfig("elysium:admin:roleRule", []string{"10.4.61.61:9001", "10.4.61.61:9002", "10.4.61.61:9003"}...),
	)
	if err != nil {
		logger.Fatal(err)
	}
	handler, err := NewHandler(e)
	if err != nil {
		logger.Fatal(err)
	}
	req := &pb.SelectRuleReq{}
	rsp := &pb.SelectRuleRsp{}
	_ = handler.SelectRule(ctx, req, rsp)
	t.Log(rsp)
```