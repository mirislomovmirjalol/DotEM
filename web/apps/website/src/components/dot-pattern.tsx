import {useId} from "react";
import {cn} from "../lib/utils.tsx";

interface DotPatternProps {
    width?: number;
    height?: number;
    x?: number;
    y?: number;
    cx?: number;
    cy?: number;
    cr?: number;
    className?: string;

    [key: string]: unknown;
}

export default function DotPatternComponent({
                                       className,
                                       cr = 1,
                                       cy = 1,
                                       height = 16,
                                       width = 16,
                                       x = 0,
                                       y = 0,
                                       ...props
                                   }: DotPatternProps): JSX.Element {
    const id = useId();

    return (
        <svg
            aria-hidden="true"
            className={cn(
                "pointer-events-none absolute inset-0 h-full w-full fill-neutral-400/80 -z-10",
                className,
            )}
            {...props}
        >
            <defs>
                <pattern
                    height={height}
                    id={id}
                    patternContentUnits="userSpaceOnUse"
                    patternUnits="userSpaceOnUse"
                    width={width}
                    x={x}
                    y={y}
                >
                    <circle cx={cy} cy={cy} id="pattern-circle" r={cr}/>
                </pattern>
            </defs>
            <rect fill={`url(#${id})`} height="100%" strokeWidth={0} width="100%"/>
        </svg>
    );
}