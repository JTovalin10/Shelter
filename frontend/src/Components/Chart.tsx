import { AreaChart, Area, YAxis, Tooltip, ResponsiveContainer } from 'recharts'

interface ChartPoint {
  value: number
}

interface Props {
  data: ChartPoint[]
  color: string
  unit: string
  gradientId: string
  domain?: [number | string, number | string]
}

export default function Chart({ data, color, unit, gradientId, domain = [0, 'auto'] }: Props) {
  return (
    <ResponsiveContainer width="100%" height={80}>
      <AreaChart data={data} margin={{ top: 4, right: 0, left: 0, bottom: 0 }}>
        <defs>
          <linearGradient id={gradientId} x1="0" y1="0" x2="0" y2="1">
            <stop offset="5%" stopColor={color} stopOpacity={0.35} />
            <stop offset="95%" stopColor={color} stopOpacity={0} />
          </linearGradient>
        </defs>
        <YAxis domain={domain} hide />
        <Tooltip
          content={({ active, payload }) => {
            if (active && payload?.length) {
              return (
                <div className="chart-tooltip">
                  {payload[0].value}{unit}
                </div>
              )
            }
            return null
          }}
        />
        <Area
          type="monotone"
          dataKey="value"
          stroke={color}
          strokeWidth={2}
          fill={`url(#${gradientId})`}
          dot={false}
          isAnimationActive={false}
        />
      </AreaChart>
    </ResponsiveContainer>
  )
}
