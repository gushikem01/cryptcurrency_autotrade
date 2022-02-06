package traderscobra

import (
	"context"

	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Cmd コマンド
type Cmd struct {
	logger       *zap.Logger
	coincheckSrv *coincheck.Service
	gmoSrv       *gmo.Service
	liquidSrv    *liquid.Service
	bitbankSrv   *bitbank.Service
	btcboxSrv    *btcbox.Service
	bitflayer    *bitflayer.Service
}

// NewCmd コマンド作成
func NewCmd(
	logger *zap.Logger,
	coincheckSrv *coincheck.Service,
	gmoSrv *gmo.Service,
	liquidSrv *liquid.Service,
	bitbankSrv *bitbank.Service,
	btcboxSrv *btcbox.Service,
	bitflayerSrv *bitflayer.Service,
) *Cmd {
	return &Cmd{logger, coincheckSrv, gmoSrv, liquidSrv, bitbankSrv, btcboxSrv, bitflayerSrv}
}

// Run ...
func (h *Cmd) Run() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "traders",
		Short: "traders",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			srv := NewService(h)
			if err := srv.cmd.GetPriceByTraders(ctx); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
