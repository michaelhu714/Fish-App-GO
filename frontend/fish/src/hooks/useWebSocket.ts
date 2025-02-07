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
			const data = JSON.parse(event.data);
			console.log("Recieved Message: " + data);
			setMessages((prev) => [...prev, data]);
		}

		setWs(socket);
		return () => {
			socket.close();
		};
	}, [url]);

	const sendMessage = (type: string, content: string) => {
		if (ws && ws.readyState === WebSocket.OPEN) {
			const message = {
				type: type,
				content: content
			}
			ws.send(JSON.stringify(message));
		} else {
			console.log("WebSocket is closed");
		}
	};

	return { ws, messages, sendMessage }
}
