  import { useNavigate } from "react-router";
  import { useState } from "react";
  import React from "react"; 

  function Home() {
    const navigate = useNavigate();
    const [roomInput, setRoomInput] = useState("");
    const [isRoomEntered, setRoomEntered] = useState(false);

    const [nicknameInput, setName] = useState("");
    

    const joinRoom = async () => {
      if (!roomInput.trim()) {
        alert("Please enter a room name");
        return;
      }
  
      try {
        const response = await fetch("http://localhost:8080/api/join", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ room: roomInput }),
        });
        if (!response.ok) {
          throw new Error("Failed to join room");
        }
        const data = await response.json();
        setRoomEntered(true); // set state to show nickname input
      } catch (error) {
        console.error("Error joining room:", error);
        alert("Error joining room. Try again.");
      }
    };
  
    const handleNicknameSubmit = () => {
      if (!nicknameInput.trim()) {
        alert("Please enter a nickname");
        return;
      }
      navigate(`/room?room=${encodeURIComponent(roomInput)}&nickname=${encodeURIComponent(nicknameInput)}`);
    };
  
    return (
      <div style={styles.container}>
        <h1 style={styles.title}>Fish Home Page</h1>
        <div style={styles.formContainer}>
          {!isRoomEntered ? (
            <>
              <input
                type="text"
                value={roomInput}
                onChange={(e) => setRoomInput(e.target.value)}
                style={styles.input}
                placeholder="Enter a room to get started"
              />
              <button onClick={joinRoom} style={styles.button}>
                Join Room
              </button>
            </>
          ) : (
            <>
              <input
                type="text"
                value={nicknameInput}
                onChange={(e) => setName(e.target.value)}
                style={styles.input}
                placeholder="Choose a nickname"
              />
              <button onClick={handleNicknameSubmit} style={styles.button}>
                Join Room
              </button>
            </>
          )}
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