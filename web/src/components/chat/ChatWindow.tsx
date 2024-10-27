// src/components/ChatWindow.tsx
import React, { useState } from 'react';
import MessageList from './MessageList';
import MessageInput from './MessageInput';

interface ChatWindowProps {
    friend: {
        id: string;
        name: string;
    };
}

interface Message {
    id: string;
    text: string;
    createdAt: string;
    sender: 'me' | 'friend';
}

const ChatWindow: React.FC<ChatWindowProps> = ({ friend }) => {
    const [messages, setMessages] = useState<Message[]>([]);

    const handleSendMessage = (text: string) => {
        const newMessage = {
            id: Date.now().toString(),
            text,
            createdAt: new Date().toISOString(),
            sender: 'me',
        };
        // @ts-ignore
        setMessages((prevMessages) => [...prevMessages, newMessage]);
    };

    return (
        <div className="chat-window">
            <h2>Chat with {friend.name}</h2>
            <MessageList messages={messages} />
            <MessageInput onSendMessage={handleSendMessage} />
        </div>
    );
};

export default ChatWindow;
