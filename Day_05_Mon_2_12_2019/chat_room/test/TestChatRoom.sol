pragma solidity >=0.4.25 <0.6.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/ChatRoom.sol";

contract TestChatRoom {
  function testOpenChatRoom() public {
    ChatRoom room = ChatRoom(DeployedAddresses.ChatRoom());
    room.close();
    Assert.equal(room.isOpen(), false, "Expected ChatRoom to be closed");
    room.open();
    Assert.equal(room.isOpen(), true, "Expected ChatRoom to be opened");
  }

  function testCloseChatRoom() public {
    ChatRoom room = ChatRoom(DeployedAddresses.ChatRoom());
    room.open();
    Assert.equal(room.isOpen(), true, "Expected ChatRoom to be opened");
    room.close();
    Assert.equal(room.isOpen(), false, "Expected ChatRoom to be closed");
  }

  function testRegisterUser() public {
    ChatRoom room = ChatRoom(DeployedAddresses.ChatRoom());

    room.register("IA");
    Assert.equal(room.isRegistered("IA"), true, "ChatRoom should have a user known as 'IA' registered");
  }

  function testUnregisterUser() public {
    ChatRoom room = ChatRoom(DeployedAddresses.ChatRoom());

    room.register("IA");
    room.unregister("IA");

    Assert.equal(room.isRegistered("IA"), false, "ChatRoom should not have any users registered");
  }

  function testBroadcast() public {
    ChatRoom room = ChatRoom(DeployedAddresses.ChatRoom());
      
    room.open();

    room.register("IA");
    room.register("AI");

    room.broadcast("IA", "Hello World!");
    room.broadcast("AI", "Hello World As Well!");

    Assert.equal(room.getMessageHistorySize(), 2, "ChatRoom should have 2 messages in its message history");
  }
}