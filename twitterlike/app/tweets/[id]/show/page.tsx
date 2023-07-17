"use client";
import Header from "@/app/components/header";
import { NextApiRequest, NextApiResponse } from "next";
import { useParams, useSearchParams } from "next/navigation";

export default function ShowTweet() {
    const params = useParams();
    return (
        <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
        <Header showSpinner={false}></Header>
        <div className=" mt-[60px] bg-white">
            <p>
                {params['id']}
            </p>  
        </div>
        </main>
    )
}
