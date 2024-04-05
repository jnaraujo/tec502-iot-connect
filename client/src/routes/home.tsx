import { SendCommandBox } from "@/components/send-command"
import { SensorListBox } from "@/components/sensor-list-box"
import { SensorResponseList } from "@/components/sensor-response-list"

export function Home() {
  return (
    <main className="container grid h-screen grid-cols-[1fr_400px] gap-6 py-6">
      <section className="grid h-fit grid-cols-2">
        <SendCommandBox />

        <SensorListBox />
      </section>

      <article className="bg-background h-full space-y-4 overflow-hidden rounded-lg border p-6">
        <h2 className="text-lg font-medium text-zinc-900">
          Respostas dos sensores
        </h2>
        <SensorResponseList />
      </article>
    </main>
  )
}
