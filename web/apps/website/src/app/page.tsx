import React from "react";
import Image from "next/image";
import {WelcomeSection} from "../components/welcome-section.tsx";
import DotPatternComponent from "../components/dot-pattern.tsx";

export default function Page(): JSX.Element {
    return (
        <main className="container mx-auto flex flex-col items-center justify-center gap-8 p-8 h-screen">
            <section className="w-full flex flex-col lg:flex-row justify-between px-4 items-center text-left">
                <WelcomeSection/>
                <Image alt="screenshot" className="mt-6" height={600} src="/screenshot.png" width={600}/>
            </section>
            <DotPatternComponent className="opacity-30" />
            <div className="absolute inset-0 m-auto max-w-xs h-[357px] blur-[118px] sm:max-w-md md:max-w-lg -z-10"
                 style={{background: "linear-gradient(106.89deg, rgba(192, 132, 252, 0.11) 15.73%, rgba(14, 165, 233, 0.41) 15.74%, rgba(232, 121, 249, 0.26) 56.49%, rgba(79, 70, 229, 0.4) 115.91%)"}}/>
        </main>
    );
}