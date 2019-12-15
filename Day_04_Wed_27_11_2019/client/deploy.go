package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"

    store "github.com/web3assignments/BC1_Ilyas/contracts" // for demo
)

func main() {
    client, err := ethclient.Dial("https://ropsten.infura.io")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
	}

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    address, tx, instance, err := store.DeployStore(auth, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(address.Hex())   // 0x0671D425710B34f3959930d7F5b8277f00445b6C
    fmt.Println(tx.Hash().Hex()) // 0xe0ab724ce0afde47670c8dec759ae559e99380ec0b1aa53ded419190edd38ef9

    _ = instance
}