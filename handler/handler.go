package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/lights-T/rbac/enforcer"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lights-T/rbac/cache"
	. "github.com/lights-T/rbac/config"
	"github.com/lights-T/rbac/constant"
	"github.com/lights-T/rbac/model/authgroup"
	"github.com/lights-T/rbac/model/authgroupaccess"
	"github.com/lights-T/rbac/model/authrule"
	pb "github.com/lights-T/rbac/proto"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type Handler struct{}

func NewHandler(enforcer *enforcer.Enforcer) (*Handler, error) {
	AdapterContext = &Adapter{
		MasterDB:             enforcer.MasterDB,
		RedisClient:          enforcer.RedisClient,
		TableNameRule:        enforcer.TableNameRule,
		TableNameGroup:       enforcer.TableNameGroup,
		TableNameGroupAccess: enforcer.TableNameGroupAccess,
		RoleRuleHashKey:      enforcer.RoleRuleHashKey,
	}
	if err := InitGroupRule(); err != nil {
		return nil, err
	}
	return new(Handler), nil
}

func (s *Handler) CreateRule(ctx context.Context, req *pb.CreateRuleReq, rsp *pb.EmptyRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}
	if err = AdapterContext.MasterDB.WithTx(func(tx *goqu.TxDatabase) error {
		for _, info := range req.Data {

			if err = checkTypeByPid(info.Pid, info.Type-1); err != nil {
				return err
			}
			if info.Pid > 0 && info.Type == 1 {
				err = fmt.Errorf("pid(%d) and type(%d) error", info.Pid, info.Type)
				return err
			}

			params := make(map[string]interface{})
			params["url_path"] = info.UrlPath
			needFields := []string{"id", "is_delete"}
			list, err := authrule.SearchHAuthRuleWithFields(ctx, params, needFields...)
			if err != nil {
				return err
			}

			if list != nil && len(list) > 0 {
				if list[0].IsDelete == constant.UnDelete {
					return fmt.Errorf("ruleID: %d is exist", list[0].Id)
				}
				upData := make(map[string]interface{})
				if len(info.UrlPath) > 0 {
					upData["url_path"] = info.UrlPath
				}
				if info.Pid > 0 {
					upData["pid"] = info.Pid
				}
				if len(info.Title) > 0 {
					upData["title"] = info.Title
				}
				if info.Type > 0 {
					upData["type"] = info.Type
				}
				if info.Sort > 0 {
					upData["sort"] = info.Sort
				}
				upData["is_delete"] = constant.UnDelete
				if _, err = authrule.UpdateHAuthRule(ctx, tx, upData, map[string]interface{}{"id": list[0].Id}); err != nil {
					return err
				}
				continue
			}
			record := &authrule.HAuthRule{
				UrlPath: info.UrlPath,
				Pid:     info.Pid,
				Title:   info.Title,
				Type:    info.Type,
			}
			if info.Sort > 0 {
				record.Sort = info.Sort
			}
			if err = authrule.CreateHAuthRule(ctx, tx, record); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// 检查pid类型
// _type为pid的类型
func checkTypeByPid(pid int64, _type int32) error {
	//-1为一级id
	if pid == 0 || pid == -1 {
		return nil
	}
	ctx := context.TODO()
	params := make(map[string]interface{})
	params["id"] = pid
	needFields := []string{"id", "type"}
	list, err := authrule.SearchHAuthRuleWithFields(ctx, params, needFields...)
	if err != nil {
		return fmt.Errorf("SearchHAuthRuleWithFields error:%s", err.Error())
	}
	if list == nil || len(list) == 0 {
		return fmt.Errorf("rule: %d is not exist", pid)
	}
	if list[0].Type != _type {
		return fmt.Errorf("ruleId: %d, pid(%d) type is not %d", pid, list[0].Type, _type)
	}
	return nil
}

func (s *Handler) UpdateRule(ctx context.Context, req *pb.UpdateRuleReq, rsp *pb.EmptyRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}

	if req.IsDelete == 1 {
		upData := make(map[string]interface{})
		upData["is_delete"] = constant.IsDelete
		if _, err = authrule.UpdateHAuthRule(ctx, nil, upData, map[string]interface{}{"url_path": req.UrlPath}); err != nil {
			return err
		}
		if err = InitGroupRule(); err != nil {
			return err
		}

		return nil
	}

	upData := make(map[string]interface{})
	if len(req.Title) > 0 {
		upData["title"] = req.Title
	}
	if req.Sort > 0 {
		upData["sort"] = req.Sort
	}
	if len(upData) == 0 {
		return nil
	}
	upData["is_delete"] = constant.UnDelete
	upData["update_time"] = time.Now().Format(constant.DatetimeLayout)
	if _, err = authrule.UpdateHAuthRule(ctx, nil, upData, map[string]interface{}{"url_path": req.UrlPath}); err != nil {
		return err
	}
	if err = InitGroupRule(); err != nil {
		return err
	}

	return nil
}

func (s *Handler) SelectRule(ctx context.Context, req *pb.SelectRuleReq, rsp *pb.SelectRuleRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}
	params := make(map[string]interface{})
	if req.Type > 0 {
		params["type"] = req.Type
	}
	params["is_delete"] = constant.UnDelete
	list, err := authrule.SearchHAuthRuleWithFields(ctx, params)
	if err != nil {
		return err
	}

	for _, v := range list {
		rsp.List = append(rsp.List, &pb.RuleInfo{
			UrlPath: v.UrlPath,
			Title:   v.Title,
			Type:    v.Type,
			Pid:     v.Pid,
			Sort:    v.Sort,
		})
	}

	return nil
}

func (s *Handler) CreateAuthGroup(ctx context.Context, req *pb.CreateAuthGroupReq, rsp *pb.EmptyRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}
	groups, err := CheckGroupExist(req.Title, []int64{})
	if err != nil {
		return err
	}
	if groups > 0 {
		return err
	}

	rules, err := CheckRule(req.RuleIds)
	if err != nil {
		return err
	}

	record := &authgroup.HAuthGroup{
		Title: req.Title,
		Rules: rules,
	}
	if err = authgroup.CreateHAuthGroup(ctx, record); err != nil {
		return err
	}
	if err = InitGroupRule(); err != nil {
		return err
	}

	return nil
}

