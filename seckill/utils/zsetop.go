package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func AddToSortedSet(ctx context.Context, rdb *redis.Client, key string, members []*redis.Z) (bool, error) {
	// 添加元素到有序集合
	err := rdb.ZAdd(ctx, key, members...).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteToSortedSet(ctx context.Context, rdb *redis.Client, key string, members string) (bool, error) {
	// 添加元素到有序集合
	err := rdb.ZRem(ctx, key, members).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func FindToSortedSet(ctx context.Context, rdb *redis.Client, key string, opts *redis.ZRangeBy) ([]redis.Z, error) {
	result, err := rdb.ZRangeByScoreWithScores(ctx, key, opts).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
