import React, { useState } from 'react';
import ChatBox from '../components/ChatBox';
import FriendList from '../components/FriendList';
import FindFriend from '../components/FindFriend';
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
        console.log(`Sending message to ${selectedFriend?.name}: ${message}`);
    };

    const refreshFriendList = () => {
        // Trigger a refresh of the friend list in the FriendList component
    };

    return (
        <div className="chat-page">
            <div className="friends-section">
                <FindFriend onFriendAdded={refreshFriendList} />
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
