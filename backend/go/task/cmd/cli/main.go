package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gushikem01/cryptcurrency_autotrade/task/cmd/cli/di"
)

func main() {
	ctx := context.Background()
	cli, cleanup, err := di.InilializeTask(ctx)
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
		return
	}

	defer cleanup()

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		defer os.Exit(1)
		return
	}
}
