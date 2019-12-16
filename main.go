package main

import (
	"fmt"
	"context"
	"cloud.google.com/go/storage"
)

func main() {
	// setup GCS client
	ctx := context.Background()
	_, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("Hello World")
}
