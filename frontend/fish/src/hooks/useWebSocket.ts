import { useEffect, useState } from 'react';

export function useWebSocket(url: string) {
	const [ws, setWs] = useState<WebSocket | null>(null);
	const [messages, setMessages] = useState<string[]>([]);

	useEffect(() => {
		const socket = new WebSocket(url);

		socket.onopen = () => {
			console.log("WebSocket initialized");
		}

		socket.onclose = () => {
			console.log("?? WebSocket Disconnected");
		};

		socket.onmessage = (event) => {
			console.log("Recieved Message: " + event.data);
			setMessages((prev) => [...prev, event.data]);
		}

		setWs(socket);
		return () => {
			socket.close();
		};
	}, [url]);

	const sendMessage = (msg: string) => {
		if (ws && ws.readyState === WebSocket.OPEN) {
			ws.send(msg);
		} else {
			console.log("WebSocket is closed");
		}
	};

	return { ws, messages, sendMessage }
}
