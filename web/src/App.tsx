import React, {useEffect, useState} from 'react';
import './App.css';
import AuthPage from "./components/auth/AuthPage";
import http from "./configs/http";

const App: React.FC = () => {
    const [item, setItem] = useState<string>("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzNTQxMjAsInN1YiI6M30.RmQqJTw0NPML9IS9Q16-l37jZl1UmQDatAygkPUDtqQMMStlfc4KCHspHNeOnML5Gpb5nn9SZY7tHEa2mTE05g")

    useEffect(() => {
        // const data = localStorage.getItem('session_token')
        // if (data !== null) {
        //     setItem(data)
        // }
    })

    useEffect(() => {
        // if (item.length > 0) {
        const configs = {
            headers: {
                session: item,
                "Access-Control-Allow-Headers": "Origin, Content-Type, Accept",
                // "Access-Control-Allow-Origin": "https://www.example.com",
                "Access-Control-Allow-Methods": "GET, POST, OPTIONS",
                "Access-Control-Allow-Origin": "*"
            }
        }
        http.get("/sessions/verify", configs)
            .then(r =>
                alert(JSON.stringify(r))
            )
            .catch(er =>
                alert(er)
            )

        // }
    }, [item]);

    return (
        <div className="app">
            <AuthPage/>
        </div>
    );
};

export default App;
