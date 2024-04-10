import { env } from "@/env"
import { useMutation, useQueryClient } from "@tanstack/react-query"

interface Sensor {
  id: string
}

export function useDeleteSensor() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (sensor: Sensor) => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/sensor/${sensor.id}`, {
        method: "DELETE",
      })

      const data = await resp.json()

      if (!resp.ok) {
        throw new Error(data.message)
      }

      return data
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["getSensors"],
      })
    },
  })
}
