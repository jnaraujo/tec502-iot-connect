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
    <article className="bg-background w-[350px] space-y-4 rounded-lg border p-6">
      <h2 className="text-lg font-medium text-zinc-900">Lista de sensores</h2>

      <div className="h-[300px]">
        <List
          onAddNewSensor={addNewSensor}
          open={open}
          onOpenChange={setOpen}
        />
      </div>
    </article>
  )
}
