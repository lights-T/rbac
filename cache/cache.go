package cache

import (
	"context"
	"strconv"

	. "github.com/lights-T/rbac/config"

	"github.com/go-redis/redis/v8"
	"github.com/lights-T/lib-go/logger"
)

// HSetRoleRule 设置角色权限map
// roleRuleKey: groupId_ruleUrlPath
func HSetRoleRule(ctx context.Context, roleRuleKey string, value int64) error {
	if err := AdapterContext.RedisClient.HSetNX(ctx, AdapterContext.RoleRuleHashKey, roleRuleKey, value).Err(); err != nil {
		logger.Errorf("HSetRoleRule error:%s", err.Error())
		return err
	}
	return nil
}

// Del 删除
func Del(ctx context.Context, key string) error {
	if err := AdapterContext.RedisClient.Del(ctx, key).Err(); err != nil {
		logger.Errorf("Del error:%s", err.Error())
		return err
	}
	return nil
}

// HDel 哈希删除
func HDel(ctx context.Context, key string, fields ...string) error {
	if err := AdapterContext.RedisClient.HDel(ctx, key, fields...).Err(); err != nil {
		logger.Errorf("HDel error:%s", err.Error())
		return err
	}
	return nil
}

// HGetRoleRule 获取角色权限map
func HGetRoleRule(ctx context.Context, roleRuleKey string) (int64, error) {
	str, err := AdapterContext.RedisClient.HGet(ctx, AdapterContext.RoleRuleHashKey, roleRuleKey).Result()
	if err != nil && err != redis.Nil {
		logger.Errorf("HGetRoleRule error:%s", err.Error())
		return 0, err
	}
	if len(str) == 0 {
		return 0, nil
	}
	ruleId, _ := strconv.ParseInt(str, 10, 64)
	return ruleId, nil
}
