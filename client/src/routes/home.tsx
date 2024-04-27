import { SendCommandBox } from "@/components/send-command-box"
import { SensorListBox } from "@/components/sensor-list-box"
import { SensorResponseListBox } from "@/components/sensor-response-list-box"
import { useSensorResponses } from "@/hooks/use-sensor-responses"
import { cn } from "@/lib/utils"

export function Home() {
  const { status } = useSensorResponses()
  const isBrokerOnline = status == "success"

  return (
    <main className="container grid h-screen grid-cols-[1fr_400px] gap-6 py-6">
      <section className="grid grid-cols-[1fr_240px] grid-rows-[120px_1fr] gap-x-6 gap-y-4 overflow-auto">
        <article className="flex items-center justify-between rounded-lg border bg-background p-6">
          <div className="space-y-2">
            <h1 className="text-xl font-medium text-zinc-900">
              Bem-vindo ao IoT Connect!
            </h1>
            <p>
              Envie comandos, receba dados e gerencia sensores através do
              painel.
            </p>
          </div>
        </article>

        <article className="flex flex-col justify-center gap-2 rounded-lg border bg-background p-6">
          <h3 className="text-lg font-medium text-zinc-900">
            Status do Broker:
          </h3>
          <div className="flex items-center gap-2">
            <div
              role="none"
              className={cn("block size-3 rounded-full drop-shadow-sm", {
                "bg-green-500": isBrokerOnline,
                "bg-red-600": !isBrokerOnline,
              })}
            />
            <p className="text-zinc-700">
              {isBrokerOnline ? "Online" : "Offline"}
            </p>
          </div>
        </article>

        <SensorResponseListBox className="col-span-full" />
      </section>

      <section className="grid grid-rows-[350px_1fr] gap-4 overflow-auto">
        <SendCommandBox />
        <SensorListBox />
      </section>
    </main>
  )
}
