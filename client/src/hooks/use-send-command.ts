import { env } from "@/env"
import { useMutation, useQueryClient } from "@tanstack/react-query"

interface NewCommand {
  sensorId: string
  command: string
  content?: string
}

export function useSendCommand() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (newCommand: NewCommand) => {
      const resp = await fetch(`${env.VITE_BROKER_URL}/message`, {
        method: "POST",
        headers: {
          "content-type": "application/json",
        },
        body: JSON.stringify({
          sensor_id: newCommand.sensorId,
          command: newCommand.command,
          content: newCommand.content,
        }),
      })

      const data = await resp.json()

      if (!resp.ok) {
        throw new Error(data.message)
      }

      return data
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["getSensorsData"],
      })
    },
  })
}
