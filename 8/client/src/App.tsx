import './App.css';
import React from 'react';
import { BrowserRouter, Link, Route, Routes} from 'react-router-dom';
import NotFound from './components/NotFound';
import Login from './components/Login';
import Register from './components/Register';


function App() {
  return (
    <div className="App">
      <div className="App-div">
        <BrowserRouter>
          <ul>
            <li><Link to="/">Login</Link></li>
            <li><Link to="/Register">Register</Link></li>
          </ul>
          <Routes>
            <Route path="/" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </BrowserRouter>
      </div>
    </div>
  );
}

export default App;
