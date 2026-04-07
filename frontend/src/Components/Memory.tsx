import type { MemoryMetrics } from "../Interfaces/Metrics"
import Chart from "./Chart"

export default function Memory({ history }: { history: MemoryMetrics[] }) {
  const current = history[history.length - 1]
  if (!current) return null

  const usedGB = (current.used / 1e9).toFixed(1)
  const totalGB = (current.total / 1e9).toFixed(1)
  const swapUsedGB = (current.swapUsed / 1e9).toFixed(1)
  const swapTotalGB = (current.swapTotal / 1e9).toFixed(1)
  const data = history.map(m => ({ value: parseFloat(m.usedPercent.toFixed(1)) }))

  return (
    <section className="card">
      <div className="card-header">
        <div className="card-title-block">
          <h2>Memory</h2>
        </div>
        <span className="badge">{current.usedPercent.toFixed(1)}%</span>
      </div>

      <span className="section-label">Usage over time</span>
      <Chart data={data} color="#10b981" unit="%" gradientId="mem-grad" domain={[0, 100]} />

      <div className="card-divider" />

      <div className="card-section">
        <div className="mem-stats">
          <div className="mem-stat">
            <span className="section-label">RAM</span>
            <span className="mem-stat-value">{usedGB} / {totalGB} GB</span>
          </div>
          <div className="mem-stat">
            <span className="section-label">Swap</span>
            <span className="mem-stat-value">{swapUsedGB} / {swapTotalGB} GB</span>
          </div>
        </div>
      </div>
    </section>
  )
}
