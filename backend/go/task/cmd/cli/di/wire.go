//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/google/wire"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/log"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank"
	bitbankpublic "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitbank/public"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer"
	bitflayerpublic "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/bitflayer/public"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox"
	btcboxpublic "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/btcbox/public"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck"
	coincheckpublic "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/coincheck/public"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo"
	gmopublic "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/gmo/public"
	"github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid"
	liquidpublic "github.com/gushikem01/cryptcurrency_autotrade/pkg/traders/liquid/public"
	"github.com/gushikem01/cryptcurrency_autotrade/task/internal/app/appcobra"
	"github.com/gushikem01/cryptcurrency_autotrade/task/internal/traders/traderscobra"
)

// InilializeTask ...
func InilializeTask(ctx context.Context) (*appcobra.CLI, func(), error) {
	wire.Build(
		log.NewZap,
		liquid.NewService,
		liquidpublic.NewRepository,
		coincheck.NewService,
		coincheckpublic.NewRepository,
		gmo.NewService,
		gmopublic.NewRepository,
		bitbank.NewService,
		bitbankpublic.NewRepository,
		btcbox.NewService,
		btcboxpublic.NewRepository,
		bitflayer.NewService,
		bitflayerpublic.NewRepository,
		traderscobra.NewCmd,
		appcobra.NewCLI,
	)
	return nil, nil, nil
}
