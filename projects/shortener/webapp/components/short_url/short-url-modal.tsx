"use client";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { CreateShortUrlSchema } from "@/lib/validators/short_url";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { createShortUrlQuery } from "@/queries/short_url";
import { useState } from "react";

// TODO: seperate modal and form
const ShortUrlModal = () => {
  const [open, setOpen] = useState(false);

  const BaseCreateShortUrlSchema = CreateShortUrlSchema();
  const form = useForm<z.infer<typeof BaseCreateShortUrlSchema>>({
    resolver: zodResolver(BaseCreateShortUrlSchema),
    defaultValues: {
      target: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof BaseCreateShortUrlSchema>) => {
    try {
      await createShortUrlQuery({
        target: values?.target,
      });
      setOpen(false);
      return;
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild className="mt-4">
        <Button variant="outline" onClick={() => setOpen(true)}>
          Add Link
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <DialogHeader>
              <DialogTitle>Add Link</DialogTitle>
              <DialogDescription>Add a new short url link</DialogDescription>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-1 items-center">
                <FormField
                  control={form.control}
                  name="target"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel className="text-right">Target</FormLabel>
                      <FormControl>
                        <Input
                          className="w-full"
                          placeholder="Target"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </div>
            <DialogFooter>
              <Button type="submit">Submit</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};

export default ShortUrlModal;
