import { env } from "@/env"
import { useQuery } from "@tanstack/react-query"

interface Sensor {
  id: number
  sensor_id: string
  command: string
  content: string
  response: string
  created_at: string
  received_at: string
}

export function useSensorDataList() {
  return useQuery({
    queryFn: async () => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/sensor/data`)

      if (!resp.ok) {
        throw await resp.json()
      }

      return (await resp.json()) as Array<Sensor>
    },
    queryKey: ["getSensorsData"],
    refetchInterval: 1_000,
    staleTime: 5_000,
  })
}
