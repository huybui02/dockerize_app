"use client";

import { SessionProvider } from "next-auth/react";

export interface AuthProviderProps {
  children: React.ReactNode;
}

export default function AuthProvider({ children }: AuthProviderProps) {
  console.log("this is AuthProvider ");
  return <SessionProvider>{children}</SessionProvider>;
}
