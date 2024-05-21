"use client"

import {useState} from "react";

export default function CopyButton({children, value}: { children: React.ReactNode, value: string }): JSX.Element {
    const [copied, setCopied] = useState(false)

    const handleCopy = (): void => {
        void navigator.clipboard.writeText(value).then(r => r)
        setCopied(true)
    }
    return (
        <button
            onClick={() => {
                handleCopy()
                setCopied(true)
            }} type="button"
        >
            {copied ? "Copied" : children}
        </button>
    )
}