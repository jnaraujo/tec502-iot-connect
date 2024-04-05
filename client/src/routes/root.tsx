import { TooltipProvider } from "@/components/ui/tooltip"
import { Toaster } from "react-hot-toast"
import { Home } from "./home"

export function Root() {
  return (
    <div className="bg-muted flex min-h-[100svh] flex-col font-sans">
      <TooltipProvider>
        <Home />
      </TooltipProvider>
      <Toaster />
    </div>
  )
}
