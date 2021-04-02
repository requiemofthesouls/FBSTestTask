package service

import (
	r "FBSTestTask/internal/repository"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
)

type FibonacciService struct {
	repo r.FibonacciRepository
}

func NewFibonacciService(repo r.FibonacciRepository) *FibonacciService {
	return &FibonacciService{repo: repo}
}

func (f *FibonacciService) getFibSequence(n int) []int {
	fibSequence := make([]int, n+1, n+2)
	if n < 2 {
		fibSequence = fibSequence[0:2]
	}
	fibSequence[0] = 0
	fibSequence[1] = 1
	for i := 2; i <= n; i++ {
		fibSequence[i] = fibSequence[i-1] + fibSequence[i-2]
	}

	return fibSequence
}

func (f *FibonacciService) GetFibSlice(ctx context.Context, start, end int) (string, error) {
	if start < 0 || end < 0 {
		return "", errors.New("start & end should be positive")
	}

	if start > end {
		return "", errors.New("start can't be greater than end")
	}

	cachedSequence, err := f.repo.GetCachedFibSequence(ctx, end)
	if err != nil {
		log.Println(err)
	}

	if cachedSequence != nil {
		return f.filterFibSequence(cachedSequence, start, end), nil
	}

	sequence := f.getFibSequence(end)
	err = f.repo.SetFibSequence(ctx, sequence, end)
	if err != nil {
		fmt.Println("Can't save to redis: ", err)
	}
	return f.filterFibSequence(sequence, start, end), nil
}

func (f *FibonacciService) filterFibSequence(fibSequence []int, start int, end int) string {
	var b strings.Builder

	for i := start; i <= end; i++ {
		_, _ = fmt.Fprintf(&b, "%d:%d, ", i, fibSequence[i])
	}

	s := b.String()
	s = s[:b.Len()-2]
	return s
}
