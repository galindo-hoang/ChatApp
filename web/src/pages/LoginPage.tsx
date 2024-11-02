// LoginPage.tsx
import React from 'react';
import Login from '../components/Login';
import '../styles/LoginPage.css';

const LoginPage: React.FC = () => {
    return (
        <div className="login-page">
            <h1>Login</h1>
            <Login/>
        </div>
    );
};

export default LoginPage;
