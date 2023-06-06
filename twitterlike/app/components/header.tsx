"use client";

import { AppRouterInstance } from "next/dist/shared/lib/app-router-context";
import { usePathname } from "next/navigation";
import { MouseEventHandler, ReactElement, useEffect, useState } from "react";
import { FiCloudRain, FiZap } from "react-icons/fi";
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

export default function Header() {
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
            <div className="relative top-[50%] w-full flex flex-col items-center -translate-y-1/2 text-gray-600">
                <span>If you are already signed up, this will log you in. </span>
                <span>If not, this will <b>sign you up</b></span>
                <div onClick={()=>{
                    location.href ="https://data.tribist.com/users/login?source=twitter&redirect_url="+
                        path;
                }}>
                {twitterLogin()}
                </div>
                <span><i>Your username will be your twitter handle.</i></span>
            </div>
        )
    };

    const aboutDiv = function(): ReactElement {
        return (
            <div className="pt-[50px] px-[5px] md:px-[50px] text-gray-600">
                <p>I am a twitter addict. I tweet throughout the day even on bad days.
                My tweets cover everything from my random thoughts, to stuff I have read, to politics. Some of these are
                only output and seldom re-read while others I like to revisit over and over again.</p><br/>
                <p>I wanted a place to organize and re-organize my tweets to extract different perspectives and share these new perspectives
                with others. Unfortunately twitter is not great for this. So I decided to develop my own microblog.</p><br/>
                <p>This blog is optimized for viewing, organizing, studying and finally sharing tweets.
                In time, I hope you will be able to sell your organized tweets to your readers. I am not a fan of ad-revenue and want this
                place to be where you find your tribe of real people who you will inspire and gather inspiration from.</p><br/>
                <p>You can follow me on twitter <a className="text-indigo-600" href="https://twitter.com/param108">{"@param108"}</a><br/>
                My microblog on tribist is <a className="text-indigo-600" href="https://ui.tribist.com/param108">{"@param108"}</a></p><br/>
                <p>This <b>microblog</b> is work in progress and is <b>Open Source.</b> If you would like to contribute,
                send me a pull request or create an issue at <a href="https://github.com/param108/profile">github.com/param108/profile</a>.</p>
            </div>
        )
    };

    return (
        <div className="fixed bg-black h-[50px] md:p-[5px] mb-[5px] w-full md:items-center">
            <div className="hidden md:block w-full">
                <button className="text-white float-right p-[5px] mr-[50px]"
                    onClick={menuClick("login")}>{"Login/Signup"}</button>
                <button className="text-white float-right p-[5px] mr-[50px]"
                    onClick={menuClick("about")}>{"About"}</button>
            </div>
            <div className="flex flex-col overflow-y-visible md:hidden ">
            <div className="h-[50px] flex items-center">
                <div className="mx-[15px] text-white">
            <FiCloudRain size={30} onClick={()=>{setShowVMenu(!showVMenu)}}/>
                </div>
            </div>
            {((show: Boolean)=> {
                if (show) {
                    return (
                <div className="bg-black text-white max-w-[70%]">
                    <ul className="">
                            <li className="bg-black hover:bg-slate-500 w-full pl-[5px] py-[5px] block"
                                onClick={menuClick("login")}>{"Login/Signup"}</li>
                            <li className="bg-black hover:bg-slate-500 w-full pl-[5px] py-[5px] block"
                                onClick={menuClick("about")}>{"About"}</li>
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
