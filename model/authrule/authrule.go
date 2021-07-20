package authrule

import (
	"context"
	"time"

	. "github.com/lights-T/rbac/config"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

const (
	MaxLimit = 1000
)

var (
	ColumnFields = []interface{}{"id", "url_path", "title", "type", "is_delete", "pid", "sort", "icon", "create_time", "update_time"}
)

type HAuthRule struct {
	Id         int64     `db:"id" json:"id,omitempty" goqu:"pk,skipinsert,skipupdate"`    //
	UrlPath    string    `db:"url_path" json:"urlPath,omitempty" goqu:"defaultifempty"`   // 规则唯一标识路径 模块/方法
	Title      string    `db:"title" json:"title,omitempty" goqu:"defaultifempty"`        // 规则中文名称 要逻辑化定义，给运营人员使用
	Type       int32     `db:"type" json:"type,omitempty" goqu:"defaultifempty"`          // 类型 1一级菜单 2二级菜单 3三级菜单 -1详情/按钮
	IsDelete   int32     `db:"is_delete" json:"isDelete,omitempty" goqu:"defaultifempty"` // 是否删除 1删除
	Pid        int64     `db:"pid" json:"pid,omitempty" goqu:"defaultifempty"`            // 父级ID
	Sort       int32     `db:"sort" json:"sort,omitempty" goqu:"defaultifempty"`          // 排序
	Icon       string    `db:"icon" json:"icon,omitempty" goqu:"defaultifempty"`          //
	CreateTime time.Time `json:"createTime" db:"create_time" goqu:"skipinsert,skipupdate"`
	UpdateTime time.Time `json:"updateTime" db:"update_time" goqu:"skipinsert,skipupdate"`

	NextMenus []*HAuthRule `json:"nextMenus,omitempty" goqu:"skipinsert,skipupdate"` //下一等级菜单
}

// GetHAuthRule exps 支持 map[string]interface{} 或 goqu 表达式（eq: exp.NewExpressionList(exp.AndType).Append(goqu.C(k).Eq(v))）
func GetHAuthRule(ctx context.Context, exps interface{}, excludeFields ...string) (*HAuthRule, error) {
	self := &HAuthRule{}
	cols := make([]interface{}, 0, len(ColumnFields))
	if len(excludeFields) > 0 {
		_tempMap := make(map[string]struct{})
		for _, e := range excludeFields {
			_tempMap[e] = struct{}{}
		}
		for _, s := range ColumnFields {
			_s := s.(string)
			if _, ok := _tempMap[_s]; !ok {
				cols = append(cols, s)
			}
		}
	} else {
		cols = ColumnFields
	}
	conditions := exp.NewExpressionList(exp.AndType)
	switch exps.(type) {
	case map[string]interface{}:
		for k, v := range exps.(map[string]interface{}) {
			conditions = conditions.Append(goqu.I(k).Eq(v))
		}
	case exp.ExpressionList:
		conditions = exps.(exp.ExpressionList)
	}
	if _, err := AdapterContext.MasterDB.From(AdapterContext.TableNameRule).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		ScanStructContext(ctx, self); err != nil {
		return nil, err
	}

	return self, nil
}

func GetHAuthRuleWithFields(ctx context.Context, exps interface{}, includeFields ...string) (*HAuthRule, error) {
	self := &HAuthRule{}
	cols := make([]interface{}, 0, len(ColumnFields))
	if len(includeFields) > 0 {
		_tempMap := make(map[string]struct{})
		for _, e := range includeFields {
			_tempMap[e] = struct{}{}
		}
		for _, s := range ColumnFields {
			_s := s.(string)
			if _, ok := _tempMap[_s]; ok {
				cols = append(cols, s)
			}
		}

	} else {
		cols = ColumnFields
	}
	conditions := exp.NewExpressionList(exp.AndType)
	switch exps.(type) {
	case map[string]interface{}:
		for k, v := range exps.(map[string]interface{}) {
			conditions = conditions.Append(goqu.I(k).Eq(v))
		}
	case exp.ExpressionList:
		conditions = exps.(exp.ExpressionList)
	}
	if _, err := AdapterContext.MasterDB.From(AdapterContext.TableNameRule).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		ScanStructContext(ctx, self); err != nil {
		return nil, err
	}

	return self, nil
}

func SearchHAuthRule(ctx context.Context, exps interface{}, excludeFields ...string) ([]*HAuthRule, error) {
	var self []*HAuthRule
	cols := make([]interface{}, 0, len(ColumnFields))
	if len(excludeFields) > 0 {
		_tempMap := make(map[string]struct{})
		for _, e := range excludeFields {
			_tempMap[e] = struct{}{}
		}
		for _, s := range ColumnFields {
			_s := s.(string)
			if _, ok := _tempMap[_s]; !ok {
				cols = append(cols, s)
			}
		}
	} else {
		cols = ColumnFields
	}
	conditions := exp.NewExpressionList(exp.AndType)
	switch exps.(type) {
	case map[string]interface{}:
		for k, v := range exps.(map[string]interface{}) {
			conditions = conditions.Append(goqu.I(k).Eq(v))
		}
	case exp.ExpressionList:
		conditions = exps.(exp.ExpressionList)
	}
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameRule).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		Limit(MaxLimit).
		ScanStructsContext(ctx, &self); err != nil {
		return nil, err
	}
	if len(self) == 0 {
		return nil, nil
	}

	return self, nil
}

