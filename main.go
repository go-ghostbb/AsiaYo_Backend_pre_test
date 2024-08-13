package main

import (
	"AsiaYo-Backend-pre-test/internal/cmd"
	"context"
)

func main() {
	err := cmd.MainFn(context.Background())
	if err != nil {
		panic(err)
	}
}
