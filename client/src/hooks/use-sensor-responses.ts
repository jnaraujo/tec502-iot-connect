import { SENSOR_RESPONSE_REFETCH_INTERVAL } from "@/constants/query"
import { env } from "@/env"
import { getRelativeTimeString } from "@/util/time"
import { useQuery } from "@tanstack/react-query"

interface Sensor {
  sensor_id: string
  name: string
  content: number[]
  created_at: string
  updated_at: string
  relative_time: string
}

export function useSensorResponses() {
  return useQuery({
    queryFn: async () => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/sensor/data`)

      if (!resp.ok) {
        throw await resp.json()
      }

      const data = (await resp.json()) as Array<Sensor>
      return data.map((sensor) => ({
        ...sensor,
        relative_time: getRelativeTimeString(new Date(sensor.updated_at)),
      }))
    },
    queryKey: ["getSensorsData"],
    refetchInterval: SENSOR_RESPONSE_REFETCH_INTERVAL,
  })
}
