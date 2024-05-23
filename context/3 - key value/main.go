package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value666")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("key").(string)
	fmt.Println(token)
}
