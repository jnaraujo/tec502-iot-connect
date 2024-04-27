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
import { cn } from "@/lib/utils"
import { Plus, Trash } from "lucide-react"
import { useState } from "react"
import { Input } from "../ui/input"

interface Props {
  onCreateSensorClick: () => void
  onDeleteSensor: (sensorId: string) => void
}

export function List(props: Props) {
  const { data } = useSensorList()
  const [search, setSearch] = useState("")

  const filteredData = data
    ?.filter((sensor) => {
      const term = search.toLowerCase()

      return (
        sensor.address.toLowerCase().includes(term) ||
        sensor.id.toLowerCase().includes(term)
      )
    })
    .sort((a, b) => {
      return a.id.localeCompare(b.id)
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
              <TableHead className="w-12 text-center"></TableHead>
              <TableHead className="w-20">ID</TableHead>
              <TableHead className="w-28">Endereço IP</TableHead>
              <TableHead className="w-12 text-center">#</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredData?.map((sensor) => (
              <TableRow key={sensor.id}>
                <TableCell>
                  <Tooltip>
                    <TooltipTrigger>
                      <span
                        className={cn(
                          "block size-3 rounded-full drop-shadow-sm",
                          {
                            "bg-green-500": sensor.is_online,
                            "bg-red-600": !sensor.is_online,
                          },
                        )}
                      />
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>
                        {sensor.is_online ? "Sensor online" : "Sensor offline"}
                      </p>
                    </TooltipContent>
                  </Tooltip>
                </TableCell>
                <TableCell
                  className="max-w-20 truncate font-medium"
                  title={sensor.id}
                >
                  {sensor.id}
                </TableCell>
                <TableCell className="max-w-24 truncate" title={sensor.address}>
                  {sensor.address}
                </TableCell>
                <TableCell className="max-w-12 text-center">
                  <button
                    onClick={() => {
                      props.onDeleteSensor(sensor.id)
                    }}
                  >
                    <Trash
                      size={16}
                      className="text-zinc-400 transition-colors duration-200 hover:text-red-700"
                    />
                  </button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </ScrollArea>
    </div>
  )
}
