// pages/HomePage.tsx
import React, {useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import api from "../api/api";
import LocalStorageWrapper from "../utils";

const HomePage: React.FC = () => {
    const navigate = useNavigate();

    useEffect(() => {
        const token = LocalStorageWrapper.getToken()
        alert(`hello world`)

        if (token) {
            const config = {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            }
            api.get("/sessions/verify", config)
                .then(r => {
                    if (r.status === 202) navigate("/chat")
                    else navigate("/login")
                })
                .catch(err => {
                    navigate("/login")
                })
        } else {
            // Otherwise, navigate to the login page
            navigate('/login');
        }
    }, [navigate]);

    // @ts-ignore
    return (
        <div>
            <h1>Home Page</h1>
            <p>Checking user status...</p>
        </div>
    );
};

export default HomePage;
