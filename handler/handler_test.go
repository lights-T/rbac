package handler

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/lights-T/lib-go/logger"
	"github.com/lights-T/rbac/enforcer"
	pb "github.com/lights-T/rbac/proto"
	"testing"
	"time"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		logger.Fatal("Error loading .env file")
	}
}

func Test(t *testing.T) {
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

}

//func TestRule_Select(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.SelectRuleReq{}
//	var RequestData = func(exps domain.SelectRuleReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		SelectRule(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("搜索：", func(t *testing.T) {
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestRule_Create(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.CreateRuleReq{}
//	var RequestData = func(exps domain.CreateRuleReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		CreateRule(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("创建：", func(t *testing.T) {
//		data := make([]*domain.RuleInfo, 0, 5)
//		//data = append(data, &domain.RuleInfo{
//		//	UrlPath: "model/action1",
//		//	Title:   "一级菜单1",
//		//	Type:    1,
//		//	Pid:     0,
//		//})
//		//data = append(data, &domain.RuleInfo{
//		//	UrlPath: "model/action2",
//		//	Title:   "一级菜单2",
//		//	Type:    1,
//		//	Pid:     0,
//		//})
//		//data = append(data, &domain.RuleInfo{
//		//	UrlPath: "model/action3",
//		//	Title:   "一级菜单3",
//		//	Type:    1,
//		//	Pid:     0,
//		//})
//		//data = append(data, &domain.RuleInfo{
//		//	UrlPath: "model/action4",
//		//	Title:   "二级菜单1",
//		//	Type:    2,
//		//	Pid:     20,
//		//})
//		//data = append(data, &domain.RuleInfo{
//		//	UrlPath: "model/action5",
//		//	Title:   "三级菜单1",
//		//	Type:    3,
//		//	Pid:     22,
//		//})
//		data = append(data, &domain.RuleInfo{
//			UrlPath: "model/action6",
//			Title:   "三级菜单2",
//			Type:    3,
//			Pid:     24,
//			Sort:    51,
//		})
//		exps.Data = data
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestRule_Update(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.UpdateRuleReq{}
//	var RequestData = func(exps domain.UpdateRuleReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		UpdateRule(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("修改：", func(t *testing.T) {
//		exps = domain.UpdateRuleReq{
//			UrlPath:  "model/action1",
//			Title:    "up - 一级菜单1",
//			IsDelete: 0,
//		}
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestAuthGroup_Create(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.CreateAuthGroupReq{}
//	var RequestData = func(exps domain.CreateAuthGroupReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		CreateAuthGroup(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("创建：", func(t *testing.T) {
//		exps = domain.CreateAuthGroupReq{
//			Title:   "超级管理员",
//			RuleIds: []int64{20, 21, 22},
//		}
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestAuthGroup_Update(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.UpdateAuthGroupReq{}
//	var RequestData = func(exps domain.UpdateAuthGroupReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		UpdateAuthGroup(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("修改角色：", func(t *testing.T) {
//		exps = domain.UpdateAuthGroupReq{
//			Id:      34,
//			Title:   "超级管理员",
//			RuleIds: []int64{20, 21, 22, 24, 25},
//		}
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestGroup_Select(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.SelectAuthGroupReq{}
//	var RequestData = func(exps domain.SelectAuthGroupReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		SelectAuthGroup(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("搜索全部：", func(t *testing.T) {
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//	t.Run("id搜索：", func(t *testing.T) {
//		exps.Id = 34
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestAuthGroupAccess_Create(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.UpdateAuthGroupAccessReq{}
//	var RequestData = func(exps domain.UpdateAuthGroupAccessReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		UpdateAuthGroupAccess(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("创建：", func(t *testing.T) {
//		exps = domain.UpdateAuthGroupAccessReq{
//			Uid:      1,
//			GroupIds: []int64{1, 2},
//		}
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestVerifyUserAuth(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.VerifyUserAuthReq{}
//	var RequestData = func(exps domain.VerifyUserAuthReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		VerifyUserAuth(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("验证：", func(t *testing.T) {
//		exps = domain.VerifyUserAuthReq{
//			GroupId: 34,
//			UrlPath: "model/action1",
//		}
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
//
//func TestGetUserMenu(t *testing.T) {
//	buf := bytes.NewBuffer(nil)
//	httpTest := &httptest.ResponseRecorder{Body: buf}
//	c, _ := gin.CreateTestContext(httpTest)
//	exps := domain.GetUserMenuReq{}
//	var RequestData = func(exps domain.GetUserMenuReq, t *testing.T, expected int) (bool, json.RawMessage) {
//		b, _ := json.Marshal(exps)
//		c.Request = &http.Request{Body: ioutil.NopCloser(bytes.NewReader(b))}
//		GetUserMenu(c)
//		var res *domain.Rsp
//		if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
//			assert.NoError(t, err, "json unmarshal")
//		}
//		logger.Infof("%v", string(res.Data))
//		if res.Code != 200 {
//			t.Fatal(res.Msg)
//		}
//
//		return assert.Equal(t, expected, res.Code), res.Data
//	}
//	t.Run("获取用户菜单：", func(t *testing.T) {
//		exps = domain.GetUserMenuReq{
//			Uid: 1,
//		}
//		RequestData(exps, t, 200)
//		buf.Reset()
//	})
//}
