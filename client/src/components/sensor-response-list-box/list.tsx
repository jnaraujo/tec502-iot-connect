import { Input } from "@/components/ui/input"
import { ScrollArea } from "@/components/ui/scroll-area"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { useState } from "react"

interface Sensor {
  id: number
  sensor_id: string
  command: string
  content: string
  response: string
  created_at: string
  received_at: string
}

const data: Array<Sensor> = [
  {
    id: 0,
    sensor_id: "temp1",
    command: "get_timed",
    content: "",
    response: "Command not found",
    created_at: "2024-04-05T11:25:48.7342823-03:00",
    received_at: "2024-04-05T11:25:48.735409-03:00",
  },
  {
    id: 2,
    sensor_id: "temp1",
    command: "get_time",
    content: "test",
    response: "2024-04-05 11:26:45",
    created_at: "2024-04-05T11:26:45.2796895-03:00",
    received_at: "2024-04-05T11:26:45.2818259-03:00",
  },
  {
    id: 3,
    sensor_id: "temp1",
    command: "test",
    content: "",
    response: "Hello from sensor!",
    created_at: "2024-04-05T11:27:03.8902475-03:00",
    received_at: "2024-04-05T11:27:03.8910043-03:00",
  },
]

export function List() {
  const [search, setSearch] = useState("")

  const filteredData = data
    .filter((sensor) => {
      const term = search.toLowerCase()

      return (
        sensor.sensor_id.toLowerCase().includes(term) ||
        sensor.command.toLowerCase().includes(term) ||
        sensor.response.toLowerCase().includes(term) ||
        sensor.content.toLowerCase().includes(term)
      )
    })
    .sort((a, b) => {
      return Date.parse(b.received_at) - Date.parse(a.received_at)
    })

  return (
    <div className="flex h-full w-full flex-col space-y-2 overflow-auto p-1">
      <Input
        placeholder="Pesquisa"
        value={search}
        onChange={(event) => {
          setSearch(event.target.value)
        }}
        className="h-8 max-w-sm shrink-0"
      />

      <ScrollArea className="flex-1 pr-2">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[100px]">Sensor ID</TableHead>
              <TableHead>Comando</TableHead>
              <TableHead className="text-right">Resposta</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredData.map((sensor) => (
              <TableRow key={sensor.id}>
                <TableCell className="font-medium">
                  {sensor.sensor_id}
                </TableCell>
                <TableCell>{sensor.command}</TableCell>
                <TableCell className="text-right">{sensor.response}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </ScrollArea>
    </div>
  )
}
