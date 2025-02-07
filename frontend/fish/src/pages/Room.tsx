import { useWebSocket } from '../hooks/useWebSocket'
import { useNavigate } from "react-router";
import { useState } from 'react';
import { useLocation } from 'react-router';

function Room() {
  const params = new URLSearchParams(useLocation().search);
  const roomName = params.get("room");
  const wsURL = `ws://localhost:8080/ws?room=${encodeURIComponent(roomName!)}`
  const { messages, sendMessage } = useWebSocket(wsURL);
  const [input, setInput] = useState("");
  const navigate = useNavigate();
  return (
    <div>
      <h1>Room</h1>
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
          sendMessage("CHAT", input);
          setInput("");
        }}>
          Send
        </button>
        <button onClick={() => {
          sendMessage("LEAVE", "");
          navigate("/");
        }}>Leave</button>
      </div>
    </div>
  )
}

export default Room;
