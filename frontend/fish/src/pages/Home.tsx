import { useNavigate } from "react-router";
import { useState } from "react";

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
  }

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
      <button onClick={joinRoom}>Go to Room</button>
    </div >
  )
}

export default Home;
