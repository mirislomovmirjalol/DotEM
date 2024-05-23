import "./globals.css";
import "@repo/ui/styles.css";
import type {Metadata} from "next";
import {Inter} from "next/font/google";
import React from "react";
import Navbar from "../components/navbar.tsx";

const inter = Inter({subsets: ["latin"]});

export const metadata: Metadata = {
    title: "Create Turborepo",
    description: "Generated by create turbo",
};

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode;
}): JSX.Element {
    return (
        <html lang="en">
        <body className={`${inter.className} bg-gray-950 text-white min-h-screen`}>
        <Navbar/>
        {children}
        </body>
        </html>
    );
}