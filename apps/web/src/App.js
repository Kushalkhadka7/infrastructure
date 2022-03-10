import './App.css';

import { Button } from '@mui/material';
import { useState } from 'react';
import ReactJson from 'react-json-view';

function App() {
  const [appData, setAppData] = useState(null);
  const handleAuthClick = async () => {
    fetch('http://dev-nginx-service:80/auth/info')
      .then((response) => response.json())
      .then((data) => setAppData(data.data));
  };
  const handleNginxClick = async () => {
    fetch('http://dev-nginx-service:80')
      .then((response) => response.json())
      .then((data) => setAppData(data.data));
  };

  return (
    <div className='App'>
      <header className='App-header'>
        <Button variant='contained' color='success' onClick={handleNginxClick}>
          Call nginx
        </Button>
        <Button variant='contained' color='success' onClick={handleAuthClick}>
          Call Auth
        </Button>

        <Button className='manager' variant='contained' color='success'>
          Call Manager
        </Button>
        <div className='data'>{appData && <ReactJson src={appData} />}</div>
      </header>
    </div>
  );
}

export default App;
