import React, { useState } from 'react';
import FriendList from '../components/FriendList';
import ChatBox from '../components/ChatBox';
import MessageList from '../components/MessageList';
import '../styles/ChatPage.css';

type Friend = {
    id: number;
    name: string;
};

const ChatPage: React.FC = () => {
    const [selectedFriend, setSelectedFriend] = useState<null | Friend>(null);

    const handleSelectFriend = (friend: Friend) => {
        setSelectedFriend(friend);
    };

    const handleSendMessage = (message: string) => {
        // Implement sending message to the server
        console.log(`Sending message to ${selectedFriend?.name}: ${message}`);
    };

    return (
        <div className="chat-page">
            <div className="friends-section">
                <FriendList onSelectFriend={handleSelectFriend} />
            </div>
            <div className="chat-section">
                {selectedFriend ? (
                    <>
                        <MessageList friendId={selectedFriend.id} />
                        <ChatBox onSendMessage={handleSendMessage} />
                    </>
                ) : (
                    <p className="select-prompt">Select a friend to start chatting.</p>
                )}
            </div>
        </div>
    );
};

export default ChatPage;
