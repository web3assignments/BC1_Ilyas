#!/bin/bash

solc --abi ../../solidity/contracts/ChatRoom.sol -o abi --overwrite && solc --bin ../../solidity/contracts/ChatRoom.sol -o abi --overwrite

abigen --bin=../../go/abi/ChatRoom.bin --sol ../../solidity/contracts/ChatRoom.sol --pkg contracts --out contracts/chatroom.go