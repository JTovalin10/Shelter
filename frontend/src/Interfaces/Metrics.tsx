export interface CpuMetrics {
  overall_percent: number;
  per_core_percent: number[];
  model_name: string;

  timestamp: number;
}

export interface MemoryMetrics {
  total: number;
  used: number;
  usedPercent: number;
  swapTotal: number;
  swapUsed: number;
}

export interface NetworkMetrics {
  bytesSent: number;
  bytesRecv: number;
  ip: string;
  mac: string;
}

export interface HostMetrics {
  uptime: number;
  procs: number;
  os: string;
}

export interface Metrics {
  cpu: CpuMetrics;
  memory: MemoryMetrics;
  network: NetworkMetrics;
  host: HostMetrics;
}