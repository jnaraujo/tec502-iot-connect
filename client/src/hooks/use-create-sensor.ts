import { env } from "@/env"
import { useMutation, useQueryClient } from "@tanstack/react-query"

interface NewSensor {
  address: string
  id: string
}

export function useCreateSensor() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (newSensor: NewSensor) => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/sensor`, {
        method: "POST",
        headers: {
          "content-type": "application/json",
        },
        body: JSON.stringify(newSensor),
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
