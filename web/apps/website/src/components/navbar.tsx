import {Github} from "lucide-react";
import React from "react";
import Link from "next/link";

export default function Navbar(): JSX.Element {
    return (
        <nav className="fixed flex flex-row items-center justify-between w-screen p-4 border-b border-gray-800 backdrop-blur">
            <h1 className="text-primary px-3 py-1 rounded text-xl font-bold">.EM</h1>
            <div className="flex flex-row items-center space-x-4">
                <Link href="https://github.com/mirislomovmirjalol/DotEM">
                    <Github className="w-6 h-6"/>
                </Link>
                <Link className="bg-primary text-primary-foreground text-center px-4 py-2 rounded" href="/docs">
                    Get Started
                </Link>
            </div>
        </nav>
    );
}