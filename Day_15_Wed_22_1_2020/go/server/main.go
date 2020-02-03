package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	chatRoom "./contracts"
)

const (
	accountPrivateKey = "aedf2b2bbc210ec0755d66026b0fea64d2c572d1f8db962f0fba26fb07777c93"
	contractAddress   = "0x3F58c3d19aE978d2984e927e995129D2a2854a47"
)

type Event struct {
	DisplayName string
}

func main() {
	ctx := context.Background()
	client, err := connectToChain(ctx, "ws://localhost:8545/ws")
	if err != nil {
		panic(err)
	}

	hexContractAddress := common.HexToAddress(contractAddress)
	cr, err := makeNewChatRoom(client, hexContractAddress)
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(5 * time.Second)

		log.Println("Shooting in a transaction!")
		transactOpts, err := createTransactOpts(ctx, client, accountPrivateKey)
		if err != nil {
			panic(err)
		}

		tx, err := cr.Register(transactOpts, "Sino")
		if err != nil {
			panic(err)
		}

		log.Printf("Success with hash: %v\n", tx.Hash().String())
	}()

	contractAbi, err := abi.JSON(strings.NewReader(chatRoom.ChatRoomABI))
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Waiting for an event...")
		eventLog, err := consumeEventLog(ctx, client, hexContractAddress)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("New event received with block number %v\n", eventLog.BlockNumber)
		log.Printf("TX hash: %v\n", eventLog.TxHash.String())
		log.Printf("Block address: %v\n", eventLog.Address.String())

		event := &Event{}
		err = contractAbi.Unpack(event, "ParticipantJoined", eventLog.Data)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(event)
	}
}

func connectToChain(ctx context.Context, url string) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func createTransactOpts(ctx context.Context, client *ethclient.Client, privateKeyStr string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

func makeNewChatRoom(client *ethclient.Client, address common.Address) (*chatRoom.ChatRoom, error) {
	chatRoomInstance, err := chatRoom.NewChatRoom(address, client)
	if err != nil {
		return nil, err
	}

	return chatRoomInstance, nil
}

func consumeEventLog(ctx context.Context, client *ethclient.Client, address common.Address) (*types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
		FromBlock: big.NewInt(10),
		ToBlock:   big.NewInt(40),
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return nil, err
	}

	select {
	case err := <-sub.Err():
		return nil, err
	case vLog := <-logs:
		return &vLog, nil
	}
}
