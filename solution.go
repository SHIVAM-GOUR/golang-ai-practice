package main

import (
	"context"
	"fmt"
)

type ctxKey string

const (
	userKey  ctxKey = "user"
	otherKey ctxKey = "other"
)

func main() {
	// var wg sync.WaitGroup

	ctx := context.WithValue(context.Background(), userKey, "martha-wayne")

	fn1(ctx)

}

func fn1(ctx context.Context) {
	ctx2 := context.WithValue(ctx, otherKey, "other")
	fn2(ctx2)
}

func fn2(ctx context.Context) {
	val := ctx.Value(userKey)

	if val == nil {
		println("key value missing from the context: ", userKey)
	} else {
		user, ok := val.(string)
		if !ok {
			println("unexpected type for value: ", userKey)
		} else {
			fmt.Println("context value: ", user)
		}

	}

}
