package service

import (
	r "FBSTestTask/internal/repository"
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestFibonacciService_GetFibSlice_ExpectError(t *testing.T) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"redis": ":6379",
		},
	})

	rc := cache.New(&cache.Options{
		Redis: ring,
	})

	repo := r.NewFibonacciRepository(rc)
	s := NewFibonacciService(*repo)

	expectedErr := "start & end should be positive"

	ctx := context.TODO()

	_, err := s.GetFibSlice(ctx, -10, 20)
	if err.Error() != expectedErr {
		t.Error("Fail")
	}

	_, err = s.GetFibSlice(ctx, 0, -20)
	if err.Error() != expectedErr {
		t.Error("Fail")
	}

	_, err = s.GetFibSlice(ctx, -10, -20)
	if err.Error() != expectedErr {
		t.Error("Fail")
	}

	expectedErr = "start can't be greater than end"
	_, err = s.GetFibSlice(ctx, 40, 20)
	if err.Error() != expectedErr {
		t.Error("Fail")
	}
}

func TestFibonacciService_GetFibSlicedMap(t *testing.T) {
	ctx := context.TODO()

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"redis": ":6379",
		},
	})

	rc := cache.New(&cache.Options{
		Redis: ring,
	})

	repo := r.NewFibonacciRepository(rc)
	s := NewFibonacciService(*repo)

	res, err := s.GetFibSlice(ctx, 240, 24400)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(res))
}
