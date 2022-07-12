package utils

import (
	"github.com/janotchain/janotchain/eth"
	"github.com/janotchain/janotchain/eth/downloader"
	"github.com/janotchain/janotchain/ethstats"
	"github.com/janotchain/janotchain/les"
	"github.com/janotchain/janotchain/node"
	"github.com/janotchain/janotchain/janot"
	"github.com/janotchain/janotchain/janotlending"
	whisper "github.com/janotchain/janotchain/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var tomoXServ *janot.TomoX
			ctx.Service(&tomoXServ)
			var lendingServ *janotlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, tomoXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterTomoXService(stack *node.Node, cfg *janot.Config) {
	tomoX := janot.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return tomoX, nil
	}); err != nil {
		Fatalf("Failed to register the TomoX service: %v", err)
	}

	// register janotlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return janotlending.New(tomoX), nil
	}); err != nil {
		Fatalf("Failed to register the TomoXLending service: %v", err)
	}
}
