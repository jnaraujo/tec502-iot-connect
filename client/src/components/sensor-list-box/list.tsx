import { Button } from "@/components/ui/button"
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
import { useSensorList } from "@/hooks/use-sensor-list"
import { Plus } from "lucide-react"
import { useState } from "react"
import { Input } from "../ui/input"

interface Props {
  onCreateSensorClick: () => void
}

export function List(props: Props) {
  const { data } = useSensorList()
  const [search, setSearch] = useState("")

  const filteredData = data
    ?.filter((sensor) => {
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
    <div className="flex h-full w-full flex-col space-y-2 overflow-auto p-1">
      <div className="flex shrink-0 gap-4">
        <Input
          placeholder="Pesquisa"
          value={search}
          onChange={(event) => {
            setSearch(event.target.value)
          }}
          className="h-8"
        />

        <Tooltip>
          <TooltipTrigger asChild>
            <Button
              className="size-8 shrink-0 p-0"
              onClick={props.onCreateSensorClick}
            >
              <Plus size={18} />
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Adicionar novo sensor</p>
          </TooltipContent>
        </Tooltip>
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
            {filteredData?.map((sensor) => (
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
