"use client";

import { AppRouterInstance } from "next/dist/shared/lib/app-router-context";
import { usePathname, useRouter } from "next/navigation";
import { MouseEventHandler, ReactElement, useEffect, useReducer, useState } from "react";
import { FiCloudRain, FiZap } from "react-icons/fi";
import ReactModal from "react-modal";
import RingLoader from "react-spinners/RingLoader"
import Toggle from 'react-toggle'
import 'react-toggle/style.css'
import './header.css'
const largeModalStyle = {
    content: {
        left: "25%",
        right: "25%",
        top: "100px",
        background: "#64748b"
    }
};

const bigModalStyle = {
    content: {
        left: "10%",
        right: "10%",
        top: "100px",
        background: "#64748b"
    }
};

const smallModalStyle = {
    content: {
        left: "2%",
        right: "2%",
        top: "100px",
        background: "#64748b"
    }
};

type HeaderProps = {
    showSpinner: boolean,
    changeDarkMode: Function|null
}

export default function Header(props:HeaderProps) {
    var [showVMenu, setShowVMenu] = useState(false)

    const path = usePathname();

    // Which modal is open
    var [openModal, setOpenModal] = useState("")

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

    const menuClick = function (itemName:string):MouseEventHandler {
        return function() {
            setOpenModal(itemName);
            setShowVMenu(false);
        }
    }

    const twitterLogin = function() : ReactElement {
        return (
            <picture className="relative">
                <source srcSet="https://cdn.cms-twdigitalassets.com/content/dam/developer-twitter/auth-docs/sign-in-with-twitter-gray.png.twimg.768.png" media="(max-width: 767px)"/>
                <source srcSet="https://cdn.cms-twdigitalassets.com/content/dam/developer-twitter/auth-docs/sign-in-with-twitter-gray.png.twimg.2560.png" media="(min-width: 1920px)"/>

                <img src="https://cdn.cms-twdigitalassets.com/content/dam/developer-twitter/auth-docs/sign-in-with-twitter-gray.png.twimg.1920.png"
                    style={{
                        width: "158px"
                    }} width="158px" height="28px"
                    data-src="https://cdn.cms-twdigitalassets.com/content/dam/developer-twitter/auth-docs/sign-in-with-twitter-gray.png.twimg.1920.png"
                    alt="" data-object-fit="cover"/>
            </picture>);
    };

    const loginDiv = function(): ReactElement {
        return (
            <div className="relative top-[50%] w-full flex flex-col items-center -translate-y-1/2 text-gray-50">
                <span>If you are already signed up, this will log you in. </span>
                <span>If not, this will <b>sign you up</b></span>
                <div onClick={()=>{
                    location.href =`${process.env.NEXT_PUBLIC_BE_URL}/users/login?source=twitter&redirect_url=${path}`;
                }}>
                {twitterLogin()}
                </div>
                <span><i>Your username will be your twitter handle.</i></span>
                <span>Or login as <b>guest</b> to play around.</span>
                <span>
                    <a
                        href={`${process.env.NEXT_PUBLIC_BE_URL}/users/login?source=guest&redirect_url=${path}&guest=true`}>
                    Login as Guest</a></span>
            </div>
        )
    };

    var [loggedInUser, setLoggedInUser] = useState("")
    var [useDarkMode, setUseDarkMode] = useState(true)
    useEffect(()=>{
        const username = localStorage.getItem('username')
        if (username && username.length > 0) {
            setLoggedInUser(username)
        }
        const darkMode = localStorage.getItem('dark_mode')
        if (darkMode) {
            if (darkMode == "dark") {
                setUseDarkMode(true)
            } else {
                setUseDarkMode(false)
            }
        } else {
            setUseDarkMode(false)
        }
        if (props.changeDarkMode) {
            props.changeDarkMode(darkMode)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    const aboutDiv = function(): ReactElement {
        return (
            <div className="pt-[50px] px-[5px] md:px-[50px] text-gray-50">
                <p>I am a twitter addict. I tweet throughout the day even on bad days.
                My tweets cover everything from my random thoughts, to stuff I have read, to politics. Some of these are
                only output and seldom re-read while others I like to revisit over and over again.</p><br/>
                <p>I wanted a place to organize and re-organize my tweets to extract different perspectives and share these new perspectives
                with others. Unfortunately twitter is not great for this. So I decided to develop my own microblog.</p><br/>
                <p>This blog is optimized for viewing, organizing, studying and finally sharing tweets.
                In time, I hope you will be able to sell your organized tweets to your readers. I am not a fan of ad-revenue and want this
                place to be where you find your tribe of real people who you will inspire and gather inspiration from.</p><br/>
                <p>You can follow me on twitter <a className="text-indigo-600" href="https://twitter.com/param108">{"@param108"}</a><br/>
                My microblog on tribist is <a className="text-indigo-600" href="https://ui.tribist.com/user/param108/tweets">{"@param108"}</a></p><br/>
                <p>This <b>microblog</b> is work in progress and is <b>Open Source.</b> If you would like to contribute,
                send me a pull request or create an issue at <a href="https://github.com/param108/profile">github.com/param108/profile</a>.</p>
            </div>
        )
    };

    return (
        <div className="fixed bg-black h-[50px] md:p-[5px] mb-[5px] w-full md:items-center">
            <div className="hidden md:block w-full">
            {(loggedInUser && loggedInUser.length > 0)?(
                <button className="text-pink-600 float-left p-[5px] mr-[5px]"
                onClick={()=>(location.href=`/user/${loggedInUser}/tweets`)}>{"@"+loggedInUser}</button>): null}
                <RingLoader className="inline-block float-left" color="#EC4899"
                        loading={props.showSpinner} size={30}/>
                <div className="float-right mr-[50px] p-[5px]">
                <Toggle
                checked={useDarkMode}
                onChange={()=>{
                    setUseDarkMode(!useDarkMode);
                    let mode = (!useDarkMode)?"dark":"light";
                    if (props.changeDarkMode) {
                        props.changeDarkMode(mode)
                    }
                    localStorage.setItem("dark_mode", mode)
                }} icons={false}/>
                </div>
                <button className="text-white float-right p-[5px] mr-[50px]"
                    onClick={menuClick("login")}>{"Login/Signup"}</button>
                <button className="text-white float-right p-[5px] mr-[50px]"
                    onClick={menuClick("about")}>{"About"}</button>
            {((!loggedInUser) || loggedInUser.length === 0)?(
            <button className="text-white float-right p-[5px] mr-[50px]"
                    onClick={()=>{location.href=`${process.env.NEXT_PUBLIC_BE_URL}/users/login?source=guest&redirect_url=${path}&guest=true`}}>{"Guest Login"}</button>
            ):null}
            </div>
            <div className="flex flex-col overflow-y-visible md:hidden ">
            <div className="h-[50px] flex items-center">
                <div className="mx-[15px] flex flex-row text-white">
            <FiCloudRain size={30} onClick={()=>{setShowVMenu(!showVMenu)}}/>
            <RingLoader className="ml-[5px]" color="#EC4899"
                        loading={props.showSpinner} size={30}/>
                </div>
            </div>
            {((show: Boolean)=> {
                if (show) {
                    return (
                <div className="bg-black text-white max-w-[70%]">
                    <ul className="">
                            {(loggedInUser && loggedInUser.length > 0)?(
                                <li className="select-none bg-black hover:bg-slate-500 w-full pl-[15px] py-[5px] block"
                                    onClick={() =>(location.href=`/user/${loggedInUser}/tweets`)}>{"@"+loggedInUser}</li>): null}
                            <li className="select-none bg-black hover:bg-slate-500 w-full pl-[15px] py-[5px] block"
                                onClick={menuClick("login")}>{"Login/Signup"}</li>
                            <li className="select-none bg-black hover:bg-slate-500 w-full pl-[15px] py-[5px] block"
                                onClick={menuClick("about")}>{"About"}</li>
                            {((!loggedInUser) || loggedInUser.length === 0)?(
                            <li className="select-none bg-black hover:bg-slate-500 w-full pl-[15px] py-[5px] block"
                                onClick={()=>{location.href=`${process.env.NEXT_PUBLIC_BE_URL}/users/login?source=guest&redirect_url=${path}&guest=true`}}>{"Guest Login"}</li>
                            ):null}
                            <li className="select-none bg-black hover:bg-slate-500 w-full pl-[15px] py-[5px] block">
                                <Toggle
                                    checked={useDarkMode}
                                        onChange={() => {
                                            setUseDarkMode(!useDarkMode);
                                            let mode = (!useDarkMode) ? "dark" : "light";
                                            if (props.changeDarkMode) {
                                                props.changeDarkMode(mode)
                                            }
                                            localStorage.setItem("dark_mode", mode)
                                        }} icons={false} />
                            </li>
                    </ul>
                </div>
                    )
                }
            })(showVMenu)}
            </div>
            <div>
                <ReactModal
                        style={largeScreen?largeModalStyle:(bigScreen)?bigModalStyle:smallModalStyle}
                        isOpen={openModal=="login"}>
                        <FiZap
                            size={30}
                            className="text-pink-600 float-right"
                            onClick={()=>{setOpenModal("closed")}}/>
                        {loginDiv()}
                </ReactModal>
                <ReactModal
                        style={largeScreen?largeModalStyle:(bigScreen)?bigModalStyle:smallModalStyle}
                        isOpen={openModal=="about"}>
                    <FiZap
                            size={30}
                            className="text-pink-600 float-right"
                            onClick={()=>{setOpenModal("closed")}}/>
                        {aboutDiv()}
                </ReactModal>
            </div>
        </div>
    );
}
