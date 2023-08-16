import { ToastProvider } from "@/providers/toast-provider";
import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { ModalProvider } from "@/providers/modal-provider";
import AuthProvider from "@/providers/auth-provider";
// import AuthProvider from "@/providers/auth-provider";
const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Admin",
  description: "Admin"
};

export default function RootLayout({
  children
}: {
  children: React.ReactNode;
}) {
  console.log("this is RootLayout");
  return (
    <html lang="en">
      <body className={inter.className}>
        <AuthProvider>
          <ToastProvider />
          <ModalProvider />
          {children}
        </AuthProvider>
      </body>
    </html>
  );
}
