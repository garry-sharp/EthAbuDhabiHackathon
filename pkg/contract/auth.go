package contract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"xxx/pkg/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func getAuth() (*bind.TransactOpts, *ecdsa.PrivateKey, *ecdsa.PublicKey, *common.Address) {
	c, _ := GetClient()
	pk, _ := crypto.HexToECDSA(config.GetConfig().PrivateKey)
	auth := bind.NewKeyedTransactor(pk)

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Error converting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := c.PendingNonceAt(context.Background(), address)
	auth.Nonce = big.NewInt(int64(nonce))
	return auth, pk, publicKeyECDSA, &address
}
