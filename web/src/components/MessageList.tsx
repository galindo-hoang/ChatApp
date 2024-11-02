import React, { useEffect, useState } from 'react';
import axios from 'axios';
import '../styles/MessageList.css';

type Message = {
    id: number;
    text: string;
    sender: string;
    timestamp: string;
};

const MessageList: React.FC<{ friendId: number }> = ({ friendId }) => {
    const [messages, setMessages] = useState<Message[]>([]);

    useEffect(() => {
        axios.get(`/api/messages?friendId=${friendId}`)
            .then(response => {
                setMessages(response.data);
            })
            .catch(error => {
                console.error('Error fetching messages:', error);
            });
    }, [friendId]);

    return (
        <ul className="message-list">
            {messages.map(message => (
                <li key={message.id} className="message-item">
                    <span className="message-sender">{message.sender}:</span>
                    <span className="message-text">{message.text}</span>
                    <span className="message-timestamp">{message.timestamp}</span>
                </li>
            ))}
        </ul>
    );
};

export default MessageList;
