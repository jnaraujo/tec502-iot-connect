import { TooltipProvider } from "@/components/ui/tooltip"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { ReactQueryDevtools } from "@tanstack/react-query-devtools"
import { Toaster } from "react-hot-toast"
import { Home } from "./home"

const queryClient = new QueryClient()

export function Root() {
  return (
    <div className="flex min-h-[100svh] flex-col bg-muted font-sans">
      <QueryClientProvider client={queryClient}>
        <TooltipProvider>
          <Home />
        </TooltipProvider>

        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
      <Toaster />
    </div>
  )
}
