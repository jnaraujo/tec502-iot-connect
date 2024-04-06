import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { useSendCommand } from "@/hooks/use-send-command"
import { useSensorList } from "@/hooks/use-sensor-list"
import { cn } from "@/lib/utils"
import { useState } from "react"
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
  const { mutate: sendCommand, error } = useSendCommand()
  const { data: sensors } = useSensorList()

  function handleSendCommand(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const formData = new FormData(event.currentTarget)

    const command = (formData.get("command") as string) || ""
    const content = (formData.get("content") as string) || ""

    sendCommand(
      {
        command,
        sensorId,
        content,
      },
      {
        onSuccess: () => {
          toast.success("Comando enviado com sucesso")
        },
      },
    )
  }

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
              <Select onValueChange={setSensorId} defaultValue={sensorId}>
                <SelectTrigger id="sensor_id">
                  <SelectValue placeholder="Sensor" />
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
              <Input
                id="command"
                name="command"
                placeholder="Ex: get_time"
                className="col-span-3"
                required
              />
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
          {error?.message}
        </span>

        <Button type="submit" className="w-fit">
          Enviar comando
        </Button>
      </form>
    </article>
  )
}
