package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("jane", "abc123")
}

func ProcessRequest(userId, authToken string) {
	ctx := context.WithValue(context.Background(), "userId", userId)
	ctx = context.WithValue(ctx, "authToken", authToken)
	HandleRequest(ctx)
}

func HandleRequest(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (%v)",
		ctx.Value("userId"),
		ctx.Value("authToken"),
	)
}
