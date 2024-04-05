import { fontFamily } from "tailwindcss/defaultTheme"

/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  fontFamily: {
    sans: ["Inter", fontFamily.sans],
  },
  theme: {
    container: {
      center: true,
      padding: "1rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      colors: {
        muted: "hsl(var(--muted))",
        background: "hsl(var(--background))",
      },
    },
  },
  plugins: [],
}
