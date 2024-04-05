import { useState } from "react"
import { toast } from "react-hot-toast"
import { List } from "./list"

export function SensorListBox() {
  const [open, setOpen] = useState(false)

  function addNewSensor(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const formData = new FormData(event.currentTarget)

    const address = formData.get("address")
    const name = formData.get("name")

    toast.success("Sensor adicionado!")
    setOpen(false)

    console.log(address, name)
  }

  return (
    <article className="bg-background flex flex-col gap-4 overflow-hidden rounded-lg border p-6">
      <h2 className="text-lg font-medium text-zinc-900">Lista de sensores</h2>

      <List onAddNewSensor={addNewSensor} open={open} onOpenChange={setOpen} />
    </article>
  )
}
