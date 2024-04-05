import "./index.css"
import React from "react"
import ReactDOM from "react-dom/client"
import { Root } from "./routes/root"

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Root />
  </React.StrictMode>,
)
