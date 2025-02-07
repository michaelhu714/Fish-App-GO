import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { useWebSocket } from './hooks/useWebSocket'

function App() {
  const [count, setCount] = useState(0)
  useWebSocket("ws://localhost:8080/ws");
  return (
    <>
      <div>
        <ul>
        </ul>
        <input type='text' />
        <button type='button'>Send msg</button>
      </div>
    </>
  )
}

export default App
