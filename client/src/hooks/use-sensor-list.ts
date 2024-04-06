import { env } from "@/env"
import { useQuery } from "@tanstack/react-query"

interface Sensor {
  address: string
  id: string
}

export function useSensorList() {
  return useQuery({
    queryFn: async () => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/sensor`)

      if (!resp.ok) {
        throw new Error("Não foi possível listar os sensores")
      }

      return (await resp.json()).sensors as Array<Sensor>
    },
    queryKey: ["getSensors"],
    refetchInterval: 10_000,
    staleTime: 5_000,
  })
}
