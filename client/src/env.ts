import { z } from "zod"

const envSchema = z.object({
  VITE_BROKER_URL: z.string().url().min(1),
})

export const env = envSchema.parse(import.meta.env)
