import React, { useState } from 'react';
import '../styles/ChatBox.css';

const ChatBox: React.FC<{ onSendMessage: (message: string) => void }> = ({ onSendMessage }) => {
    const [message, setMessage] = useState('');

    const handleSendMessage = (e: React.FormEvent) => {
        e.preventDefault();
        if (message.trim()) {
            onSendMessage(message);
            setMessage('');
        }
    };

    return (
        <form onSubmit={handleSendMessage} className="chat-box-form">
            <input
                type="text"
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                placeholder="Type a message"
                required
                className="input-field"
            />
            <button type="submit" className="send-button">Send</button>
        </form>
    );
};

export default ChatBox;
