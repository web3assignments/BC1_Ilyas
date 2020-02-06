package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"

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
			event, err := chatClient.Consume(ctx)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Println(event)
		}
	}()

	r := mux.NewRouter()

	r.HandleFunc("/join", chat.JoinEndpoint(chatRoom)).
		Methods("POST")

	r.HandleFunc("/send", chat.SendEndpoint(chatRoom)).
		Methods("POST")

	if err := http.ListenAndServe(":8079", r); err != nil {
		log.Fatal(err)
	}
}
