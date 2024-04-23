import { SENSOR_LIST_REFETCH_INTERVAL } from "@/constants/query"
import { env } from "@/env"
import { useQuery } from "@tanstack/react-query"

interface Sensor {
  address: string
  id: string
  is_online: boolean
}

export function useSensorList() {
  return useQuery({
    queryFn: async () => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/sensor`)

      if (!resp.ok) {
        throw new Error("Não foi possível listar os sensores")
      }

      const sensors = (await resp.json()).sensors as Array<Sensor>

      return sensors.sort((a, b) => {
        return a.id.localeCompare(b.id)
      })
    },
    queryKey: ["getSensors"],
    refetchInterval: SENSOR_LIST_REFETCH_INTERVAL,
  })
}
