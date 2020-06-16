package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/vitelabs/go-vite/client"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/rpc"
	"github.com/vitelabs/go-vite/rpcapi/api"
)

var defaultHost = "http://127.0.0.1:48132"
var defaultAddress = "vite_0000000000000000000000000000000000000000a4f3a0cb58"
var address = flag.String("address", defaultAddress, "account address")
var startHeight = flag.Uint64("start", 1, "start height")
var endHeight = flag.Uint64("end", 1, "end height")
var host = flag.String("host", defaultHost, "full node http rpc host")

func main() {
	flag.Parse()

	rpc, err := client.NewRpcClient(*host)
	if err != nil {
		panic(err)
	}
	err = printAccBlocks(rpc, types.HexToAddressPanic(*address), *startHeight, *endHeight, func(block *api.AccountBlock) (b bool, b2 bool) {
		return true, true
	})

	if err != nil {
		panic(err)
	}
	fmt.Println()
}

type Matcher func(block *api.AccountBlock) (bool, bool)

func printAccBlocks(rpc client.RpcClient, address types.Address, start uint64, end uint64, matcherFunc Matcher) error {
	if start > end {
		end = start
	}
	index := end
	count := uint64(30)
	for {
		if index-start < count {
			count = index - start + 1
		}
		blocks, err := getBlocksByHeight(rpc.GetClient(), address, index, count)
		if err != nil {
			return err
		}
		for _, block := range blocks {
			if ledger.IsSendBlock(block.BlockType) {
				fmt.Printf("[S]%s,%s,%s,%d,%s,%d,%s\n",
					block.Height, block.Address,
					block.TokenInfo.TokenSymbol, block.TokenInfo.Decimals,
					*block.Amount, block.Timestamp,
					time.Unix(block.Timestamp, 0).String())
			} else if block.BlockType == ledger.BlockTypeReceive {
				fmt.Printf("[R]%s,%s,%s,%d,%s,%d,%s\n",
					block.Height, block.Address,
					block.TokenInfo.TokenSymbol, block.TokenInfo.Decimals,
					*block.Amount, block.Timestamp,
					time.Unix(block.Timestamp, 0).String())
			}
			heightUint64, err := strconv.ParseUint(block.Height, 10, 64)
			if err != nil {
				return err
			}
			index = heightUint64
		}

		if index <= start {
			break
		}
	}
	return nil
}

// tmp implements
func getBlocksByHeight(cc *rpc.Client, addr types.Address, height interface{}, count uint64) (blocks []*api.AccountBlock, err error) {
	err = cc.Call(&blocks, "ledger_getBlocksByHeight", addr, height, count)
	return
}