func SearchHAuthRuleWithFields(ctx context.Context, exps interface{}, includeFields ...string) ([]*HAuthRule, error) {
	var self []*HAuthRule
	cols := make([]interface{}, 0, len(ColumnFields))
	if len(includeFields) > 0 {
		_tempMap := make(map[string]struct{})
		for _, e := range includeFields {
			_tempMap[e] = struct{}{}
		}
		for _, s := range ColumnFields {
			_s := s.(string)
			if _, ok := _tempMap[_s]; ok {
				cols = append(cols, s)
			}
		}
	} else {
		cols = ColumnFields
	}
	conditions := exp.NewExpressionList(exp.AndType)
	switch exps.(type) {
	case map[string]interface{}:
		for k, v := range exps.(map[string]interface{}) {
			conditions = conditions.Append(goqu.I(k).Eq(v))
		}
	case exp.ExpressionList:
		conditions = exps.(exp.ExpressionList)
	}
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameRule).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		Limit(MaxLimit).
		ScanStructsContext(ctx, &self); err != nil {
		return nil, err
	}
	if len(self) == 0 {
		return nil, nil
	}

	return self, nil
}

func SearchHAuthRuleWithFieldsLimit(ctx context.Context, exps interface{}, offset, limit uint, includeFields ...string) ([]*HAuthRule, error) {
	var self []*HAuthRule
	cols := make([]interface{}, 0, len(ColumnFields))
	if limit > MaxLimit {
		limit = MaxLimit
	}
	if len(includeFields) > 0 {
		_tempMap := make(map[string]struct{})
		for _, e := range includeFields {
			_tempMap[e] = struct{}{}
		}
		for _, s := range ColumnFields {
			_s := s.(string)
			if _, ok := _tempMap[_s]; ok {
				cols = append(cols, s)
			}
		}
	} else {
		cols = ColumnFields
	}
	conditions := exp.NewExpressionList(exp.AndType)
	switch exps.(type) {
	case map[string]interface{}:
		for k, v := range exps.(map[string]interface{}) {
			conditions = conditions.Append(goqu.I(k).Eq(v))
		}
	case exp.ExpressionList:
		conditions = exps.(exp.ExpressionList)
	}
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameRule).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		Offset(offset).
		Limit(limit).
		ScanStructsContext(ctx, &self); err != nil {
		return nil, err
	}
	if len(self) == 0 {
		return nil, nil
	}

	return self, nil
}

func CreateHAuthRule(ctx context.Context, tx *goqu.TxDatabase, h *HAuthRule, excludeFields ...string) error {
	var builder *goqu.InsertDataset
	if tx != nil {
		builder = tx.Insert(AdapterContext.TableNameRule)
	} else {
		builder = AdapterContext.MasterDB.Insert(AdapterContext.TableNameRule)
	}
	cols := make([]interface{}, 0, 10)
	if len(excludeFields) > 0 {
		_tempMap := make(map[string]struct{})
		for _, e := range excludeFields {
			_tempMap[e] = struct{}{}
		}
		for _, s := range ColumnFields {
			_s := s.(string)
			if _, ok := _tempMap[_s]; !ok {
				cols = append(cols, s)
			}
		}
	} else {
		cols = ColumnFields
	}
	_, err := builder.Prepared(true).
		Cols(cols...).
		Rows(h).
		Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func UpdateHAuthRule(ctx context.Context, tx *goqu.TxDatabase, data map[string]interface{}, exps interface{}) (int64, error) {
	var builder *goqu.UpdateDataset
	if tx != nil {
		builder = tx.Update(AdapterContext.TableNameRule)
	} else {
		builder = AdapterContext.MasterDB.Update(AdapterContext.TableNameRule)
	}
	rc := make(goqu.Record)
	for k, v := range data {
		rc[k] = v
	}
	conditions := exp.NewExpressionList(exp.AndType)
	switch exps.(type) {
	case map[string]interface{}:
		for k, v := range exps.(map[string]interface{}) {
			conditions = conditions.Append(goqu.I(k).Eq(v))
		}
	case exp.ExpressionList:
		conditions = exps.(exp.ExpressionList)
	}
	u, err := builder.Set(rc).Where(conditions).Executor().ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	return u.RowsAffected()
}