func CheckRule(ruleIds []int64) (string, error) {
	var rules string
	params := make(map[string]interface{})
	params["is_delete"] = constant.UnDelete
	needFields := []string{"id"}
	ruleList, err := authrule.SearchHAuthRuleWithFields(context.TODO(), params, needFields...)
	if err != nil {
		return rules, err
	}
	if ruleList == nil {
		return rules, fmt.Errorf("rule不存在")
	}
	rulesHash := make(map[int64]int64, len(ruleList))
	for _, info := range ruleList {
		if _, ok := rulesHash[info.Id]; !ok {
			rulesHash[info.Id] = info.Id
		}
	}
	for _, id := range ruleIds {
		if _, ok := rulesHash[id]; !ok {
			return rules, fmt.Errorf("ruleId: %d, 不存在", id)
		}
		idStr := strconv.Itoa(int(id))
		if len(rules) == 0 {
			rules = idStr
		} else {
			rules = fmt.Sprintf("%s,%s", rules, idStr)
		}
	}
	return rules, nil
}

func (s *Handler) UpdateAuthGroup(ctx context.Context, req *pb.UpdateAuthGroupReq, rsp *pb.EmptyRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}
	group, err := CheckGroupExist(req.Title, []int64{req.Id})
	if err != nil {
		return err
	}
	if group == 0 {
		return errors.New("data not found")
	}
	rules, err := CheckRule(req.RuleIds)
	if err != nil {
		return err
	}

	upData := make(map[string]interface{})
	if len(req.Title) > 0 {
		upData["title"] = req.Title
	}
	if len(req.RuleIds) > 0 {
		upData["rules"] = rules
	}
	if len(upData) == 0 {
		return nil
	}
	upData["update_time"] = time.Now().Format(constant.DatetimeLayout)
	if _, err = authgroup.UpdateHAuthGroup(ctx, upData, map[string]interface{}{"id": req.Id}); err != nil {
		return err
	}
	if err = InitGroupRule(); err != nil {
		return err
	}

	return nil
}

