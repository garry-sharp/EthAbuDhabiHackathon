package contract

import (
	"context"
	"log"
	"math/big"
	"sync"
	"time"
	iuniswapv2factory "xxx/pkg/abis/IUniswapV2Factory"
	iuniswapv2pair "xxx/pkg/abis/IUniswapV2Pair"
	"xxx/pkg/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var factory *iuniswapv2factory.Iuniswapv2factory

type Storage struct {
	Base            *config.Token
	Quote           *config.Token
	Inverse         bool
	Dex             *config.Dex
	Price           big.Float
	BaseLiquidity   big.Int
	QuoteLiquidity  big.Int
	ContractAddress common.Address
}

// TODO this would be redis or a DB or something in real life. But this is a hackathon
var storage map[*config.Dex]map[*config.Token]map[*config.Token]*Storage

func init() {
	storage = make(map[*config.Dex]map[*config.Token]map[*config.Token]*Storage)
}

func getStorageAndInit(dex *config.Dex, base *config.Token, quote *config.Token) *Storage {
	if storage[dex] == nil {
		storage[dex] = make(map[*config.Token]map[*config.Token]*Storage)
	}
	if storage[dex][base] == nil {
		storage[dex][base] = make(map[*config.Token]*Storage)
	}
	return storage[dex][base][quote]
}

func UniswapV2DexQuery(dex *config.Dex, tokens []*config.Token) error {
	log.Printf("Starting price Jobs for dex %s at address %s\n", dex.Name, dex.ContractAddress)

	if len(tokens) < 2 {
		log.Println("Less than 2 tokens, nothing to look up price for")
	}

	c, _ := GetClient()
	var err error
	factory, err = iuniswapv2factory.NewIuniswapv2factory(common.HexToAddress(dex.ContractAddress), c)

	for i := 0; i < len(tokens); i++ {
		for j := 0; j < len(tokens); j++ {
			if j <= i {
				continue
			}
			go UniswapV2PriceQuery(dex, tokens[i], tokens[j])
		}
	}

	time.Sleep(time.Second * 10)

	if err != nil {
		return err
	}
	return nil
}

func UniswapV2PriceQuery(dex *config.Dex, tokenA, tokenB *config.Token) error {
	log.Printf("Initializing price job for %s and %s on %s\n", tokenA.Name, tokenB.Name, dex.Name)

	contractAddress, err := factory.GetPair(&bind.CallOpts{
		Context: context.Background(),
	}, common.HexToAddress(tokenA.Address), common.HexToAddress(tokenB.Address))
	if err != nil {
		return err
	}
	c, err := GetClient()
	if err != nil {
		return err
	}
	contract, err := iuniswapv2pair.NewIuniswapv2pair(contractAddress, c)
	if err != nil {
		return err
	}

	for {
		reserves, err := contract.GetReserves(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			return err
		}

		token0address, _ := contract.Token0(nil)
		inverse := true
		if token0address.Hex() == common.HexToAddress(tokenA.Address).Hex() {
			inverse = false
		}

		var baseLiquidity big.Int = *reserves.Reserve0
		var quoteLiquidity big.Int = *reserves.Reserve1

		var price *big.Float
		if !inverse {
			price = new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve0), new(big.Float).SetInt(reserves.Reserve1))
		} else {
			_t := baseLiquidity
			baseLiquidity = quoteLiquidity
			quoteLiquidity = _t
			price = new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve1), new(big.Float).SetInt(reserves.Reserve0))
		}

		//log.Printf("Got new price of %.4f for %s/%s on %s\n", price, tokenA.Symbol, tokenB.Symbol, dex.Name)
		getStorageAndInit(dex, tokenA, tokenB)
		storage[dex][tokenA][tokenB] = &Storage{
			Base:            tokenA,
			Quote:           tokenB,
			Dex:             dex,
			Inverse:         inverse,
			Price:           *price,
			BaseLiquidity:   baseLiquidity,
			QuoteLiquidity:  quoteLiquidity,
			ContractAddress: contractAddress,
		}

		time.Sleep(config.GetConfig().PriceRefreshInterval)
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func UniswapV2PricePusher(tokens []*config.Token) {
	dexes := []*config.Dex{}
	for dex := range storage {
		dexes = append(dexes, dex)
	}
	go func() {
		for {
			var wg sync.WaitGroup
			wg.Add(factorial(len(tokens) - 1))
			for i := 0; i < len(tokens); i++ {
				for j := i + 1; j < len(tokens); j++ {
					base := tokens[i]
					quote := tokens[j]
					var bestBuy *Storage
					var bestSell *Storage
					for dexI := range dexes {
						//No race condition of values changing mid execution since the price fetcher doesn't update the value of the storage map by reference rathers adds a whole new pointer value
						s := getStorageAndInit(dexes[dexI], base, quote)
						if s == nil {
							continue
						} else {
							if bestBuy == nil || s.Price.Cmp(&bestBuy.Price) < 0 {
								bestBuy = s
							}
							if bestSell == nil || s.Price.Cmp(&bestSell.Price) > 0 {
								bestSell = s
							}
						}
					}
					WritePriceData(bestBuy, bestSell)
					wg.Done()
				}
			}
			wg.Wait()
			time.Sleep(config.GetConfig().PriceRefreshInterval)
		}
	}()
}
