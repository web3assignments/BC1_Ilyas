pragma solidity >=0.4.25 <0.6.0;

/**
 * A smart contract that simulates a chat room where
 * participants can chat with each other.
 */
contract ChatRoom {
    /**
     * Defines an event of a participant having joined the chat room.
     */
    event ParticipantJoined(string displayName);

    /**
     * Defines a message broadcast event that subscribers may listen
     * to, to trigger an action in a separate contract.
     */
    event MessageBroadcasted(string displayName, string text);

    /**
     * Defines an event of a participant having left the chat room.
     */
    event ParticipantLeft(string displayName);

    /**
     * Defines an event of the chat room having been opened again.
     */
    event ChatRoomOpened();

    /**
     * Defines an event of the chat room having been closed again.
     */
    event ChatRoomClosed();

    /**
     * A published chat message.
     */
    struct Message {
        string text;
    }

    /**
     * A participant.
     */
    struct Participant {
        uint index;
        string displayName;
        bool deleted;
    }

    /**
     * The current state of the chat room.
     */
    enum State {
        OPEN, CLOSED
    }

    mapping (string => Participant) private participantsByNames;
    Message[] private history;
    State private state = State.OPEN;

    /**
     * Registers a new 'Participant' to this chat room.
     */
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

    /**
     * Opens up the chat room.
     */
    function open() public {
        state = State.OPEN;
        emit ChatRoomOpened();
    }

    /**
     * Closes off the chat room.
     */
    function close() public {
        state = State.CLOSED;
        emit ChatRoomClosed();
    }

    /**
     * Returns whether this chat room is open for participants or not.
     */
    function isOpen() public view returns(bool isOpened) {
        return state == State.OPEN;
    }

    /**
     * Broadcasts the given text message as being sent by the specified user.
     */
    function broadcast(string memory displayName, string memory text) public {
        require(state != State.CLOSED, "Chat room is closed");
        require(bytes(text).length != 0, "Cannot publish empty text values");

        history.push(Message(text));
        emit MessageBroadcasted(displayName, text);
    }

    /**
     * Unregisters a 'Participant' from this chat room.
     */
    function unregister(string memory displayName) public {
        participantsByNames[displayName].deleted = true;

        // and notify all contracts of a participant having left the chat room
        emit ParticipantLeft(displayName);
    }

    /**
     * Returns whether a user by the specified display name exists
     * within this chat room.
     */
    function isRegistered(string memory displayName) public view returns(bool exists) {
        return !participantsByNames[displayName].deleted;
    }

    /**
     * Returns the amount of messages that have been published in this
     * chat room.
     */
    function getMessageHistorySize() public view returns(uint count) {
        return history.length;
    }
}