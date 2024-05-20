import React from 'react';
import { useNavigate } from 'react-router-dom';

const WelcomePage: React.FC = () => {
  const navigate = useNavigate();

  const handleStartChat = () => {
    navigate('/chat');
  };

  return (
    <div>
      <h1>Welcome to the our software company Chat App</h1>
      <button onClick={handleStartChat}>Start Chat</button>
    </div>
  );
};

export default WelcomePage;
