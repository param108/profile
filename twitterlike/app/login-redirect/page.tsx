"use client";
import Header from "@/app/components/header";
import { NextApiRequest, NextApiResponse } from "next";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";
import { FiCloudLightning } from "react-icons/fi";
import ReactModal from "react-modal";
import { getOneTime, getProfile } from "../apis/login";
import CircularProgressBar from "../components/circular_progress_bar";
import dynamic from 'next/dynamic';

const ProgressBar = dynamic(() =>
    import("../components/circular_progress_bar"), {
        ssr: false,
    });

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

// This page takes the onetime parameter and the redirecturl
// from the url will extract the token and store in browser storage
// and redirect to the redirecturl
export default function LoginUser() {
    var [ oneTime, setOneTime ] = useState("")
    var [ token, setToken ] = useState("")
    var [ failureVisible, setFailureVisible ] = useState(false)
    var [ tokenWritten, setTokenWritten ] = useState(false)
    var [ redirectURL, setRedirectURL ] = useState("")
    var [ progress, setProgress ] = useState(1);
    const router = useRouter();
    const searchParams = useSearchParams();

    // screen prelims for modal
    const [largeScreen, setLargeScreen] = useState(
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 1024px)").matches
    )
    const [bigScreen, setBigScreen] = useState(
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 768px)").matches
    )

    useEffect(() => {
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 768px)").addEventListener('change', (e) => {
            setBigScreen( e.matches );
        });
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 1024px)").addEventListener('change', (e) => {
            setLargeScreen( e.matches );
        });
    }, []);

        // check if the parameters passed are valid
    useEffect(()=>{
        const onetimeVal :(string|null) = searchParams.get("onetime");

        const oneTimeValue:string = (onetimeVal?onetimeVal:"");

        const redirectVal: (string|null) = searchParams.get("redirect_url")
        const redirectURLValue:string = (redirectVal?redirectVal:"");

        if (oneTimeValue === "") {
            setFailureVisible(true);
            return;
        }

        if (redirectURLValue === "" || !redirectURLValue.startsWith("/")) {
            setFailureVisible(true);
            return;
        }

        setRedirectURL(redirectURLValue);

        setOneTime(oneTimeValue);

        setProgress(3)
    }, [searchParams]);

    // use onetime to get the token
    useEffect(()=>{
        if (oneTime == "") {
            return;
        }

        getOneTime(oneTime)
            .then((p)=> {
                setToken(p.data.data.value)
                setProgress(5)
            })
            .catch(() => setFailureVisible(true));

    }, [oneTime])

    // store the token in localstorage
    useEffect(()=>{
        if (token == "") {
            return;
        }
        try {
            getProfile(token).
                then((res)=>{
                    localStorage.setItem('username', res.data.data.username)
                    localStorage.setItem('user_id', res.data.data.user_id)
                    localStorage.setItem('api_token', token);
                    setProgress(6)
                    setTokenWritten(true)
                }).
                catch(()=>{
                    // clear out the api_token
                    localStorage.removeItem('api_token')
                    setFailureVisible(true)
                })

        } catch(e) {
            setFailureVisible(true)
        }
    }, [token])

    // all done, redirect to the redirectURL
    useEffect(()=>{
        if (!tokenWritten) {
            return;
        }

        router.push(redirectURL)

    }, [tokenWritten, redirectURL, router])

    return (
        <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
        <Header showSpinner={false}></Header>
        <div className="mt-[100px] md:mt-[150px] flex flex-col items-center">
            <span className="text-gray-500 md:text-xl">Logging you in...</span>
            <ProgressBar
                maxValue={8}
                selectedValue={progress}
                radius={largeScreen?180:bigScreen?120:80}
                strokeWidth={largeScreen?12:bigScreen?10:6}
                label=''
                activeStrokeColor='#05a168'
                inactiveStrokeColor='#ddd'
                backgroundColor='#fff'
                textColor='#ddd'
                labelFontSize={largeScreen?12:bigScreen?10:6}
                valueFontSize={largeScreen?60:bigScreen?40:20}
                withGradient={false}
                anticlockwise={false}
                initialAngularDisplacement={0}/>
        </div>
        <ReactModal
            style={largeScreen?largeModalStyle:(bigScreen)?bigModalStyle:smallModalStyle}
            isOpen={failureVisible}>
            <div className="flex w-full h-full items-center flex-col">
                <FiCloudLightning
                    size={30}
                    className="text-red-600"/>
                <p>
                    {"Failed to log you in. It's not you,"}
                </p>
                <p>(maybe it is),</p>
                <p>
                    but if its not you its us.</p>
                <p>
                    Return <a className="text-blue-600" href="/">Home</a>!
                </p>
            </div>
        </ReactModal>
        </main>
    )
}
