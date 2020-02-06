package main

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/web3assignments/BC1_Ilyas/internal/chat-room/chat"
)

const (
	accountPrivateKey = "aedf2b2bbc210ec0755d66026b0fea64d2c572d1f8db962f0fba26fb07777c93"
)

var (
	contractAddress = common.HexToAddress("0xe94506240068647aF19D9Db6b4f025513B4066cc")
)

func main() {
	ctx := context.Background()
	ethClient, err := chat.ConnectToChain(ctx, "ws://localhost:8545/ws")
	if err != nil {
		panic(err)
	}

	chatRoom, err := chat.MakeNewChatRoom(ethClient, contractAddress)
	if err != nil {
		panic(err)
	}

	config := chat.ClientConfig{
		PrivateKey:      accountPrivateKey,
		ContractAddress: contractAddress,
	}

	chatClient, err := chat.NewClient(config, ethClient, chatRoom)

	go func() {
		for {
			time.Sleep(1 * time.Second)

			err := chatClient.Send(ctx, "General", "Sino", "Hello World!")
			if err != nil {
				log.Println(err)
			}
		}
	}()

	for {
		event, err := chatClient.Consume(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(event)
	}
}
