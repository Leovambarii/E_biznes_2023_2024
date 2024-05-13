import React, { useState } from 'react';

function Register() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [response, setResponse] = useState('');
  const [token, setToken] = useState(localStorage.getItem('token'));

  const handleRegister = async () => {
    try {
      const res = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({username, password}),
      });

      if (res.ok) {
        setResponse(`Registration successful!`);
      } else {
        const responseText = await res.text();
        setResponse(responseText);
      }
    } catch (error) {
      console.log(error);
      setResponse(`Registration error!`);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    setToken(null);
    setResponse('');
  };


  return (
    <div>
      {token ? (
        <div>
          <h2>Logged In</h2>
          <button onClick={handleLogout}>Logout</button>
        </div>
      ) : (
        <div>
          <h2>Register</h2>
          <input type="text" placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
          <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
          <button onClick={handleRegister}>Register</button>
        </div>
      )}
      {response && <p style={{ color: response.includes('successful') ? 'green' : 'red' }}>{response}</p>}
    </div>
  );
}

export default Register;
