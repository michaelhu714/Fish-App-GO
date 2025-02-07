import React from "react";
import { useNavigate } from "react-router";

function Home() {
  const navigate = useNavigate();
  return (
    <div>
      <h1>Home Page</h1>
      <button onClick={() => navigate("/room")}>Go to Room</button>
    </div>
  )
}

export default Home;
