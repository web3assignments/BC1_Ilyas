pragma solidity >=0.4.25 <0.6.0;

/// @title A chat room for users to participate in.
/// @author I.A Baas
/// @notice You can use this contract for only the most basic chatting.
contract ChatRoom {
    /// @notice Defines an event of a participant having joined the chat room.
    event ParticipantJoined(string displayName);

    /// @notice Defines a message broadcast event that subscribers may listen
    /// to, to trigger an action in a separate contract.
    event MessageBroadcasted(string displayName, string text);

    /// @notice Defines an event of a participant having left the chat room.
    event ParticipantLeft(string displayName);

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
        bool deleted;
    }

    /// @notice The state of the chat room.
    enum State {
        OPEN, CLOSED
    }

    mapping (string => Participant) private participantsByNames;
    Message[] private history;
    State private state = State.OPEN;

    /// @notice Registers a new 'Participant' to this chat room.
    /// @param displayName The name of the participant to display to
    /// other users.
    function register(string memory displayName) public {
        // in every other language you'd do something like : participants.push(Participant(displayName, new Participant[]));
        // but solidity is an incredibly odd language so we'd have to instantiate it this way:

        Participant memory p = participantsByNames[displayName];

        p.deleted = false;
        p.displayName = displayName;

        participantsByNames[displayName] = p;

        // and notify all contracts of a participant having joined the chat room
        emit ParticipantJoined(displayName);
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
    /// @return Whether the chat room is open or not.
    function isOpen() public view returns(bool isOpened) {
        return state == State.OPEN;
    }

    /// @notice Broadcasts the given text message as being sent by the specified user.
    /// @param displayName The name of the participant sending the message.
    /// @param text The text to broadcast.
    function broadcast(string memory displayName, string memory text) public {
        require(state != State.CLOSED, "Chat room is closed");
        require(bytes(text).length != 0, "Cannot publish empty text values");

        history.push(Message(text));
        emit MessageBroadcasted(displayName, text);
    }

    /// @notice Unregisters a 'Participant' from this chat room.
    /// @param displayName The name put up for display of a participant.
    function unregister(string memory displayName) public {
        participantsByNames[displayName].deleted = true;

        // and notify all contracts of a participant having left the chat room
        emit ParticipantLeft(displayName);
    }

    /// @notice Returns whether a user by the specified display name exists within
    /// this chat room.
    /// @return Whether the specified user exists in this chat room.
    function isRegistered(string memory displayName) public view returns(bool exists) {
        return !participantsByNames[displayName].deleted;
    }

    /// @notice Returns the amount of messages that have been published
    /// in this chat room.
    /// @return The amount of messages that have been published.
    function getMessageHistorySize() public view returns(uint count) {
        return history.length;
    }
}