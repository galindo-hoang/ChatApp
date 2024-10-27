// src/components/MessageList.tsx
import React, { useEffect, useRef } from 'react';

interface Message {
    id: string;
    text: string;
    createdAt: string;
    sender: 'me' | 'friend';
}

interface MessageListProps {
    messages: Message[];
}

const MessageList: React.FC<MessageListProps> = ({ messages }) => {
    const messagesEndRef = useRef<HTMLDivElement | null>(null);

    // Scroll to the bottom when a new message is added
    useEffect(() => {
        messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
    }, [messages]);

    return (
        <div className="message-list">
            {messages.map((message) => (
                <div key={message.id} className={`message ${message.sender}`}>
                    <span>{message.text}</span>
                    <span className="time">
            {new Date(message.createdAt).toLocaleTimeString()}
          </span>
                </div>
            ))}
            {/* Invisible div to act as the scroll target */}
            <div ref={messagesEndRef} />
        </div>
    );
};

export default MessageList;
