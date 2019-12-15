const ChatR = artifacts.require("ChatRoom");

module.exports = function(deployer) {
  deployer.deploy(ChatR);
};
