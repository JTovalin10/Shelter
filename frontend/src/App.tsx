import { useEffect, useState } from "react";
import type { Metrics } from "./Interfaces/Metrics";
import { Cpu, Memory, Host, Network } from "./Components";
import "./App.css";

const MAX_HISTORY = 60;

export default function App() {
  const [history, setHistory] = useState<Metrics[]>([])
  const [connected, setConnected] = useState(false)

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080/ws')
    ws.onopen = () => setConnected(true)
    ws.onclose = () => setConnected(false)
    ws.onmessage = (e) => {
      const data: Metrics = JSON.parse(e.data)
      setHistory(prev => [...prev.slice(-(MAX_HISTORY - 1)), data])
    }
    return () => ws.close()
  }, [])

  const metrics = history[history.length - 1] ?? null

  if (!connected) return <div className="status">Connecting...</div>
  if (!metrics) return <div className="status">Waiting for data...</div>

  return (
    <div className="dashboard">
      <h1>Shelter</h1>
      <div className="grid">
        <Cpu history={history.map(m => m.cpu)} />
        <Memory history={history.map(m => m.memory)} />
        <Host metrics={metrics.host} />
        <Network history={history.map(m => m.network)} />
      </div>
    </div>
  )
}
