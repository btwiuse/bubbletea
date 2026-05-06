package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"examples/hackertea/internal/cache"
	"examples/hackertea/internal/client"
	"examples/hackertea/internal/constants"
	"examples/hackertea/internal/hn"
	"examples/hackertea/internal/tui/model"

	"github.com/btwiuse/boba"
)

func main() {
	ctx := context.Background()

	c := client.New(constants.BaseURL, &http.Client{Timeout: 10 * time.Second})
	memCache := cache.New()
	hnClient := hn.New(c, memCache)

	m, err := model.New(ctx, hnClient)
	if err != nil {
		fmt.Println("Error creating model: ", err)
		os.Exit(1)
	}

	p := boba.NewProgram(m)

	if _, err = p.Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
