import { SENSOR_RESPONSE_REFETCH_INTERVAL } from "@/constants/query"
import { env } from "@/env"
import { useQuery } from "@tanstack/react-query"

interface Sensor {
  sensor_id: string
  name: string
  content: number[]
  created_at: string
  updated_at: string
}

export function useSensorResponses() {
  return useQuery({
    queryFn: async () => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/sensor/data`)

      if (!resp.ok) {
        throw await resp.json()
      }

      return (await resp.json()) as Array<Sensor>
    },
    queryKey: ["getSensorsData"],
    refetchInterval: SENSOR_RESPONSE_REFETCH_INTERVAL,
  })
}
