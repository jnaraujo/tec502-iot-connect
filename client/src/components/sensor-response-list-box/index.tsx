import { List } from "./list"

export function SensorResponseListBox() {
  return (
    <article className="flex flex-col gap-4 overflow-hidden rounded-lg border bg-background p-6">
      <h2 className="text-lg font-medium text-zinc-900">
        Respostas dos sensores
      </h2>
      <List />
    </article>
  )
}