func (s *Handler) SelectAuthGroup(ctx context.Context, req *pb.SelectAuthGroupReq, rsp *pb.SelectAuthGroupRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}

	params := make(map[string]interface{})
	if req.Id > 0 {
		params["id"] = req.Id
	}
	groupList, err := authgroup.SearchHAuthGroupWithFields(ctx, params, []string{"id", "title", "rules"}...)
	if err != nil {
		return err
	}
	if groupList == nil {
		return errors.New("group data not found")
	}
	exps := exp.NewExpressionList(exp.AndType)
	ruleIdsHash := make(map[string]string, 0)
	if req.Id > 0 {
		for _, v := range strings.Split(groupList[0].Rules, ",") {
			if _, ok := ruleIdsHash[v]; !ok {
				ruleIdsHash[v] = ""
			}
		}
	} else {
		for _, g := range groupList {
			rsp.List = append(rsp.List, &pb.AuthGroupRuleInfo{
				Id:         g.Id,
				GroupTitle: g.Title,
				Rules:      g.Rules,
			})
		}
		return nil
	}
	ruleList, err := authrule.SearchHAuthRuleWithFields(ctx, exps)
	if err != nil {
		return err
	}
	if ruleList == nil {
		return errors.New("rule data not found")
	}
	rules := make([]*pb.RuleInfo, 0, len(ruleList))
	for _, r := range ruleList {
		rule := &pb.RuleInfo{
			UrlPath:  r.UrlPath,
			Title:    r.Title,
			Type:     r.Type,
			Pid:      r.Pid,
			Sort:     r.Sort,
			Selected: false,
		}
		if _, ok := ruleIdsHash[strconv.Itoa(int(r.Id))]; ok {
			rule.Selected = true
		}
		rules = append(rules, rule)
	}
	rsp.List = append(rsp.List, &pb.AuthGroupRuleInfo{
		Id:         groupList[0].Id,
		GroupTitle: groupList[0].Title,
		Rules:      groupList[0].Rules,
		RuleList:   rules,
	})

	return nil
}

func CheckGroupExist(title string, ids []int64) (int, error) {
	var self int
	needFields := []string{"id"}
	exps := exp.NewExpressionList(exp.AndType)
	if len(title) > 0 && len(ids) > 0 {
		exps = exps.Append(goqu.L(" ( id in ? or title = ? ) ", ids, title))
	} else if len(ids) > 0 {
		exps = exps.Append(goqu.I("id").In(ids))
	} else if len(title) > 0 {
		exps = exps.Append(goqu.I("title").Eq(title))
	}
	roles, err := authgroup.SearchHAuthGroupWithFields(context.TODO(), exps, needFields...)
	if err != nil {
		return self, err
	}
	if roles == nil {
		return self, nil
	}
	return len(roles), nil
}

