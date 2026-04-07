import type { CpuMetrics } from "../Interfaces/Metrics"
import Chart from "./Chart"

export default function Cpu({ history }: { history: CpuMetrics[] }) {
  const current = history[history.length - 1]
  if (!current) return null

  const data = history.map(m => ({ value: parseFloat(m.overall_percent.toFixed(1)) }))

  return (
    <section className="card card--full">
      <div className="card-header">
        <div className="card-title-block">
          <h2>CPU</h2>
          <span className="card-subtitle">{current.model_name}</span>
        </div>
        <span className="badge">{current.overall_percent.toFixed(1)}%</span>
      </div>

      <span className="section-label">Overall Usage</span>
      <Chart data={data} color="var(--accent)" unit="%" gradientId="cpu-grad" domain={[0, 100]} />

      <div className="card-divider" />

      <div className="card-section">
        <span className="section-label">Per Core</span>
        <div className="core-grid">
          {current.per_core_percent.map((pct, i) => (
            <div key={i} className="core-row">
              <span className="core-label">C{i}</span>
              <div className="core-bar-track">
                <div className="core-bar-fill" style={{ width: `${pct}%` }} />
              </div>
              <span className="core-pct">{pct.toFixed(0)}%</span>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}
