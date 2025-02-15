import { useWebSocket } from '../hooks/useWebSocket';
import { useNavigate } from "react-router";
import { useState } from 'react';
import { useLocation } from 'react-router';

function Room() {
  const params = new URLSearchParams(useLocation().search);
  const roomName = params.get("room");
  const nickname = params.get("nickname");
  const wsURL = `ws://localhost:8080/ws?room=${encodeURIComponent(roomName!)}&nickname=${encodeURIComponent(nickname!)}`;
  const { messages, sendMessage } = useWebSocket(wsURL);
  const [input, setInput] = useState("");
  const navigate = useNavigate();

  const handleSendMessage = () => {
    if (input.trim()) {
      sendMessage("CHAT", input);
      setInput("");
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleSendMessage();
    }
  };

  return (
    <div style={styles.container}>
      <h1 style={styles.title}>Room: {roomName}</h1>
      <h2 style={styles.name}>You: {nickname}</h2> 
      <div style={styles.messagesContainer}>
        <h2 style={styles.subtitle}>Messages</h2>
        <ul style={styles.messagesList}>
          {messages.map((msg:string, index:number) => (
            <li key={index} style={styles.messageItem}>{msg}</li>
          ))}
        </ul>
      </div>
      <div style={styles.inputContainer}>
        <input
          type='text'
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={handleKeyDown} 
          style={styles.input}
          placeholder="Type a message..."
        />
        <button
          onClick={() => {
            sendMessage("LEAVE", "");
            navigate("/");
          }}
          style={{ ...styles.button, ...styles.leaveButton }}
        >
          Leave Room
        </button>
      </div>
    </div>
  );
}

export default Room;

// Inline styles (same as before)
const styles: { [key: string]: React.CSSProperties } = {
  title: {
    fontSize: "2rem",
    marginBottom: "20px",
    marginRight: "50px",
  },
  subtitle: {
    fontSize: "1.5rem",
    color: "#555",
    marginBottom: "10px",
  },
  messagesContainer: {
    backgroundColor: "#fff",
    padding: "20px",
    borderRadius: "10px",
    boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
    width: "80%",
    maxWidth: "600px",
    height: "400px",
    overflowY: "auto",
    marginBottom: "20px",
  },
  messagesList: {
    listStyle: "none",
    padding: "0",
    margin: "0",
  },
  messageItem: {
    padding: "10px",
    borderBottom: "1px solid #eee",
    fontSize: "1rem",
    color: "#333",
  },
  inputContainer: {
    display: "flex",
    gap: "10px",
    width: "80%",
    maxWidth: "600px",
  },
  input: {
    flex: 1,
    padding: "10px",
    fontSize: "1rem",
    borderRadius: "5px",
    border: "1px solid #ccc",
    outline: "none",
  },
  leaveButton: {
    backgroundColor: "#dc3545",
    marginLeft: "5px"
  },
  name: {
    marginRight: "45px"
  }
};