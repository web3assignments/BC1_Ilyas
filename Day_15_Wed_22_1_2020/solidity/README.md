# Chat Room Contract in Solidity

To generate the necessary ABI and BIN files:

```
solc --abi contracts/ChatRoom.sol -o ../go/abi --overwrite && solc --bin contracts/ChatRoom.sol -o ../go/abi --overwrite
```

And then translate the ABI files to a format we can work with in Go:

```
abigen --bin=../go/ChatRoom.bin --sol contracts/ChatRoom.sol --pkg contracts --out ../go/contracts/chatroom.go
```