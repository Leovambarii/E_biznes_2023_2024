import React from 'react';
import { BrowserRouter, Route, Routes} from 'react-router-dom';
import WelcomePage from './components/WelcomePage';
import ChatPage from './components/ChatPage';
import NotFoundPage from "./components/NotFoundPage";

const App: React.FC = () => {
  return (
      <div className="App">
        <div className="App-div">
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<WelcomePage />} />
              <Route path="/chat" element={<ChatPage />} />
              <Route path="*" element={<NotFoundPage />} />
            </Routes>
          </BrowserRouter>
        </div>
      </div>
  );
};

export default App;
