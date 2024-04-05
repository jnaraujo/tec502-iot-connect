import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { ScrollArea } from "@/components/ui/scroll-area"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { Plus } from "lucide-react"
import { useState } from "react"

interface Sensor {
  address: string
  name: string
}

const data: Array<Sensor> = [
  {
    address: "localhost:3333",
    name: "temp1",
  },
  {
    address: "localhost:3334",
    name: "temp2",
  },
]

interface Props {
  onAddNewSensor: (event: React.FormEvent<HTMLFormElement>) => void
  open: boolean
  onOpenChange: (open: boolean) => void
}

export function List(props: Props) {
  const [search, setSearch] = useState("")

  const filteredData = data
    .filter((sensor) => {
      const term = search.toLowerCase()

      return (
        sensor.address.toLowerCase().includes(term) ||
        sensor.name.toLowerCase().includes(term)
      )
    })
    .sort((a, b) => {
      return a.name.localeCompare(b.name)
    })

  return (
    <div className="flex h-full w-full flex-col space-y-2 overflow-auto">
      <div className="flex shrink-0 gap-4">
        <Input
          placeholder="Pesquisa"
          value={search}
          onChange={(event) => {
            setSearch(event.target.value)
          }}
          className="h-8"
        />

        <Dialog open={props.open} onOpenChange={props.onOpenChange}>
          <Tooltip>
            <TooltipTrigger asChild>
              <DialogTrigger asChild>
                <Button className="size-8 shrink-0 p-0">
                  <Plus size={18} />
                </Button>
              </DialogTrigger>
            </TooltipTrigger>
            <TooltipContent>
              <p>Adicionar novo sensor</p>
            </TooltipContent>
          </Tooltip>

          <DialogContent className="sm:max-w-[425px]">
            <form onSubmit={props.onAddNewSensor}>
              <DialogHeader>
                <DialogTitle>Adicionar novo sensor</DialogTitle>
                <DialogDescription>
                  Ao adicionar um novo sensor ao sistema, ele estará disponível
                  à todos os usuários e poderá ser usado para enviar e receber
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
              <DialogFooter>
                <Button type="submit">Adicionar sensor</Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      <ScrollArea className="flex-1 pr-2">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-8">Nome</TableHead>
              <TableHead className="text-right">Endereço IP</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredData.map((sensor) => (
              <TableRow key={sensor.name}>
                <TableCell className="font-medium">{sensor.name}</TableCell>
                <TableCell className="text-right">{sensor.address}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </ScrollArea>
    </div>
  )
}
