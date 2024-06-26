import React from "react";
import Link from "next/link";
import {ExternalLink} from "lucide-react";

export function WelcomeSection(): JSX.Element {
    return (
        <div className="flex flex-col items-start w-full max-w-lg">
            <h2 className="text-4xl text-white font-extrabold md:text-4xl">
                Welcome to <span className="bg-primary text-primary-foreground px-3 py-1 rounded">.EM</span>
            </h2>
            <p className="text-gray-400 md:text-lg text-sm my-4">
                Manage your environment files in local storage and export/import if needed.
                Your Dynamic Environment Management Tool
            </p>
            <div className="bg-gray-950 border-r-4 border-b-4 border-primary p-4 w-full flex flex-col items-center justify-between rounded mt-4">
                <div className="flex items-center w-full justify-between">
                    <code className="w-full">
                        Open installation guide
                    </code>
                    <Link className="border p-2 rounded flex flex-row items-center" href="/docs">
                        Open
                        <ExternalLink className="ml-2" />
                    </Link>
                </div>
            </div>
        </div>
    );
}