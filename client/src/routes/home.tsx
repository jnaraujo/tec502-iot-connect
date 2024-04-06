import { SendCommandBox } from "@/components/send-command-box"
import { SensorListBox } from "@/components/sensor-list-box"
import { SensorResponseListBox } from "@/components/sensor-response-list-box"

export function Home() {
  return (
    <main className="container grid h-screen grid-cols-[1fr_400px] gap-6 py-6">
      <section className="grid grid-rows-[120px_1fr] gap-4 overflow-auto">
        <article className="col-span-1 flex flex-col space-y-2 rounded-lg border bg-background p-6">
          <h1 className="text-xl font-medium text-zinc-900">
            Bem-vindo ao IoT Connect!
          </h1>
          <p>
            Envie comandos, receba dados e gerencia sensores através do painel.
          </p>
        </article>

        <SensorResponseListBox />
      </section>

      <section className="grid grid-rows-[350px_1fr] gap-4 overflow-auto">
        <SendCommandBox />
        <SensorListBox />
      </section>
    </main>
  )
}
