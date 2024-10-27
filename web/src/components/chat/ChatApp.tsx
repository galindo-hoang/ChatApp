// src/components/ChatApp.tsx
import React, { useState } from 'react';
import FriendsList from './FriendsList';
import ChatWindow from './ChatWindow';
import './ChatApp.css'

interface Friend {
    id: string;
    name: string;
    isOnline: boolean;
}

const ChatApp: React.FC = () => {
    const [friends] = useState<Friend[]>([
        { id: '1', name: 'Alice', isOnline: true },
        { id: '2', name: 'Bob', isOnline: true },
        { id: '3', name: 'Charlie', isOnline: false },
    ]);
    const [selectedFriend, setSelectedFriend] = useState<Friend | null>(null);

    return (
        <div className="chat-app">
            <FriendsList
                friends={friends}
                onSelectFriend={setSelectedFriend}
                selectedFriend={selectedFriend}
            />
            {selectedFriend && <ChatWindow friend={selectedFriend} />}
        </div>
    );
};

export default ChatApp;
