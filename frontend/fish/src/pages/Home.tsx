import React from "react";
import { useNavigate } from "react-router";
import { useWebSocket } from '../hooks/useWebSocket'
import { useState } from "react";

function Home() {
  const navigate = useNavigate();
  const { sendMessage } = useWebSocket("ws://localhost:8080/ws");
  const [input, setInput] = useState("");
  return (
    <div>
      <h1>Home Page</h1>
      <div>
        <p>Enter room name</p>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)} />
      </div>
      <button onClick={() => {
        sendMessage("JOIN", input);
        navigate("/room");
      }}>Go to Room</button>
    </div >
  )
}

export default Home;
