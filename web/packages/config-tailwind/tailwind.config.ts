import type {Config} from "tailwindcss";

const config: Omit<Config, "content"> = {
    theme: {
        extend: {
            backgroundImage: {
                "glow-conic":
                    "conic-gradient(from 180deg at 50% 50%, #2a8af6 0deg, #a853ba 180deg, #e92a67 360deg)",
            },
            colors: {
                "primary": "#a855f7",
                "primary-foreground": "#f3f4f6",
                "secondary": "#f9fafb",
                "secondary-foreground:": "#111827",
            },
        },
    },
    plugins: [],
};
export default config;
