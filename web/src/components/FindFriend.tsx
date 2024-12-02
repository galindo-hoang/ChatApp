import React, { useState } from 'react';
import api from '../api/api';
import '../styles/FindFriend.css';
import LocalStorageWrapper from '../utils'
import {useNavigate} from "react-router-dom";

type User = {
    id: number;
    name: string;
};

const FindFriend: React.FC<{ onFriendAdded: () => void }> = ({ onFriendAdded }) => {
    const navigate = useNavigate();
    const [query, setQuery] = useState('');
    const [searchResults, setSearchResults] = useState<User[]>([]);
    const [error, setError] = useState('');

    const handleSearch = (e: React.FormEvent) => {
        e.preventDefault();
        api.get(`/v1/users?query=${query}`)
            .then(response => {
                setSearchResults(response.data);
                setError('');
            })
            .catch(err => {
                setError('Error searching for users.');
                console.error(err);
            });
    };

    const handleAddFriend = (userId: number) => {
        const token = LocalStorageWrapper.getToken()
        if (!token) {
            navigate('/login')
            return
        }

        const configs = {
            headers: {
                Authorization: `Bearer ${token}`
            }
        }

        api.post(`/v1/friends`, { userId }, configs)
            .then(() => {
                onFriendAdded();
                alert('Friend added successfully!');
            })
            .catch(err => {
                setError('Error adding friend.');
                console.error(err);
            });
    };

    return (
        <div className="find-friend">
            <form onSubmit={handleSearch} className="find-friend-form">
                <input
                    type="text"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    placeholder="Search for friends"
                    className="search-input"
                    required
                />
                <button type="submit" className="search-button">Search</button>
            </form>
            {error && <p className="error-message">{error}</p>}
            <ul className="search-results">
                {searchResults.map(user => (
                    <li key={user.id} className="result-item">
                        {user.name}
                        <button
                            onClick={() => handleAddFriend(user.id)}
                            className="add-friend-button"
                        >
                            Add Friend
                        </button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default FindFriend;
