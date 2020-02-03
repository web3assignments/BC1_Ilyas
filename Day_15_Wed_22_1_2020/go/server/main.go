package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"reflect"
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

var (
	participantJoinedSigHash = crypto.Keccak256Hash([]byte("ParticipantJoined(string)"))
	participantLeftSigHash   = crypto.Keccak256Hash([]byte("ParticipantLeft(string)"))

	channelCreatedSigHash   = crypto.Keccak256Hash([]byte("ChannelCreated(string)"))
	channelDestroyedSigHash = crypto.Keccak256Hash([]byte("ChannelDestroyed(string)"))

	chatRoomOpenedSigHash = crypto.Keccak256Hash([]byte("ChatRoomOpened()"))
	chatRoomClosedSigHash = crypto.Keccak256Hash([]byte("ChatRoomClosed"))
)

type ParticipantJoinedEvent struct {
	DisplayName string
}

type ParticipantLeftEvent struct {
	DisplayName string
}

type ChannelCreated struct {
	Name string
}

type ChannelDestroyed struct {
	Name string
}

type ChatRoomOpened struct{}

type ChatRoomClosed struct{}

type Event interface{}

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

		transactOpts, err := createTransactOpts(ctx, client, accountPrivateKey)
		if err != nil {
			panic(err)
		}

		_, err = cr.Register(transactOpts, "Sino")
		if err != nil {
			panic(err)
		}

		transactOpts, err = createTransactOpts(ctx, client, accountPrivateKey)
		if err != nil {
			panic(err)
		}

		_, err = cr.Register(transactOpts, "Sino1")
		if err != nil {
			panic(err)
		}

		transactOpts, err = createTransactOpts(ctx, client, accountPrivateKey)
		if err != nil {
			panic(err)
		}

		_, err = cr.Register(transactOpts, "Sino2")
		if err != nil {
			panic(err)
		}
	}()

	contractAbi, err := abi.JSON(strings.NewReader(chatRoom.ChatRoomABI))
	if err != nil {
		log.Fatal(err)
	}

	eventChannel := make(chan Event)
	go handleEvents(eventChannel)

	for {
		log.Println("Waiting for an event...")
		eventLog, err := consumeEventLog(ctx, client, hexContractAddress)
		if err != nil {
			log.Println(err)
			continue
		}

		eventTopic := eventLog.Topics[0]
		topicHex := eventTopic.Hex()
		switch topicHex {
		case participantJoinedSigHash.Hex():
			event := &ParticipantJoinedEvent{}
			err = contractAbi.Unpack(event, "ParticipantJoined", eventLog.Data)
			if err != nil {
				log.Println(err)
				continue
			}

			eventChannel <- event

		case participantLeftSigHash.Hex():
			event := &ParticipantLeftEvent{}
			err = contractAbi.Unpack(event, "ParticipantLeft", eventLog.Data)
			if err != nil {
				log.Println(err)
				continue
			}

			eventChannel <- event

		case channelCreatedSigHash.Hex():
			event := &ChannelCreated{}
			err = contractAbi.Unpack(event, "ChannelCreated", eventLog.Data)
			if err != nil {
				log.Println(err)
				continue
			}

			eventChannel <- event

		case channelDestroyedSigHash.Hex():
			event := &ChannelDestroyed{}
			err = contractAbi.Unpack(event, "ChannelDestroyed", eventLog.Data)
			if err != nil {
				log.Println(err)
				continue
			}

			eventChannel <- event

		case chatRoomOpenedSigHash.Hex():
			event := &ChatRoomOpened{}
			err = contractAbi.Unpack(event, "ChatRoomOpened", eventLog.Data)
			if err != nil {
				log.Println(err)
				continue
			}

			eventChannel <- event

		case chatRoomClosedSigHash.Hex():
			event := &ChatRoomClosed{}
			err = contractAbi.Unpack(event, "ChatRoomClosed", eventLog.Data)
			if err != nil {
				log.Println(err)
				continue
			}

			eventChannel <- event

		default:
			log.Printf("unsupported event type of %v with hex %v\n", eventTopic, topicHex)
		}
	}
}

func handleEvents(eventChannel <-chan Event) {
	for evt := range eventChannel {
		switch evt {
		default:
			log.Printf("unsupported event of type %v\n", reflect.TypeOf(evt))
		}
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
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(100000),
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
