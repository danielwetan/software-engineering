import z from "zod";

export const CreateShortUrlSchema = () =>
  z.object({
    target: z
      .string()
      .min(1, {
        message: "Target is required",
      })
      .url({
        message: "Invalid url",
      }),
  });
