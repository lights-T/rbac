package authgroupaccess

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
	ColumnFields = []interface{}{"uid", "group_id", "is_delete", "create_time", "update_time"}
)

type HAuthGroupAccess struct {
	Uid        int64  `db:"uid" json:"uid,omitempty" goqu:"defaultifempty"`            // 用户id
	GroupId    int64  `db:"group_id" json:"groupId,omitempty" goqu:"defaultifempty"`   // 身份id
	IsDelete   int32  `db:"is_delete" json:"isDelete,omitempty" goqu:"defaultifempty"` // 是否删除 1删除
	CreateTime string `json:"createTime" db:"create_time" goqu:"skipinsert,skipupdate"`
	UpdateTime string `json:"updateTime" db:"update_time" goqu:"skipinsert,skipupdate"`
}

// GetHAuthGroupAccess exps 支持 map[string]interface{} 或 goqu 表达式（eq: exp.NewExpressionList(exp.AndType).Append(goqu.C(k).Eq(v))）
func GetHAuthGroupAccess(ctx context.Context, exps interface{}, excludeFields ...string) (*HAuthGroupAccess, error) {
	self := &HAuthGroupAccess{}
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
	if _, err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroupAccess).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		ScanStructContext(ctx, self); err != nil {
		return nil, err
	}

	return self, nil
}

func GetHAuthGroupAccessWithFields(ctx context.Context, exps interface{}, includeFields ...string) (*HAuthGroupAccess, error) {
	self := &HAuthGroupAccess{}
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
	if _, err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroupAccess).
		Prepared(true).
		Select(cols...).
		Where(conditions).
		ScanStructContext(ctx, self); err != nil {
		return nil, err
	}

	return self, nil
}

func SearchHAuthGroupAccess(ctx context.Context, exps interface{}, excludeFields ...string) ([]*HAuthGroupAccess, error) {
	var self []*HAuthGroupAccess
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
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroupAccess).
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

func SearchHAuthGroupAccessWithFields(ctx context.Context, exps interface{}, includeFields ...string) ([]*HAuthGroupAccess, error) {
	var self []*HAuthGroupAccess
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
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroupAccess).
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

func SearchHAuthGroupAccessWithFieldsLimit(ctx context.Context, exps interface{}, offset, limit uint, includeFields ...string) ([]*HAuthGroupAccess, error) {
	var self []*HAuthGroupAccess
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
	if err := AdapterContext.MasterDB.From(AdapterContext.TableNameGroupAccess).
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

func CreateHAuthGroupAccess(ctx context.Context, h *HAuthGroupAccess, excludeFields ...string) (int64, error) {
	builder := AdapterContext.MasterDB.Insert(AdapterContext.TableNameGroupAccess)
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
	b, err := builder.Prepared(true).
		Cols(cols...).
		Rows(h).
		Executor().ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	return b.LastInsertId()
}

func UpdateHAuthGroupAccess(ctx context.Context, data map[string]interface{}, exps interface{}, tx *goqu.TxDatabase) (int64, error) {
	var builder *goqu.UpdateDataset
	if tx != nil {
		builder = tx.Update(AdapterContext.TableNameGroupAccess)
	} else {
		builder = AdapterContext.MasterDB.Update(AdapterContext.TableNameGroupAccess)
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

func UpdateInsertHAuthGroupAccess(ctx context.Context, data []goqu.Record, exps interface{}, isDelete int, tx *goqu.TxDatabase) error {
	var builder *goqu.InsertDataset
	if tx != nil {
		builder = tx.Insert(AdapterContext.TableNameGroupAccess)
	} else {
		builder = AdapterContext.MasterDB.Insert(AdapterContext.TableNameGroupAccess)
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
	_, err := builder.Prepared(false).Rows(data).OnConflict(
		goqu.DoUpdate("uid,group_id", goqu.Record{"is_delete": isDelete})).Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
