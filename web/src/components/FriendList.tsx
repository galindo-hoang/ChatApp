import React, {useEffect, useState} from 'react';
import axios from 'axios';
import '../styles/FriendList.css';

type Friend = {
    id: number;
    name: string;
    status: 'online' | 'offline';
};

const FriendList: React.FC<{ onSelectFriend: (friend: Friend) => void; refresh?: boolean }> = ({ onSelectFriend, refresh }) => {
    const [friends, setFriends] = useState<Friend[]>([]);

    useEffect(() => {
        axios.get('/api/friends')
            .then(response => {
                const sortedFriends = response.data.sort((a: Friend, b: Friend) => {
                    if (a.status === 'online' && b.status === 'offline') return -1;
                    if (a.status === 'offline' && b.status === 'online') return 1;
                    return 0;
                });
                setFriends(sortedFriends);
            })
            .catch(error => {
                console.error('Error fetching friends:', error);
            });
    }, [refresh]);

    return (
        <ul className="friend-list">
            {friends.map(friend => (
                <li key={friend.id} className={`friend-item ${friend.status}`} onClick={() => onSelectFriend(friend)}>
                    {friend.name}
                </li>
            ))}
        </ul>
    );
};

export default FriendList;
