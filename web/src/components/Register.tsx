import React, {useState} from 'react';
import {useNavigate} from 'react-router-dom';
import '../styles/Register.css';
import api from "../api/api";

const Register: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const response = await api.post('/v1/accounts', {email, password});
            if (response.data.success) {
                navigate('/login');
            } else {
                alert('Registration failed');
            }
        } catch (error) {
            console.error('Registration error:', error);
        }
    };

    return (
        <form onSubmit={handleRegister} className="register-form">
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
            <button type="submit" className="submit-button">Register</button>
        </form>
    );
};

export default Register;
