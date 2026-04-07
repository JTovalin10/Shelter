import type { NetworkMetrics } from "../Interfaces/Metrics"
import Chart from "./Chart"

export default function Network({ history }: { history: NetworkMetrics[] }) {
  const current = history[history.length - 1]
  if (!current) return null

  const sentData = history.map((m, i) => {
    if (i === 0) return { value: 0 }
    const delta = (m.bytesSent - history[i - 1].bytesSent) / 1024
    return { value: Math.max(0, parseFloat(delta.toFixed(1))) }
  })

  const recvData = history.map((m, i) => {
    if (i === 0) return { value: 0 }
    const delta = (m.bytesRecv - history[i - 1].bytesRecv) / 1024
    return { value: Math.max(0, parseFloat(delta.toFixed(1))) }
  })

  const currentSentRate = sentData[sentData.length - 1]?.value ?? 0
  const currentRecvRate = recvData[recvData.length - 1]?.value ?? 0
  const totalSentMB = (current.bytesSent / 1e6).toFixed(1)
  const totalRecvMB = (current.bytesRecv / 1e6).toFixed(1)

  return (
    <section className="card card--full">
      <div className="card-header">
        <div className="card-title-block">
          <h2>Network</h2>
          <span className="card-subtitle">{current.ip} &nbsp;·&nbsp; {current.mac}</span>
        </div>
        <div className="net-header-right">
          <span className="net-total">↑ {totalSentMB} MB sent</span>
          <span className="net-total">↓ {totalRecvMB} MB received</span>
        </div>
      </div>

      <div className="net-charts">
        <div>
          <div className="net-chart-header">
            <span className="section-label">↑ Upload</span>
            <span className="badge-sm">{currentSentRate} KB/s</span>
          </div>
          <Chart data={sentData} color="#f59e0b" unit=" KB/s" gradientId="net-sent-grad" />
        </div>
        <div>
          <div className="net-chart-header">
            <span className="section-label">↓ Download</span>
            <span className="badge-sm">{currentRecvRate} KB/s</span>
          </div>
          <Chart data={recvData} color="#3b82f6" unit=" KB/s" gradientId="net-recv-grad" />
        </div>
      </div>
    </section>
  )
}
