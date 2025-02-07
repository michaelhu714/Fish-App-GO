import './App.css'
import { useWebSocket } from './hooks/useWebSocket'
import { useState } from 'react';
import React from 'react';

function App() {
  const { messages, sendMessage } = useWebSocket("ws://localhost:8080/ws");
  const [input, setInput] = useState("");
  return (
    <div>
      <h1>WebSocket Test</h1>
      <h2>Messages</h2>
      <div>
        <ul>
          {messages.map((msg: string, index: number) => (
            <li key={index}>{msg}</li>
          ))}
        </ul>
      </div>
      <div>
        <input
          type='text'
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <button onClick={() => {
          sendMessage(input);
          setInput("");
        }}>
          Send
        </button>
      </div>
    </div>
  )
}

export default App
