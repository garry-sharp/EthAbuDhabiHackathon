package contract

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"math/big"
	aggregator "xxx/pkg/abis/Aggregator"
	"xxx/pkg/config"

	"github.com/ethereum/go-ethereum/common"
)

var _aggregator *aggregator.Aggregator

type WriteArgs struct {
	token0          common.Address
	token1          common.Address
	token0Liquidity big.Int
	token1Liquidity big.Int
	contract        common.Address
}

func toECDSAFromHex(privateKey string) (*ecdsa.PrivateKey, error) {
	pk := new(ecdsa.PrivateKey)
	pk.D, _ = new(big.Int).SetString(privateKey, 16)
	pk.PublicKey.Curve = elliptic.P256()
	pk.PublicKey.X, pk.PublicKey.Y = pk.PublicKey.Curve.ScalarBaseMult(pk.D.Bytes())
	return pk, nil
}

// TODO error handling
func GetAggregator() *aggregator.Aggregator {
	c, _ := GetClient()
	if _aggregator == nil {
		addr := common.HexToAddress(config.GetConfig().AggregatorAddress)
		_aggregator, _ = aggregator.NewAggregator(addr, c)
	}
	return _aggregator
}

// TODO proper error handing
func Deploy() error {
	c, _ := GetClient()
	auth, _, _, address := getAuth()
	nonce, _ := c.PendingNonceAt(context.Background(), *address)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = 300000
	auth.Value = big.NewInt(0)
	gp, _ := c.SuggestGasPrice(context.Background())
	gp = gp.Mul(gp, big.NewInt(4))
	auth.GasPrice = gp

	deployAddress, _, contract, _ := aggregator.DeployAggregator(auth, c)
	log.Printf("Deployed Aggregator Contract to %s\n", deployAddress)
	_aggregator = contract
	return nil
}

func buildWriteArgs(storage *Storage) WriteArgs {
	if storage.Inverse {
		return WriteArgs{
			token0:          common.HexToAddress(storage.Quote.Address),
			token1:          common.HexToAddress(storage.Base.Address),
			token0Liquidity: storage.QuoteLiquidity,
			token1Liquidity: storage.BaseLiquidity,
			contract:        storage.ContractAddress,
		}
	} else {
		return WriteArgs{
			token0:          common.HexToAddress(storage.Base.Address),
			token1:          common.HexToAddress(storage.Quote.Address),
			token0Liquidity: storage.BaseLiquidity,
			token1Liquidity: storage.QuoteLiquidity,
			contract:        storage.ContractAddress,
		}
	}
}

func WritePriceData(bestBuy, bestSell *Storage) {
	c, _ := GetClient()
	auth, _, _, address := getAuth()

	nonce, _ := c.PendingNonceAt(context.Background(), *address)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = 300000
	auth.Value = big.NewInt(0)
	gp, _ := c.SuggestGasPrice(context.Background())
	gp = gp.Mul(gp, big.NewInt(4))
	auth.GasPrice = gp

	var bestBuyWriteArgs WriteArgs
	var bestSellWriteArgs WriteArgs

	if bestBuy != nil {
		bestBuyWriteArgs = buildWriteArgs(bestBuy)
	}
	if bestSell != nil {
		bestSellWriteArgs = buildWriteArgs(bestSell)
	}

	var err error = nil

	if bestBuy != nil && bestSell != nil {
		_, err = GetAggregator().UpdateBestAll(auth, bestBuyWriteArgs.token0, bestBuyWriteArgs.token1, &bestBuyWriteArgs.token0Liquidity, &bestBuyWriteArgs.token1Liquidity, bestBuyWriteArgs.contract)
	} else if bestBuy != nil {
		_, err = GetAggregator().UpdateBestBuy(auth, bestBuyWriteArgs.token0, bestBuyWriteArgs.token1, &bestBuyWriteArgs.token0Liquidity, &bestBuyWriteArgs.token1Liquidity, bestBuyWriteArgs.contract)
	} else if bestSell != nil {
		_, err = GetAggregator().UpdateBestSell(auth, bestSellWriteArgs.token0, bestSellWriteArgs.token1, &bestSellWriteArgs.token0Liquidity, &bestSellWriteArgs.token1Liquidity, bestSellWriteArgs.contract)
	}

	//THis is horribly hacky. Better to update everything in 1 go but I guess it will work
	if err != nil {
		WritePriceData(bestBuy, bestSell)
	}
}
