import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"

export function SendCommandBox() {
  return (
    <article className="bg-background flex w-[350px] flex-col space-y-2 rounded-lg border p-6">
      <h2 className="text-lg font-medium text-zinc-900">Enviar comando</h2>

      <form className="flex flex-1 flex-col justify-between gap-4">
        <div className="flex flex-col gap-4">
          <div className="space-y-1">
            <Label htmlFor="command">Sensor ID:</Label>
            <Input
              id="command"
              name="command"
              placeholder="Ex: temp1"
              className="col-span-3"
              required
            />
          </div>

          <div className="space-y-1">
            <Label htmlFor="command">Comando:</Label>
            <Input
              id="command"
              name="command"
              placeholder="Ex: get_time"
              className="col-span-3"
              required
            />
          </div>
          <div className="space-y-1">
            <Label htmlFor="content">Conteúdo:</Label>
            <Textarea
              id="content"
              name="content"
              placeholder="Ex: today"
              className="resize-none"
            />
          </div>
        </div>

        <Button type="submit" className="w-fit">
          Enviar comando
        </Button>
      </form>
    </article>
  )
}
