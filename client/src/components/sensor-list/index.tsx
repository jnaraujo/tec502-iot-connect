import { List } from "./list"

export function SensorList() {
  return (
    <article className="bg-background w-[350px] space-y-4 rounded-lg border p-6">
      <h2 className="text-lg font-medium text-zinc-900">Lista de sensores</h2>

      <div className="h-[300px]">
        <List />
      </div>
    </article>
  )
}
