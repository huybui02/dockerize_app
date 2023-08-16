import bcrypt from "bcrypt";
import prisma from "@/lib/prismadb";
import { NextResponse } from "next/server";

export async function POST(request: Request) {
  const body = await request.json();
  const {
    email,
    name,
    password,
    phone, // assuming the phone is sent in the request body
    address // assuming the address is sent in the request body
  } = body;

  const hashedPassword = await bcrypt.hash(password, 12);

  // Check if the email already exists in the database
  const existingUser = await prisma.user.findUnique({
    where: {
      email
    }
  });

  if (existingUser) {
    return NextResponse.error();
  }

  const user = await prisma.user.create({
    data: {
      email,
      name,
      hashedPassword,
      phone, // adding phone to the data being saved
      address // adding address to the data being saved
    }
  });

  return NextResponse.json(user);
}
