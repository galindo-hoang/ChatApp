import React, {useState} from 'react';
import {useNavigate} from 'react-router-dom';
import '../styles/Login.css';
import api from "../api/api";
import LocalStorageWrapper from "../utils";

const Login: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const response = await api.post('/v1/sessions', {email, password});
            if (response.data.success) {
                LocalStorageWrapper.setAccount(JSON.stringify(response.data.data.account))
                LocalStorageWrapper.setToken(response.data.data.access_token)
                navigate('/chat');
            } else {
                alert(`${response.data.message}`);
            }
        } catch (error) {
            console.error('Login error:', error);
        }
    };

    const register = () => {
        navigate('/register')
    }

    return (
        <form onSubmit={handleLogin} className="login-form">
            <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="Email"
                required
                className="input-field"
            />
            <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Password"
                required
                className="input-field"
            />
            <button type="submit" className="submit-button">Login</button>
            <button onClick={register} className="register-button">Register</button>
        </form>
    );
};

export default Login;
