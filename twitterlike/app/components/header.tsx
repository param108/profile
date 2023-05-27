"use client";

import { MouseEventHandler, ReactElement, useEffect, useState } from "react";
import { FiCloudRain, FiZap } from "react-icons/fi";
import ReactModal from "react-modal";
import { JsxElement } from "typescript";

const bigModalStyle = {
    content: {
        left: "25%",
        right: "25%",
        top: "100px"
    }
};

const smallModalStyle = {
    content: {
        left: "10%",
        right: "10%",
        top: "100px"
    }
};

export default function Header() {
    var [showVMenu, setShowVMenu] = useState(false)

    // Which modal is open
    var [openModal, setOpenModal] = useState("")

    const [largeScreen, setLargeScreen] = useState(
        window.matchMedia("(min-width: 768px)").matches
    )

    useEffect(() => {
        window
            .matchMedia("(min-width: 768px)")
            .addEventListener('change', e => setLargeScreen( e.matches ));
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
            <div className="relative top-[50%] w-full flex flex-col items-center -translate-y-1/2">
                <span>If you are already signed up, this will log you in. </span>
                <span>If not, this will <b>sign you up</b></span>
                {twitterLogin()}
                <span><i>Your username will be your twitter handle.</i></span>
            </div>
        )
    };

    const signupDiv = function(): ReactElement {
        return (
            <div className="h-full w-full flex flex-col items-center">
                <span>When you signup, your twitter handle becomes your username...</span>
                {twitterLogin()}
            </div>
        )
    };

    const aboutDiv = function(): ReactElement {
        return (
            <div></div>
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
            {(() => {
                var modalStyle = largeScreen?bigModalStyle:smallModalStyle;
                return (
                    <div>
                    <ReactModal
                        style={modalStyle}
                        isOpen={openModal=="login"}>
                        <FiZap
                            size={30}
                            className="text-pink-600 float-right"
                            onClick={()=>{setOpenModal("closed")}}/>
                        {loginDiv()}
                    </ReactModal>
                    <ReactModal
                        style={modalStyle}
                        isOpen={openModal=="about"}>
                        <FiZap
                            size={30}
                            className="text-pink-600 float-right"
                            onClick={()=>{setOpenModal("closed")}}/>
                        {aboutDiv()}
                    </ReactModal>
                    </div>
                )
            })()}
        </div>
    );
}
