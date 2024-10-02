"use client";

import { LoginSchema } from "@/lib/validators/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { loginQuery } from "@/queries/auth";
import { profileQuery } from "@/queries/profile";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

import { Input } from "@/components/ui/input";
import { useAppDispatch } from "@/stores";
import { setData } from "@/stores/slices/auth";
// import { setDialog } from "@/stores/slices/dialog";
import { createSession } from "@/lib/session";

const LoginForm = () => {
  // const { 
  //   data: { login },
  // } = useAppSelector((state) => state.dialogReducer);
  const dispatch = useAppDispatch();

  const router = useRouter();

  const BaseLoginSchema = LoginSchema();
  const form = useForm<z.infer<typeof BaseLoginSchema>>({
    resolver: zodResolver(BaseLoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof BaseLoginSchema>) => {
    try {
      const response = await loginQuery({
        email: values?.email,
        password: values?.password,
      });
      if (response?.status === 200) {
        await createSession(response?.data?.jwt_token);
      }
      const responseProfile = await profileQuery();
      if (responseProfile.status === 200) {
        dispatch(
          setData({
            data: responseProfile?.data,
          })
        );
        // dispatch(setDialog({ login: false }));
        router.push("/dashboard");
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Card className="max-w-sm">
      <CardHeader>
        <CardTitle className="text-2xl">Login</CardTitle>
        <CardDescription>
          Enter your email below to login to your account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input placeholder="Email" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input placeholder="Password" type="password" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit">Submit</Button>
          </form>
        </Form>
      </CardContent>
    </Card>
  );
};

export default LoginForm;
