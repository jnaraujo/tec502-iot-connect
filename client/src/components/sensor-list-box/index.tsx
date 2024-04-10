import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
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
import { useDeleteSensor } from "@/hooks/use-delete-sensor"
import { cn } from "@/lib/utils"
import { useState } from "react"
import { toast } from "react-hot-toast"
import { Button, buttonVariants } from "../ui/button"
import { List } from "./list"

export function SensorListBox() {
  const [openCreateDialog, setOpenCreateDialog] = useState(false)
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false)
  const [deletedSensorId, setDeletedSensorId] = useState("")
  const { mutate: createSensor, error } = useCreateSensor()
  const { mutate: deleteSensor } = useDeleteSensor()

  function addNewSensor(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const formData = new FormData(event.currentTarget)

    const address = formData.get("address") as string
    const id = formData.get("id") as string

    createSensor(
      {
        address,
        id,
      },
      {
        onSuccess: () => {
          toast.success("Sensor adicionado!")
          setOpenCreateDialog(false)
        },
      },
    )
  }

  function handleDeleteSensor() {
    deleteSensor(
      {
        id: deletedSensorId,
      },
      {
        onSuccess: () => {
          toast.success("Sensor deletado!")
        },
        onError: (error) => {
          toast.error("Erro ao deletar o sensor: " + error.message)
        },
      },
    )
  }

  return (
    <article className="flex flex-col gap-4 overflow-hidden rounded-lg border bg-background p-6">
      <h2 className="text-lg font-medium text-zinc-900">Lista de sensores</h2>

      <List
        onCreateSensorClick={() => {
          setOpenCreateDialog(true)
        }}
        onDeleteSensor={(sensorId) => {
          setOpenDeleteDialog(true)
          setDeletedSensorId(sensorId)
        }}
      />

      <Dialog open={openCreateDialog} onOpenChange={setOpenCreateDialog}>
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
                <Label htmlFor="name">ID do sensor:</Label>
                <Input
                  id="id"
                  name="id"
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

      <AlertDialog open={openDeleteDialog} onOpenChange={setOpenDeleteDialog}>
        <AlertDialogTrigger>Open</AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>
              Deletar o sensor "{deletedSensorId}"?
            </AlertDialogTitle>
            <AlertDialogDescription>
              Essa ação é irreversível. Ao deletar um sensor, ele não será mais
              acessível.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              className={buttonVariants({
                variant: "destructive",
              })}
              onClick={handleDeleteSensor}
            >
              Deletar o sensor
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </article>
  )
}
