pragma solidity >=0.4.0 <0.7.0;

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
    }

    /**
     * The current state of the chat room.
     */
    enum State {
        OPEN, CLOSED
    }

    mapping (string => Participant) private participantsByNames;
    Participant[] private participants;
    Message[] private history;
    State private state;

    /**
     * Registers a new 'Participant' to this chat room.
     */
    function register(string memory displayName) public {
        // in every other language you'd do something like : participants.push(Participant(displayName, new Participant[]));
        // but solidity is an incredibly odd language so we'd have to instantiate it this way:

        Participant memory p = participantsByNames[displayName];
        p.displayName = displayName;
        p.index = getParticipantCount() + 1;

        participantsByNames[displayName] = p;
        participants.push(p);

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
        require(participantsByNames[displayName].index != 0, "User does not exist");

        delete participants[participantsByNames[displayName].index];
        delete participantsByNames[displayName];

        // and notify all contracts of a participant having left the chat room
        emit ParticipantLeft(displayName);
    }

    /**
     * Returns the amount of participants that are currently in this chat room.
     */
    function getParticipantCount() public view returns(uint count) {
        return participants.length;
    }
}