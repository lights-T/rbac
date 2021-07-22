package authgroup

import (
	"context"

	. "github.com/lights-T/rbac/config"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

const (
	MaxLimit = 1000
)

var (
	ColumnFields = []interface{}{"id", "title", "rules", "create_time", "update_time"}
)

type HAuthGroup struct {
	Id         int64  `db:"id" json:"id,omitempty" goqu:"pk,skipinsert,skipupdate"` //
	Title      string `db:"title" json:"title,omitempty" goqu:"defaultifempty"`     // 身份名称
	Rules      string `db:"rules" json:"rules,omitempty" goqu:"defaultifempty"`     // 权限动作id
	CreateTime string `json:"createTime" db:"create_time" goqu:"skipinsert,skipupdate"`
	UpdateTime string `json:"updateTime" db:"update_time" goqu:"skipinsert,skipupdate"`
}

// GetHAuthGroup exps 支持 map[string]interface{} 或 goqu 表达式（eq: exp.NewExpressionList(exp.AndType).Append(goqu.C(k).Eq(v))）
func GetHAuthGroup(ctx context.Context, exps interface{}, excludeFields ...string) (*HAuthGroup, error) {
	self := &HAuthGroup{}
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
	if _, err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroup).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		ScanStructContext(ctx, self); err != nil {
		return nil, err
	}

	return self, nil
}

func GetHAuthGroupWithFields(ctx context.Context, exps interface{}, includeFields ...string) (*HAuthGroup, error) {
	self := &HAuthGroup{}
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
	if _, err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroup).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		ScanStructContext(ctx, self); err != nil {
		return nil, err
	}

	return self, nil
}

func SearchHAuthGroup(ctx context.Context, exps interface{}, excludeFields ...string) ([]*HAuthGroup, error) {
	var self []*HAuthGroup
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
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroup).
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

func SearchHAuthGroupWithFields(ctx context.Context, exps interface{}, includeFields ...string) ([]*HAuthGroup, error) {
	var self []*HAuthGroup
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
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroup).
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

func SearchHAuthGroupWithFieldsLimit(ctx context.Context, exps interface{}, offset, limit uint, includeFields ...string) ([]*HAuthGroup, error) {
	var self []*HAuthGroup
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
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroup).
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

func CreateHAuthGroup(ctx context.Context, h *HAuthGroup, excludeFields ...string) error {
	builder := AdapterContext.MasterDB.Insert(AdapterContext.TableNameGroup)
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

func UpdateHAuthGroup(ctx context.Context, data map[string]interface{}, exps interface{}) (int64, error) {
	builder := AdapterContext.MasterDB.Update(AdapterContext.TableNameGroup)
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