func (s *Handler) UpdateAuthGroupAccess(ctx context.Context, req *pb.UpdateAuthGroupAccessReq, rsp *pb.EmptyRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}
	group, err := CheckGroupExist("", req.GroupIds)
	if err != nil {
		return err
	}
	if group < len(req.GroupIds) {
		return errors.New("group data not found")
	}

	if err = AdapterContext.MasterDB.WithTx(func(tx *goqu.TxDatabase) error {
		upData := make(map[string]interface{})
		upData["is_delete"] = constant.IsDelete
		if _, err = authgroupaccess.UpdateHAuthGroupAccess(ctx, upData, map[string]interface{}{"uid": req.Uid}, tx); err != nil {
			return err
		}

		var dataUp []goqu.Record
		for _, id := range req.GroupIds {
			dataUp = append(dataUp, goqu.Record{
				"uid":         req.Uid,
				"group_id":    id,
				"update_time": time.Now().Format(constant.DatetimeLayout),
			})
		}
		if err = authgroupaccess.UpdateInsertHAuthGroupAccess(ctx, dataUp, nil, constant.UnDelete, tx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func InitGroupRule() error {
	ctx := context.TODO()
	//先清空
	if err := cache.Del(ctx, AdapterContext.RoleRuleHashKey); err != nil {
		return err
	}

	params := make(map[string]interface{})
	params["is_delete"] = constant.UnDelete
	ruleList, err := authrule.SearchHAuthRuleWithFields(ctx, params)
	if err != nil {
		return err
	}
	if ruleList == nil {
		fmt.Println("rule is not exist")
		return nil
	}
	ruleHash := make(map[int64]*authrule.HAuthRule, len(ruleList))
	for _, info := range ruleList {
		key := info.Id
		if _, ok := ruleHash[key]; !ok {
			ruleHash[key] = info
		}
	}
	groupParams := make(map[string]interface{})
	groupList, err := authgroup.SearchHAuthGroupWithFields(ctx, groupParams)
	if err != nil {
		return err
	}
	if groupList == nil {
		fmt.Println("group is not exist")
		return nil
	}
	for _, group := range groupList {
		if len(group.Rules) == 0 {
			continue
		}
		rules := strings.Split(group.Rules, ",")
		for _, gr := range rules {
			grInt, _ := strconv.Atoi(gr)
			if _, ok := ruleHash[int64(grInt)]; !ok {
				return errors.New(fmt.Sprintf("init failed, group: %d, ruleId: %d is not exist", group.Id, grInt))
			}
			key := fmt.Sprintf("%d_%s", group.Id, ruleHash[int64(grInt)].UrlPath)
			if err = cache.HSetRoleRule(ctx, key, int64(grInt)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Handler) VerifyUserAuth(ctx context.Context, req *pb.VerifyUserAuthReq, rsp *pb.VerifyUserAuthRsp) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}
	params := make(map[string]interface{})
	params["uid"] = req.Uid
	params["is_delete"] = constant.UnDelete
	groupList, err := authgroupaccess.SearchHAuthGroupAccessWithFields(ctx, params, []string{"uid", "group_id"}...)
	if err != nil {
		return err
	}
	if groupList == nil {
		return errors.New("group data not found")
	}
	var rowId int64
	for _, g := range groupList {
		key := fmt.Sprintf("%d_%s", g.GroupId, req.UrlPath)
		rowId, err = cache.HGetRoleRule(ctx, key)
		if err != nil {
			return err
		}
		if rowId > 0 {
			break
		}
	}

	rsp.Ok = true
	if rowId == 0 {
		rsp.Ok = false
	}
	rsp.Uid = req.Uid
	rsp.UrlPath = req.UrlPath
	return nil
}

func (s *Handler) GetUserMenu(ctx context.Context, req *pb.VerifyUserAuthReq) (map[int64]*authrule.HAuthRule, error) {
	var rsp map[int64]*authrule.HAuthRule
	var err error
	if err = req.Validate(); err != nil {
		return rsp, err
	}
	exps := exp.NewExpressionList(exp.AndType)
	exps = exps.Append(goqu.I("uid").Eq(req.Uid))
	exps = exps.Append(goqu.I("is_delete").Eq(constant.UnDelete))
	accessList, err := authgroupaccess.SearchHAuthGroupAccessWithFields(ctx, exps, "group_id")
	if err != nil {
		return rsp, err
	}
	if accessList == nil || len(accessList) == 0 {
		return rsp, errors.New("用户分组关系不存在")
	}
	groupIds := make([]int64, 0, len(accessList))
	for _, access := range accessList {
		groupIds = append(groupIds, access.GroupId)
	}
	exps = exp.NewExpressionList(exp.AndType)
	exps = exps.Append(goqu.I("id").In(groupIds))
	groupList, err := authgroup.SearchHAuthGroupWithFields(ctx, exps)
	if err != nil {
		return rsp, err
	}
	if groupList == nil {
		return rsp, errors.New("group is not exist")
	}

	rulesHash := make(map[int64]string, 500)
	rules := make([]int64, 0, 500)
	for _, group := range groupList {
		ruleArr := strings.Split(group.Rules, ",")
		for _, ruleIdStr := range ruleArr {
			ruleId, _ := strconv.Atoi(ruleIdStr)
			if _, ok := rulesHash[int64(ruleId)]; !ok {
				rulesHash[int64(ruleId)] = ""
				rules = append(rules, int64(ruleId))
			}
		}
	}
	exps = exp.NewExpressionList(exp.AndType)
	exps = exps.Append(goqu.I("id").In(rules))
	ruleList, err := authrule.SearchHAuthRuleWithFields(ctx, exps)
	if err != nil {
		return rsp, err
	}
	if ruleList == nil {
		return rsp, errors.New("rule is not exist")
	}
	firstMenu := make(map[int64]*authrule.HAuthRule, 500)    //最终组合
	secondMenu := make(map[int64][]*authrule.HAuthRule, 500) //[一级id]二级信息
	threeMenu := make(map[int64][]*authrule.HAuthRule, 500)  //[二级id]三级信息
	for _, rule := range ruleList {
		if rule.Type == constant.AuthRuleTypeFirst {
			firstMenu[rule.Id] = &authrule.HAuthRule{
				Id:        rule.Id,
				UrlPath:   rule.UrlPath,
				Title:     rule.Title,
				Type:      rule.Type,
				Pid:       rule.Pid,
				Sort:      rule.Sort,
				Icon:      rule.Icon,
				NextMenus: make([]*authrule.HAuthRule, 0, 20),
			}
		}
		if rule.Type == constant.AuthRuleTypeSecond {
			getGradeMenu(rule, secondMenu)
		}
		if rule.Type == constant.AuthRuleTypeThree {
			getGradeMenu(rule, threeMenu)
		}
	}

	//整理组合
	for _, f := range firstMenu {
		if _, ok := secondMenu[f.Id]; ok {
			for sk, sv := range secondMenu[f.Id] {
				if _, tok := threeMenu[sv.Id]; tok {
					secondMenu[f.Id][sk].NextMenus = threeMenu[sv.Id]
				}

				//大于一个，排个序
				if len(secondMenu[f.Id][sk].NextMenus) > 1 {
					sort.Slice(secondMenu[f.Id][sk].NextMenus, func(i, j int) bool {
						return secondMenu[f.Id][sk].NextMenus[i].Sort < secondMenu[f.Id][sk].NextMenus[j].Sort
					})
				}

			}
			firstMenu[f.Id].NextMenus = secondMenu[f.Id]
		}
		//大于一个，排个序
		if len(firstMenu[f.Id].NextMenus) > 1 {
			sort.Slice(firstMenu[f.Id].NextMenus, func(i, j int) bool {
				return firstMenu[f.Id].NextMenus[i].Sort < firstMenu[f.Id].NextMenus[j].Sort
			})
		}
	}

	return firstMenu, nil
}

// menu [父级id]二级信息
func getGradeMenu(rule *authrule.HAuthRule, menu map[int64][]*authrule.HAuthRule) {
	sec := &authrule.HAuthRule{
		Id:      rule.Id,
		UrlPath: rule.UrlPath,
		Title:   rule.Title,
		Type:    rule.Type,
		Pid:     rule.Pid,
		Sort:    rule.Sort,
		Icon:    rule.Icon,
	}
	if _, ok := menu[rule.Pid]; !ok {
		secondList := make([]*authrule.HAuthRule, 0, 20)
		secondList = append(secondList, sec)
		menu[rule.Pid] = secondList
	} else {
		menu[rule.Pid] = append(menu[rule.Pid], sec)
	}
}
