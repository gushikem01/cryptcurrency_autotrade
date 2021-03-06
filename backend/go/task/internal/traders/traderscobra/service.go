package traderscobra

import (
	"context"
	"fmt"

	bitbankmodel "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank/public/model"
	bitflayermodel "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer/public/model"
	coincheckmodel "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public/model"
	gmomodel "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/public/model"
	liquidmodel "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public/model"
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
	// GMOコイン
	ch1 := make(chan *gmomodel.Ticker)
	go func() {
		m, err := cmd.gmoSrv.Ticker(ctx)
		if err != nil {
			cmd.logger.Error(err.Error())
		}
		ch1 <- m
	}()
	m := <-ch1
	fmt.Printf("GMOコイン:買%s 売%s", m.Data[0].Ask, m.Data[0].Bid)

	// CoinCheck
	ch2 := make(chan *coincheckmodel.Ticker)
	go func() {
		m2, err := cmd.coincheckSrv.TickerBtcJpy(ctx)
		if err != nil {
			cmd.logger.Error(err.Error())
		}
		ch2 <- m2
	}()
	m2 := <-ch2
	fmt.Printf("コインチェック:買%s 売%s", m2.Ask, m2.Bid)

	// TODO: Liquid
	ch3 := make(chan *liquidmodel.Products)
	go func() {
		m3, err := cmd.liquidSrv.ProductBtcJpy(ctx)
		if err != nil {
			cmd.logger.Error(err.Error())
		}
		ch3 <- m3
	}()
	m3 := <-ch3
	fmt.Printf("Liquid:買%g 売%g", m3.MarketAsk, m3.MarketBid)

	ch4 := make(chan []*bitbankmodel.Ticker)
	go func() {
		m4, err := cmd.bitbankSrv.Tickers(ctx, "btc_jpy")
		if err != nil {
			cmd.logger.Error(err.Error())
		}
		ch4 <- m4
	}()
	m4 := <-ch4
	fmt.Printf("Bitbank:買%s 売%s", m4[0].Buy, m4[0].Sell)

	// TODO: Bitflayer
	ch5 := make(chan *bitflayermodel.Ticker)
	go func() {
		m5, err := cmd.bitflayer.Ticker(ctx)
		if err != nil {
			cmd.logger.Error(err.Error())
		}
		ch5 <- m5
	}()
	m5 := <-ch5
	fmt.Printf("Bitflayer:買%s 売%s", m5.BestAsk, m5.BestBid)
	return nil
}
