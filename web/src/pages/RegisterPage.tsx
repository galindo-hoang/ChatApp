// RegisterPage.tsx
import React from 'react';
import Register from '../components/Register';
import '../styles/RegisterPage.css';

const RegisterPage: React.FC = () => {
    return (
        <div className="register-page">
            <h1>Register</h1>
            <Register/>
        </div>
    );
};

export default RegisterPage;
