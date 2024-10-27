// src/components/AuthPage.tsx
import React, { useState } from 'react';
import LoginForm from './LoginForm';
import RegisterForm from './RegisterForm';
import './AuthPage.css';

const AuthPage: React.FC = () => {
    const [isLogin, setIsLogin] = useState(true);

    return (
        <div className="auth-page">
            <div className="auth-toggle">
                <button onClick={() => setIsLogin(true)} disabled={isLogin}>Login</button>
                <button onClick={() => setIsLogin(false)} disabled={!isLogin}>Register</button>
            </div>
            {isLogin ? <LoginForm /> : <RegisterForm />}
        </div>
    );
};

export default AuthPage;
