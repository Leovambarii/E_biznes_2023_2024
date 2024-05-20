import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Message } from '../interfaces/Message';
import MessageInput from '../components/MessageInput';

const ChatPage: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    const startChat = async () => {
      try {
        const response = await fetch('http://localhost:8080/greet');
        let assistantMessage: Message;
        if (!response.ok) {
          assistantMessage = { message: 'Response error, sorry!', role: 'assistant' };
        } else {
          const data = await response.json();
          assistantMessage = { message: data.message, role: 'assistant' };
        }
        addMessage(assistantMessage);
      } catch (error) {
        console.error('startChat error:', error);
      }
    };

    startChat();
  }, []);

  const addMessage = (message: Message) => {
    setMessages((prevMessages) => [...prevMessages, message]);
  };

  const sendMessage = async (message: string) => {
    try {
      const userMessage: Message = { message: message, role: 'user' };
      addMessage(userMessage);

      const response = await fetch('http://localhost:8080/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ prompt: message, messages: messages }),
      });
      let assistantMessage: Message;
      if (!response.ok) {
        assistantMessage = { message: 'Response error, sorry!', role: 'assistant' };
      } else {
        const data = await response.json();
        assistantMessage = { message: data.message, role: 'assistant' };
      }
      addMessage(assistantMessage);
    } catch (error) {
      console.error('sendMessage error:', error);
    }
  };

  const closeChat = async () => {
    try {
      const response = await fetch('http://localhost:8080/close');
      let assistantMessage: Message;
      if (!response.ok) {
        assistantMessage = { message: 'Response error, sorry!', role: 'assistant' };
      } else {
        const responseData = await response.json();
        assistantMessage = { message: responseData.message, role: 'assistant' };
      }
      addMessage(assistantMessage);
      await new Promise(resolve => setTimeout(resolve, 3000));
      setMessages([]);
      navigate('/');
    } catch (error) {
      console.error('Fetch error:', error);
    }
  };

  return (
    <div>
      <h1>AI Assistant Chat</h1>
      <div>
        {messages.map((message, index) => (
          <div key={index} className={message.role}>
           {message.role === 'user' ? 'You: ' : 'Assistant: '}{message.message}
          </div>
        ))}
      </div>
      <MessageInput sendMessage={sendMessage} />
      <button onClick={closeChat}>Close Chat</button>
    </div>
  );
};

export default ChatPage;
