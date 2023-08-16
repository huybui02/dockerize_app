import bcrypt from "bcrypt";

import type { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import prisma from "@/lib/prismadb";

export const options: NextAuthOptions = {
  providers: [
    // GitHubProvider...
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        email: { label: "Email", type: "email" },
        password: { label: "Password", type: "password" }
      },
      async authorize(credentials) {
        console.log(" this is options: NextAuthOptions");
        if (!credentials || !credentials.email || !credentials.password) {
          return null; // Incomplete credentials provided
        }

        const user = await prisma.user.findUnique({
          where: {
            email: credentials.email
          }
        });

        if (!user || !user.hashedPassword) {
          throw new Error("Invalid credentials");
        }

        const isCorrectPassword = await bcrypt.compare(
          credentials.password,
          user.hashedPassword
        );

        if (!isCorrectPassword) {
          throw new Error("Invalid credentials");
        }

        // If the credentials are valid, return the user data
        return {
          id: user.id,
          email: user.email,
          image: user.image || undefined, // optional field
          name: user.name || undefined, // optional field
          phone: user.phone,
          address: user.address
          // ... other user properties if necessary
        };
      }
    })
  ]
  // ... other options
};
