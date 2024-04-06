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
import { useSensorDataList } from "@/hooks/use-sensor-data-list"
import { useState } from "react"

export function List() {
  const { data } = useSensorDataList()
  const [search, setSearch] = useState("")

  const filteredData = data
    ?.filter((sensor) => {
      const term = search.toLowerCase()

      return (
        sensor.sensor_id.toLowerCase().includes(term) ||
        sensor.command.toLowerCase().includes(term) ||
        sensor.response.toLowerCase().includes(term) ||
        sensor.content.toLowerCase().includes(term)
      )
    })
    .sort((a, b) => {
      if (!a.received_at || !b.received_at) {
        return Date.parse(b.created_at) - Date.parse(a.created_at)
      }

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
              <TableHead className="w-[100px]">Comando</TableHead>
              <TableHead className="w-[300px]">Conteúdo</TableHead>
              <TableHead className="text-right">Resposta</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredData?.map((sensor) => (
              <TableRow key={sensor.id}>
                <TableCell className="font-medium">
                  {sensor.sensor_id}
                </TableCell>
                <TableCell>{sensor.command}</TableCell>
                <TableCell>{sensor.content}</TableCell>
                <TableCell className="text-right">{sensor.response}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </ScrollArea>
    </div>
  )
}
