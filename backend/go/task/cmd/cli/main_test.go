package main_test

import (
	"context"
	"strings"
	"testing"

	"github.com/gushikem01/cryptcurrency_autotrade/task/cmd/cli/di"
)

func TestNewRootCmd(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{args: []string{"traders"}, want: ""},
	}
	ctx := context.Background()
	cli, cleanup, err := di.InilializeTask(ctx)
	if err != nil {
		t.Log(err)
	}
	defer cleanup()

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, ""), func(t *testing.T) {
			cli.RootCmd.SetArgs(tt.args)
			if err := cli.RootCmd.Execute(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
