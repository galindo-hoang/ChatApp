// src/components/FriendsList.tsx
import React from 'react';

interface Friend {
    id: string;
    name: string;
    isOnline: boolean;
}

interface FriendsListProps {
    friends: Friend[];
    onSelectFriend: (friend: Friend) => void;
    selectedFriend: Friend | null;
}

const FriendsList: React.FC<FriendsListProps> = ({
                                                     friends,
                                                     onSelectFriend,
                                                     selectedFriend,
                                                 }) => {
    return (
        <div className="friends-list">
            <h2>Online Friends</h2>
            <ul>
                {friends.map((friend) => (
                    <li
                        key={friend.id}
                        className={`friend-item ${friend.isOnline ? 'online' : ''} ${
                            selectedFriend?.id === friend.id ? 'selected' : ''
                        }`}
                        onClick={() => onSelectFriend(friend)}
                    >
                        {friend.name} {friend.isOnline ? '🟢' : '⚪'}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default FriendsList;
