import React from 'react';
import { useWebSocket } from '../hooks/useWebSocket'
import { useState } from 'react';

function Room() {
  const { messages, sendMessage } = useWebSocket("ws://localhost:8080/ws");
  const [input, setInput] = useState("");
  return (
    <div>
      <h1>Room</h1>
      <h2>Messages</h2>
      <div>
        <ul>
          {messages.filter(m => JSON.parse(m).type == "CHAT").map((msg: string, index: number) => (
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
          sendMessage("CHAT", input);
          setInput("");
        }}>
          Send
        </button>
      </div>
    </div>
  )
}

export default Room;
