"use client";

import * as z from "zod";
import axios from "axios";
import { signIn, useSession } from "next-auth/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toast } from "react-hot-toast";

import { useEffect, useState } from "react";

import { Modal } from "@/components/ui/modal";
import { Input } from "@/components/ui/input";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import Link from "next/link";

const signUpSchema = z.object({
  email: z.string().email(),
  password: z.string().min(6),
  name: z.string(),
  phone: z.string(),
  address: z.string()
  // I'm leaving out favoriteProductId as this is typically set post sign up, through user actions
});

const SignUpModal = () => {
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const form = useForm<z.infer<typeof signUpSchema>>({
    resolver: zodResolver(signUpSchema),
    defaultValues: {
      email: "",
      password: "",
      name: "",
      phone: "",
      address: ""
    }
  });

  const onSubmit = async (values: z.infer<typeof signUpSchema>) => {
    console.log(" this is options: api.auth onSubmit");
    try {
      setLoading(true);
      const response = await axios
        .post("/api/register", values)
        .then(() =>
          signIn("credentials", {
            ...values,
            redirect: false
          })
        )
        .then((callback) => {
          if (callback?.error) {
            toast.error("Invalid credentials!");
          }

          if (callback?.ok) {
            router.push("/");
          }
        });
    } catch (error) {
      toast.error("Failed to sign up. Please check your information.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal
      title="Sign Up"
      description="Register a new account."
      isOpen={true}
      onClose={() => {}}
    >
      <div className="space-y-4 py-2 pb-4">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email address</FormLabel>
                  <FormControl>
                    <Input disabled={loading} placeholder="email" {...field} />
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
                  <FormLabel>password</FormLabel>
                  <FormControl>
                    <Input
                      disabled={loading}
                      placeholder="password"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="phone"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Phone</FormLabel>
                  <FormControl>
                    <Input
                      disabled={loading}
                      placeholder="+1234567890"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="address"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>address</FormLabel>
                  <FormControl>
                    <Input
                      disabled={loading}
                      placeholder="+1234567890"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className="pt-6 space-x-2 flex items-center justify-end w-full">
              <Button disabled={loading} variant="outline" onClick={() => {}}>
                Cancel
              </Button>
              <Button disabled={loading} type="submit">
                Sign Up
              </Button>
            </div>
            <Link href="/sign-in">
              <span className="text-blue-500 hover:underline">Sign In</span>
            </Link>
          </form>
        </Form>
      </div>
    </Modal>
  );
};

export default function Page() {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  return <div>{isClient && <SignUpModal />}</div>;
}
