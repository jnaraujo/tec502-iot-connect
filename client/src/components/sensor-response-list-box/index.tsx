import { cn } from "@/lib/utils"
import { List } from "./list"

interface Props {
  className?: string
}
export function SensorResponseListBox(props: Props) {
  return (
    <article
      className={cn(
        "flex flex-col gap-4 overflow-hidden rounded-lg border bg-background p-6",
        props.className,
      )}
    >
      <h2 className="text-lg font-medium text-zinc-900">
        Respostas dos sensores
      </h2>
      <List />
    </article>
  )
}
