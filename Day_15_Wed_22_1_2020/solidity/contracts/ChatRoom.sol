pragma solidity >=0.5 <0.7.0;

/// @title A chat room for users to participate in.
/// @author I.A Baas
/// @notice You can use this contract for only the most basic chatting.
contract ChatRoom {
    /// @notice Defines an event of a participant having joined the chat room.
    event ParticipantJoined(string displayName);

    /// @notice Defines an event of a participant having left the chat room.
    event ParticipantLeft(string displayName);

    /// @notice Defines an event of a channel having been created.
    event ChannelCreated(string name);

    /// @notice Defines an event of a channel having been destroyed.
    event ChannelDestroyed(string name);

    /// @notice Defines a message broadcast event that subscribers may listen
    /// to, to trigger an action in a separate contract.
    event MessageBroadcasted(string channelName, string displayName, string text);

    /// @notice Defines an event of the chat room having been opened again.
    event ChatRoomOpened();

    /// @notice Defines an event of the chat room having been closed again.
    event ChatRoomClosed();

    /// @notice A published chat message.
    struct Message {
        string text;
    }

    /// @notice A participant.
    struct Participant {
        uint index;
        string displayName;
        bool exists;
    }

    /// @notice The state of the chat room.
    enum State {
        OPEN, CLOSED
    }

    mapping (string => Participant) private participantsByNames;
    mapping (string => Channel) private channels;

    State private state = State.OPEN;

    constructor() public {
        addChannel("General");
        addChannel("Off-Topic");
    }

    /// @notice Registers a new 'Participant' to this chat room.
    /// @param displayName The name of the participant to display to
    /// other users.
    function register(string memory displayName) public payable {
        Participant memory p = participantsByNames[displayName];

        p.exists = true;
        p.displayName = displayName;

        participantsByNames[displayName] = p;

        emit ParticipantJoined(displayName);
    }

    /// @notice Adds a new 'Channel' to this chat room.
    /// @param name The name of the channel to add.
    function addChannel(string memory name) internal {
        Channel c = new Channel();
        channels[name] = c;

        emit ChannelCreated(name);
    }

    /// @notice Sends the given text message of the specified participant, into
    /// the specified chat channel.
    /// @param channelName The name of the channel to send the message to.
    /// @param displayName The name of the participant to display to
    /// other users.
    /// @param text The text message to send into the channel.
    function send(string calldata channelName, string calldata displayName, string calldata text) external {
        Channel c = channels[channelName];
        if (c.doesExist()) {
            emit MessageBroadcasted(channelName, displayName, text);
        }
    }

    /// @notice Removes an existing 'Channel' from this chat room.
    /// @param name The name of the channel to remove.
    function removeChannel(string memory name) internal {
        Channel c = channels[name];
        c.destroy();

        emit ChannelDestroyed(name);
    }

    /// @notice Opens up the chat room so that users may chatting.
    function open() public {
        state = State.OPEN;
        emit ChatRoomOpened();
    }

    /// @notice Closes off the chat room so that users can no longer chat.
    function close() public {
        state = State.CLOSED;
        emit ChatRoomClosed();
    }

    /// @notice Returns whether this chat room is open for participants or not.
    /// @return isOpened Whether the chat room is open or not.
    function isOpen() public view returns(bool isOpened) {
        return state == State.OPEN;
    }

    /// @notice Unregisters a 'Participant' from this chat room.
    /// @param displayName The name put up for display of a participant.
    function unregister(string memory displayName) public {
        participantsByNames[displayName].exists = false;

        // and notify all contracts of a participant having left the chat room
        emit ParticipantLeft(displayName);
    }

    /// @notice Returns whether a user by the specified display name exists within
    /// this chat room.
    /// @return exists Whether the specified user exists in this chat room.
    function isRegistered(string memory displayName) public view returns(bool exists) {
        return !participantsByNames[displayName].exists;
    }
}

/// @title A channel of messages.
/// @author I.A Baas
/// @notice You can use this contract for only the most basic chatting.
contract Channel {
    bool public exists;

    address payable owner;

    constructor() public {
        exists = true;
        owner = msg.sender;
    }

    /// @notice Returns whether this channel actually exists.
    /// @return Whether this channel exists in the chat room.
    function doesExist() public view returns(bool) {
        return exists;
    }

    /// @notice Marks this Channel as not existing anymore and then destroys it.
    function destroy() public {
        exists = false;
        selfdestruct(owner);
    }
}