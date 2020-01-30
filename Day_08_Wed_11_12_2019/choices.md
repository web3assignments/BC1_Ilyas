## Upgrade mechanism

The chat room contract is mostly stateless. Losing state after upgrading is not considered an issue as users can always rejoin the chat room.

## Pattern of choice

The pattern of choice is the self-destruct pattern. This works well with channels within the chat room that act as contracts themselves. When a channel is removed from the chat room, it can destroy itself and restore remaining Ether back to the chat room.