package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Client with timeout
	// client := http.Client{Timeout: time.Second}
	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"Ivonei"}`))

	resp, err := client.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)

	// HTTP NewRequest and Context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	// context withou timeout
	// ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
