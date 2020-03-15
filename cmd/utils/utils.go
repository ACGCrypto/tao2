package utils

import (
	"github.com/taoblockchain/tao2/eth"
	"github.com/taoblockchain/tao2/eth/downloader"
	"github.com/taoblockchain/tao2/ethstats"
	"github.com/taoblockchain/tao2/les"
	"github.com/taoblockchain/tao2/node"
	"github.com/taoblockchain/tao2/waihui"
	whisper "github.com/taoblockchain/tao2/whisper/whisperv6"
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
			var waihuiServ *waihui.Waihui
			ctx.Service(&waihuiServ)

			fullNode, err := eth.New(ctx, cfg, waihuiServ)
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

func RegisterWaihuiService(stack *node.Node, cfg *waihui.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return waihui.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Waihui service: %v", err)
	}
}
