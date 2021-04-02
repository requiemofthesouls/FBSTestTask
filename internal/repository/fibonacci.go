package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
)

type FibonacciRepository struct {
	key      string
	rdbCache *cache.Cache
}

func NewFibonacciRepository(rdbCache *cache.Cache) *FibonacciRepository {
	return &FibonacciRepository{rdbCache: rdbCache, key: "fibonacci"}
}

func (r *FibonacciRepository) GetCachedFibSequence(ctx context.Context, end int) ([]int64, error) {
	var result []int64

	if err := r.rdbCache.Get(ctx, fmt.Sprintf("%s:%d", r.key, end), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *FibonacciRepository) SetFibSequence(ctx context.Context, fibSequence []int64, end int) error {
	toSet := cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("%s:%d", r.key, end),
		Value: fibSequence,
		TTL:   0,
	}

	if err := r.rdbCache.Set(&toSet); err != nil {
		return err
	}

	return nil
}
