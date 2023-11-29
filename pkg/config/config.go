package config

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	TokenListURI         string
	Dexes                []Dex
	Tokens               []Token
	Provider             string
	PrivateKey           string //TODO this is super unsafe and only here because this is a hackathon. IRL you'd have a key signatory that signs transactions from trusted sources
	Deploy               bool
	Batch                bool
	PriceRefreshInterval time.Duration
	AggregatorAddress    string
}

type Dex struct {
	Name            string
	ContractAddress string
	Type            DexType
}

type DexType string

const (
	UNISWAPv2 DexType = "Uniswapv2"
)

type TokenInfoDTO struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	ChainID  int    `json:"chainId"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

var config *Config

func firstOrDefault(a string, b string) string {
	if len(a) > 0 {
		return a
	} else {
		return b
	}
}

func ParseConfig() {
	config = &Config{}
	_TokenListURI := flag.String("token-list-uri", firstOrDefault(os.Getenv("TOKEN_LIST_URI"), "https://raw.githubusercontent.com/XSwapProtocol/xdc-token-list/master/testnet.tokenlist.json"), "")
	_Provider := flag.String("provider", firstOrDefault(os.Getenv("PROVIDER"), "https://rpc.apothem.network"), "")
	_PK := flag.String("private-key", firstOrDefault(os.Getenv("PRIVATE_KEY"), "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"), "") //This is the first PK from the Test Test Test recovery phrase
	_Deploy := flag.Bool("deploy", false, "")
	_Batch := flag.Bool("batch", false, "")
	_AggregatorAddress := flag.String("aggregator-address", "", "Can't be used in conjunction with deploy")
	_PriceRefreshInterval := flag.Uint64("price-refresh-interval", 6, "Time to refresh prices per pair per dex (in seconds)")

	flag.Parse()

	config.TokenListURI = *_TokenListURI
	config.Provider = *_Provider
	config.PrivateKey = *_PK
	config.Deploy = *_Deploy
	config.Batch = *_Batch
	config.PriceRefreshInterval = time.Second * time.Duration(*_PriceRefreshInterval)
	config.AggregatorAddress = *_AggregatorAddress

	config.Dexes = []Dex{{Name: "Xswap", ContractAddress: "0xCae66ac135d6489BDF5619Ae8F8f1e724765eb8f", Type: UNISWAPv2}}
	config.Tokens = getTokenInfo(config)
}

func getTokenInfo(c *Config) []Token {
	resp, err := http.Get(c.TokenListURI)
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Network traffic error")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Couldn't pull in list of assets")
	}

	// Decode the JSON response into a struct
	var data TokenInfoDTO
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Fatalln("Couldn't decode asset data")
	}
	cp := data.Tokens
	return cp
}

func GetConfig() *Config {
	return config
}
