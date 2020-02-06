package chat

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	participantJoinedSigHash = crypto.Keccak256Hash([]byte("ParticipantJoined(string)"))
	participantLeftSigHash   = crypto.Keccak256Hash([]byte("ParticipantLeft(string)"))

	channelCreatedSigHash   = crypto.Keccak256Hash([]byte("ChannelCreated(string)"))
	channelDestroyedSigHash = crypto.Keccak256Hash([]byte("ChannelDestroyed(string)"))

	// NOTE: as we're creating a hash for every op, it is rather symbol
	// sensitive, including spaces!
	messageBroadcastedSigHash = crypto.Keccak256Hash([]byte("MessageBroadcasted(string,string,string)"))

	chatRoomOpenedSigHash = crypto.Keccak256Hash([]byte("ChatRoomOpened()"))
	chatRoomClosedSigHash = crypto.Keccak256Hash([]byte("ChatRoomClosed()"))
)

type ClientConfig struct {
	PrivateKey      string
	ContractAddress common.Address
}

type Client struct {
	config    ClientConfig
	ethClient *ethclient.Client
	chatRoom  *ChatRoom
	abi       abi.ABI
}

type ParticipantJoinedEvent struct {
	DisplayName string
}

type ParticipantLeftEvent struct {
	DisplayName string
}

type MessageBroadcasted struct {
	ChannelName string
	DisplayName string
	Text        string
}

type ChannelCreated struct {
	Name string
}

type ChannelDestroyed struct {
	Name string
}

type RoomOpened struct{}

type RoomClosed struct{}

type Event interface{}

func NewClient(config ClientConfig, ethClient *ethclient.Client, chatRoom *ChatRoom) (*Client, error) {
	contractAbi, err := abi.JSON(strings.NewReader(ChatRoomABI))
	if err != nil {
		return nil, err
	}

	return &Client{
		config:    config,
		ethClient: ethClient,
		chatRoom:  chatRoom,
		abi:       contractAbi,
	}, nil
}

func MakeNewChatRoom(client *ethclient.Client, address common.Address) (*ChatRoom, error) {
	chatRoomInstance, err := NewChatRoom(address, client)
	if err != nil {
		return nil, err
	}

	return chatRoomInstance, nil
}

func (client *Client) Consume(ctx context.Context) (Event, error) {
	eventLog, err := client.consumeEventLog(ctx, client.config.ContractAddress)
	if err != nil {
		return nil, err
	}

	eventTopic := eventLog.Topics[0]
	topicHex := eventTopic.Hex()
	switch topicHex {
	case participantJoinedSigHash.Hex():
		event := &ParticipantJoinedEvent{}
		err = client.abi.Unpack(event, "ParticipantJoined", eventLog.Data)
		if err != nil {
			return nil, err
		}

		return event, nil

	case participantLeftSigHash.Hex():
		event := &ParticipantLeftEvent{}
		err = client.abi.Unpack(event, "ParticipantLeft", eventLog.Data)
		if err != nil {
			return nil, err
		}

		return event, nil

	case messageBroadcastedSigHash.Hex():
		event := &MessageBroadcasted{}
		err = client.abi.Unpack(event, "MessageBroadcasted", eventLog.Data)
		if err != nil {
			return nil, err
		}

		return event, nil

	case channelCreatedSigHash.Hex():
		event := &ChannelCreated{}
		err = client.abi.Unpack(event, "ChannelCreated", eventLog.Data)
		if err != nil {
			return nil, err
		}

		return event, nil

	case channelDestroyedSigHash.Hex():
		event := &ChannelDestroyed{}
		err = client.abi.Unpack(event, "ChannelDestroyed", eventLog.Data)
		if err != nil {
			return nil, err
		}

		return event, nil

	case chatRoomOpenedSigHash.Hex():
		event := &RoomOpened{}
		err = client.abi.Unpack(event, "ChatRoomOpened", eventLog.Data)
		if err != nil {
			return nil, err
		}

		return event, nil

	case chatRoomClosedSigHash.Hex():
		event := &RoomClosed{}
		err = client.abi.Unpack(event, "ChatRoomClosed", eventLog.Data)
		if err != nil {
			return nil, err
		}

		return event, nil

	default:
		return nil, fmt.Errorf("unsupported event type of %v with hex %v\n", eventTopic, topicHex)
	}
}

func (client *Client) createTransactOpts(ctx context.Context, keyStr string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(keyStr)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.ethClient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.ethClient.SuggestGasPrice(ctx)
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

func (client *Client) consumeEventLog(ctx context.Context, address common.Address) (*types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(100000),
	}

	logs := make(chan types.Log)
	sub, err := client.ethClient.SubscribeFilterLogs(ctx, query, logs)
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

func (client *Client) Register(ctx context.Context, displayName string) error {
	transactOps, err := client.createTransactOpts(ctx, client.config.PrivateKey)
	if err != nil {
		return err
	}

	_, err = client.chatRoom.Register(transactOps, displayName)
	return err
}

func (client *Client) Unregister(ctx context.Context, displayName string) error {
	transactOps, err := client.createTransactOpts(ctx, client.config.PrivateKey)
	if err != nil {
		return err
	}

	_, err = client.chatRoom.Unregister(transactOps, displayName)
	return err
}

func (client *Client) OpenRoom(ctx context.Context) error {
	transactOps, err := client.createTransactOpts(ctx, client.config.PrivateKey)
	if err != nil {
		return err
	}

	_, err = client.chatRoom.Open(transactOps)
	return err
}

func (client *Client) CloseRoom(ctx context.Context) error {
	transactOps, err := client.createTransactOpts(ctx, client.config.PrivateKey)
	if err != nil {
		return err
	}

	_, err = client.chatRoom.Close(transactOps)
	return err
}

func (client *Client) Send(ctx context.Context, channelName string, displayName string, text string) error {
	transactOps, err := client.createTransactOpts(ctx, client.config.PrivateKey)
	if err != nil {
		return err
	}

	_, err = client.chatRoom.Send(transactOps, channelName, displayName, text)
	return err
}

func ConnectToChain(ctx context.Context, url string) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}

	return client, nil
}
