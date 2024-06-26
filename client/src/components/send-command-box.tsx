import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { useCommandList } from "@/hooks/use-command-list"
import { useSendCommand } from "@/hooks/use-send-command"
import { useSensorList } from "@/hooks/use-sensor-list"
import { cn } from "@/lib/utils"
import { useEffect, useState } from "react"
import toast from "react-hot-toast"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select"

export function SendCommandBox() {
  const [sensorId, setSensorId] = useState("")
  const [command, setCommand] = useState("")
  const [error, setError] = useState("")
  const { data: sensors } = useSensorList()
  const { data: commands } = useCommandList(sensorId)
  const { mutate: sendCommand } = useSendCommand()

  function handleSendCommand(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setError("")

    const formData = new FormData(event.currentTarget)
    const content = (formData.get("content") as string) || ""

    if (!sensorId || !command) {
      return setError("Selecione ID e/ou o Comando.")
    }

    sendCommand(
      {
        command,
        sensorId,
        content,
      },
      {
        onSuccess: (data) => {
          toast.success(data.message)
        },
        onError: (error) => {
          setError(error.message)
          toast.error(error.message)
        },
      },
    )
  }

  useEffect(() => {
    if (sensors?.length === 0) {
      setSensorId("")
    }

    if (commands?.length === 0) {
      setCommand("")
    }
  }, [commands?.length, sensors?.length])

  return (
    <article className="flex flex-col space-y-2 rounded-lg border bg-background p-6">
      <h2 className="text-lg font-medium text-zinc-900">Enviar comando</h2>

      <form
        className="flex flex-1 flex-col justify-between"
        onSubmit={handleSendCommand}
      >
        <div className="flex flex-col gap-2">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-0.5">
              <Label htmlFor="sensor_id">Sensor ID:</Label>
              <Select onValueChange={setSensorId} value={sensorId}>
                <SelectTrigger id="sensor_id">
                  <SelectValue placeholder="Selecione um sensor" />
                </SelectTrigger>
                <SelectContent>
                  {sensors?.map((sensor) => (
                    <SelectItem key={sensor.id} value={sensor.id}>
                      {sensor.id}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div className="space-y-0.5">
              <Label htmlFor="command">Comando:</Label>

              <Select onValueChange={setCommand} value={command}>
                <SelectTrigger id="command">
                  <SelectValue placeholder="Selecione um comando" />
                </SelectTrigger>
                <SelectContent>
                  {commands?.map((command) => (
                    <SelectItem key={command} value={command}>
                      {command}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <div className="space-y-0.5">
            <Label htmlFor="content">Conteúdo:</Label>
            <Textarea
              id="content"
              name="content"
              placeholder="Ex: today"
              className="resize-none"
            />
          </div>
        </div>

        <span
          className={cn("text-red-500 opacity-0", {
            "opacity-100": !!error,
          })}
        >
          {error}
        </span>

        <Button type="submit" className="w-fit">
          Enviar comando
        </Button>
      </form>
    </article>
  )
}
