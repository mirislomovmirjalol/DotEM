
import type {ClassValue} from "clsx";
import clsxFunction from "clsx";
import {twMerge} from "tailwind-merge";

export const cn = (...inputs: ClassValue[]): string => {
    return twMerge(clsxFunction(inputs));
};