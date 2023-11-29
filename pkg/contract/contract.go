package contract

import (
	"xxx/pkg/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() (*ethclient.Client, error) {
	if client != nil {
		return client, nil
	} else {
		var err error
		client, err = ethclient.Dial(config.GetConfig().Provider)
		return client, err
	}
}
