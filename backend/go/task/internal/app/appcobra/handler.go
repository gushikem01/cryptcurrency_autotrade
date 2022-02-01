package appcobra

import (
	"github.com/gushikem01/cryptcurrency_autotrade/task/internal/traders/traderscobra"
	"github.com/spf13/cobra"
)

// CLI DI
type CLI struct {
	RootCmd *cobra.Command
}

// NewCLI DI作成
func NewCLI(tc *traderscobra.Cmd) *CLI {
	cmd := &cobra.Command{
		Use:     "traders",
		Short:   "traders cli",
		Long:    "traders cli",
		Version: "1.0",
	}
	cmd.AddCommand(tc.Run())
	return &CLI{RootCmd: cmd}
}
