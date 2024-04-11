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
import { useSensorResponses } from "@/hooks/use-sensor-responses"
import { getRelativeTimeString } from "@/util/time"
import { useState } from "react"

export function List() {
  const { data } = useSensorResponses()
  const [search, setSearch] = useState("")

  const filteredData = data
    ?.filter((sensor) => {
      const term = search.toLowerCase()

      return (
        sensor.sensor_id.toLowerCase().includes(term) ||
        sensor.name.toLowerCase().includes(term) ||
        sensor.content.toLowerCase().includes(term)
      )
    })
    .sort((a, b) => {
      if(a.sensor_id < b.sensor_id) { return -1; }
      if(a.sensor_id > b.sensor_id) { return 1; }
      return 0;
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
              <TableHead className="w-[100px]">Nome</TableHead>
              <TableHead className="w-[300px]">Conteúdo</TableHead>
              <TableHead className="text-right">Atualizado em</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {filteredData?.map((sensor) => (
              <TableRow key={sensor.sensor_id}>
                <TableCell className="font-medium">
                  {sensor.sensor_id}
                </TableCell>
                <TableCell>{sensor.name}</TableCell>
                <TableCell>{sensor.content}</TableCell>
                <TableCell className="text-right">
                  {getRelativeTimeString(new Date(sensor.updated_at))}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </ScrollArea>
    </div>
  )
}
