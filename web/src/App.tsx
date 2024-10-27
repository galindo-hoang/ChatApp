import React from 'react';
import './App.css';
import ChatApp from "./components/chat/ChatApp";
import AuthPage from "./components/auth/AuthPage";

const App: React.FC = () => {
  return (
      <div className="app">
        <ChatApp/>
      </div>
  );
};

export default App;
