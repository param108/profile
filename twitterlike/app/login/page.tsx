"use client";
import Header from "@/app/components/header";
import { useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";
import { FiCloudLightning } from "react-icons/fi";
import ReactModal from "react-modal";

const largeModalStyle = {
    content: {
        left: "25%",
        right: "25%",
        top: "100px"
    }
};

const bigModalStyle = {
    content: {
        left: "10%",
        right: "10%",
        top: "100px"
    }
};

const smallModalStyle = {
    content: {
        left: "2%",
        right: "2%",
        top: "100px"
    }
};

export default function Login() {
    const [key, setKey] = useState("");
    const [error, setError] = useState("");
    const [failureVisible, setFailureVisible] = useState(false);
    
    const searchParams = useSearchParams();

    // Screen size management for modals
    const [largeScreen, setLargeScreen] = useState(
        (typeof window === "undefined") ? true : window.matchMedia("(min-width: 1024px)").matches
    );
    const [bigScreen, setBigScreen] = useState(
        (typeof window === "undefined") ? true : window.matchMedia("(min-width: 768px)").matches
    );

    useEffect(() => {
        (typeof window === "undefined") ? true : window.matchMedia("(min-width: 768px)").addEventListener('change', (e) => {
            setBigScreen(e.matches);
        });
        (typeof window === "undefined") ? true : window.matchMedia("(min-width: 1024px)").addEventListener('change', (e) => {
            setLargeScreen(e.matches);
        });
    }, []);

    // Get the key parameter from URL
    useEffect(() => {
        const keyParam = searchParams.get("key");
        if (!keyParam) {
            setError("Missing key parameter");
            setFailureVisible(true);
            return;
        }
        setKey(keyParam);
    }, [searchParams]);


    return (
        <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
            <Header changeDarkMode={null} showSpinner={false} />
            <div className="mt-[100px] md:mt-[150px] flex flex-col items-center w-full max-w-md mx-auto px-4">
                <h1 className="text-2xl font-bold text-gray-800 mb-8">Sign In</h1>
                
                <form 
                    action={`${process.env.NEXT_PUBLIC_BE_URL}/email/login?key=${key}`}
                    method="POST"
                    className="w-full space-y-4"
                >
                    <div>
                        <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-2">
                            Username
                        </label>
                        <input
                            type="text"
                            id="username"
                            name="username"
                            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                            placeholder="Enter your username"
                            pattern="[a-zA-Z0-9_]+"
                            title="Username can only contain letters, numbers, and underscores"
                            required
                        />
                    </div>
                    
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-2">
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                            placeholder="Enter your password"
                            minLength={6}
                            required
                        />
                    </div>
                    
                    {error && (
                        <div className="text-red-600 text-sm text-center">
                            {error}
                        </div>
                    )}
                    
                    <button
                        type="submit"
                        className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                        Sign In
                    </button>
                </form>
            </div>

            <ReactModal
                style={largeScreen ? largeModalStyle : (bigScreen) ? bigModalStyle : smallModalStyle}
                isOpen={failureVisible}
            >
                <div className="flex w-full h-full items-center flex-col">
                    <FiCloudLightning size={30} className="text-red-600" />
                    <p>{"Failed to load login page."}</p>
                    <p>Missing required parameters.</p>
                    <p>
                        Return <a className="text-blue-600" href="/">Home</a>!
                    </p>
                </div>
            </ReactModal>
        </main>
    );
}