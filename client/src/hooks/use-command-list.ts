import { env } from "@/env"
import { useQuery } from "@tanstack/react-query"

export function useCommandList(sensorId: string) {
  return useQuery({
    queryFn: async () => {
      if (!sensorId) {
        return []
      }

      const resp = await fetch(
        `${env.VITE_BROKER_URL}/sensor/commands/${sensorId}`,
      )

      if (!resp.ok) {
        throw new Error("Não foi possível listar os comandos")
      }

      return (await resp.json()).commands as Array<string>
    },
    queryKey: ["getSensors", sensorId],
  })
}
