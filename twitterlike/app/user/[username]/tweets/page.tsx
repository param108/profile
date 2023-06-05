"use client";
import Header from "@/app/components/header";
import { NextApiRequest, NextApiResponse } from "next";
import { useParams, useSearchParams } from "next/navigation";
import { useEffect } from "react";

export default function ShowTweet() {
    const params = useParams();
    return (
        <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
        <Header></Header>
        <div className=" mt-[60px] bg-white">
            <p>
                {params['username']}
            </p>
        </div>
        </main>
    )
}
