import z from "zod";

export const LoginSchema = () =>
  z.object({
    email: z.string().email("Invalid email").min(1, {
      message: "Email is required",
    }),
    password: z
      .string()
      .min(1, {
        message: "Password is required",
      })
      .min(8, {
        message: "Password at least 8 characters",
      }),
  });
