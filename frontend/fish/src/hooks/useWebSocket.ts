import { useEffect, useState } from 'react';

export function useWebSocket(url: string) {
	const [ws, setWs] = useState<WebSocket | null>(null);

	useEffect(() => {
		const socket = new WebSocket(url);

		socket.onopen = () => {
			console.log("WebSocket initialized");
		}

		socket.onclose = () => {
			console.log("?? WebSocket Disconnected");
		};
	});
	return { ws }
}
