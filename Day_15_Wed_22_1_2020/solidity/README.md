# Chat Room Contract in Solidity

## Compiling to Go

To generate the necessary ABI and BIN files:

```
solc --abi contracts/ChatRoom.sol -o ../go/server/abi --overwrite && solc --bin contracts/ChatRoom.sol -o ../go/server/abi --overwrite
```

And then translate the ABI files to a format we can work with in Go:

```
abigen --bin=../go/server/ChatRoom.bin --sol contracts/ChatRoom.sol --pkg contracts --out ../go/server/contracts/chatroom.go
```

## Upgrading the contract

```
truffle migrate --reset
```