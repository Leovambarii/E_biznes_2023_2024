import React from 'react';
import { useNavigate } from 'react-router-dom';

const WelcomePage: React.FC = () => {
  const navigate = useNavigate();

  const handleStartChat = () => {
    navigate('/chat');
  };

  return (
    <div className="Home-div">
      <h1>Welcome to Our<br/>Software Company Chat App!</h1>
      <button className="start-chat-button" onClick={handleStartChat}>Start Chat</button>
    </div>
  );
};

export default WelcomePage;
