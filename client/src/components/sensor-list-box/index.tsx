import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useCreateSensor } from "@/hooks/use-create-sensor"
import { cn } from "@/lib/utils"
import { useState } from "react"
import { toast } from "react-hot-toast"
import { Button } from "../ui/button"
import { List } from "./list"

export function SensorListBox() {
  const { mutate: createSensor, error } = useCreateSensor()
  const [open, setOpen] = useState(false)

  function addNewSensor(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const formData = new FormData(event.currentTarget)

    const address = formData.get("address") as string
    const name = formData.get("name") as string

    createSensor(
      {
        address,
        name,
      },
      {
        onSuccess: () => {
          toast.success("Sensor adicionado!")
          setOpen(false)
        },
      },
    )
  }

  return (
    <article className="bg-background flex flex-col gap-4 overflow-hidden rounded-lg border p-6">
      <h2 className="text-lg font-medium text-zinc-900">Lista de sensores</h2>

      <List
        onCreateSensorClick={() => {
          setOpen(true)
        }}
      />

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="sm:max-w-[425px]">
          <form onSubmit={addNewSensor}>
            <DialogHeader>
              <DialogTitle>Adicionar novo sensor</DialogTitle>
              <DialogDescription>
                Ao adicionar um novo sensor ao sistema, ele estará disponível à
                todos os usuários e poderá ser usado para enviar e receber
                comandos.
              </DialogDescription>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-4 items-center gap-4">
                <Label htmlFor="name">Nome do sensor:</Label>
                <Input
                  id="name"
                  name="name"
                  placeholder="Ex: sensor_temperatura"
                  className="col-span-3"
                  required
                />
              </div>

              <div className="grid grid-cols-4 items-center gap-4">
                <Label htmlFor="address">Endereço ip</Label>
                <Input
                  id="address"
                  name="address"
                  placeholder="Ex: 127.0.0.1:3333"
                  className="col-span-3"
                  required
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

            <DialogFooter>
              <Button type="submit">Adicionar sensor</Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>
    </article>
  )
}
