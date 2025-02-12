import { useNavigate } from "react-router";
import { useState } from "react";
import React from "react"; 

function Home() {
  const navigate = useNavigate();
  const [input, setInput] = useState("");

  const joinRoom = async () => {
    if (!input.trim()) {
      alert("Please enter a room name");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/api/join", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ room: input }),
      });
      if (!response.ok) {
        throw new Error("Failed to join room");
      }
      const data = await response.json();
      navigate(`/room?room=${encodeURIComponent(data.room)}`);
    } catch (error) {
      console.error("Error joining room:", error);
      alert("Error joining room. Try again.");
    }
  };

  return (
    <div style={styles.container}>
      <h1 style={styles.title}>Fish Home Page</h1>
      <div style={styles.formContainer}>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          style={styles.input}
          placeholder="Enter a room to get started"
        />
        <button onClick={joinRoom} style={styles.button}>
          Join Room
        </button>
      </div>
    </div>
  );
}

export default Home;

const styles: { [key: string]: React.CSSProperties } = {
  input: {
    width: "100%",
    padding: "0.8rem",
    fontSize: "1rem",
    borderRadius: "5px",
    border: "1px solid #ccc",
    marginBottom: "1rem",
    outline: "none",
  },
};