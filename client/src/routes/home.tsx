import { SendCommandBox } from "@/components/send-command"
import { SensorListBox } from "@/components/sensor-list-box"
import { SensorResponseListBox } from "@/components/sensor-response-list-box"

export function Home() {
  return (
    <main className="container grid h-screen grid-cols-[1fr_400px] gap-6 py-6">
      <section className="grid h-fit grid-cols-2">
        <SendCommandBox />
      </section>

      <section className="grid h-full grid-rows-2 gap-4 overflow-auto">
        <SensorResponseListBox />

        <SensorListBox />
      </section>
    </main>
  )
}
