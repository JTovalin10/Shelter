import type { HostMetrics } from "../Interfaces/Metrics"

const hour = 3600
const minute = 60

export default function Host({ metrics }: { metrics: HostMetrics }) {
  const uptimeHours = Math.floor(metrics.uptime / hour)
  const uptimeMinutes = Math.floor((metrics.uptime % hour) / minute)

  return (
    <section className="card">
      <div className="card-header">
        <div className="card-title-block">
          <h2>Host</h2>
        </div>
      </div>
      <div className="host-stats">
        <div className="host-stat">
          <span className="section-label">OS</span>
          <span className="host-stat-value">{metrics.os}</span>
        </div>
        <div className="host-stat">
          <span className="section-label">Uptime</span>
          <span className="host-stat-value">{uptimeHours}h {uptimeMinutes}m</span>
        </div>
        <div className="host-stat">
          <span className="section-label">Processes</span>
          <span className="host-stat-value">{metrics.procs}</span>
        </div>
      </div>
    </section>
  )
}
