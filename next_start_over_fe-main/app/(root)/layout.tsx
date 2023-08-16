"use client";
import { redirect } from "next/navigation";

import { options } from "../api/auth/[...nextauth]/options";
import prismadb from "@/lib/prismadb";
import { useSession } from "next-auth/react";

export default function SetupLayout({
  children
}: {
  children: React.ReactNode;
}) {
  console.log("this is layout root");
  const session = useSession({
    required: true,
    onUnauthenticated() {
      redirect("/sign-in");
    }
  });
  console.log("session?.data?.user?.email");
  console.log(session?.data?.user?.email);
  if (!session) {
    console.log("!!!!!!!!!!sessiond");
    console.log(session);
    redirect("/sign-in");
  }

  // const store = await prismadb.store.findFirst({
  //   where: {
  //     userId: session.user.
  //   }
  // });

  // if (store) {
  //   redirect(`/${store.id}`);
  // }

  return <>{children}</>;
}
