package traderscobra

import (
	"context"
	"fmt"

	coincheckmodel "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public/model"
	gmomodel "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/public/model"
	"github.com/shopspring/decimal"
)

// Service ...
type Service struct {
	cmd *Cmd
}

// NewService ...
func NewService(cmd *Cmd) *Service {
	return &Service{cmd: cmd}
}

// GetPriceByTraders ...
func (cmd *Cmd) GetPriceByTraders(ctx context.Context) error {
	ch1 := make(chan *gmomodel.Ticker)
	go func() {
		m, err := cmd.gmoSrv.Ticker(ctx)
		if err != nil {
			cmd.logger.Error(err.Error())
		}
		ch1 <- m
	}()
	m := <-ch1

	fmt.Printf("GMOコイン:買%s 売%s", fmt.Sprintf("%d", m.Data[0].Ask.Mul(decimal.NewFromInt(3))), fmt.Sprintf("%d", m.Data[0].Bid.Mul(decimal.NewFromInt(3))))

	ch2 := make(chan *coincheckmodel.Ticker)
	go func() {
		m2, err := cmd.coincheckSrv.TickerBtcJpy(ctx)
		if err != nil {
			cmd.logger.Error(err.Error())
		}
		ch2 <- m2
	}()
	m2 := <-ch2
	fmt.Printf("コインチェック:買%s 売%s", fmt.Sprintf("%d", m2.Ask.Mul(decimal.NewFromInt(3))), fmt.Sprintf("%d", m2.Bid.Mul(decimal.NewFromInt(3))))

	return nil
}
