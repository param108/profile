import { useEffect, useState } from "react"
import { FiImage, FiScissors, FiSend, FiUserX, FiX } from "react-icons/fi"
import { formatTweet } from "../strings"
import RingLoader from "react-spinners/RingLoader"
type EditorProps = {
    isLoggedIn: Boolean,
    defaultMessage: string,
    onSendClicked: Function,
    onChange: Function,
    onFlagsChange: Function,
    showLoading: boolean,
    hideClicked: Function,
    hideable: boolean,
    value: string,
    flags: string,
    url: string
    headerMargin?: boolean,
}

export const MaxTweetLength:number = 500

export default function Editor(props: EditorProps) {

    var topDivStyle = "bg-sky-200 w-[96%] md:w-[510px] rounded mb-[10px]";
    function onSendClick() {
        props.onSendClicked(props.value, props.flags)
    }

    return (
        <div className={topDivStyle}>
            {props.isLoggedIn ? (
                <div>
                <textarea value={props.value} onChange={(e)=> (props.onChange(e.target.value))} placeholder={"What are you thinking about ?"}
                    className="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[150px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100">
                </textarea>
                <textarea value={props.flags} onChange={(e)=> (props.onFlagsChange(e.target.value))} placeholder={"display tags..."}
                    className="block dark:bg-slate-700 w-[96%] md:w-[500px] h-[50px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none text-black dark:text-slate-100">
                </textarea>
                </div>
            ):(
                <div className="w-[96%] md:w-[500px] h-[150px] bg-white
mt-[5px] mx-[2%] md:mx-[5px] p-[5px] rounded focus:outline-none overflow-x-auto text-gray-600">
                {formatTweet(props.defaultMessage, props.url, props.flags)}
                </div>    
            )}
            {props.isLoggedIn ? (
                        <div className="block pt-[5px]">
                            <button className="px-[10px]">
                                <FiImage className="text-indigo-800" size={30} />
                            </button>
                            <div className="inline-block float-right pr-[10px]">
                                <RingLoader className="inline-block" color="#EC4899"
                                    loading={props.showLoading} size={30}/>
                            </div>
                            {(props.hideable)?(
                                <div onClick={()=>{props.hideClicked();}} className="inline-block float-left pl-[10px]">
                                <FiX className="text-red-800" size={30}/>
                                    </div>):null}
                            {(props.value.length <= MaxTweetLength)?(<div className="inline-block float-right">
                                <button
                                className="px-[10px] float-right"
                                onClick={()=>onSendClick()}>
                                    <FiSend className="text-indigo-800" size={30} />
                                </button>
                                <span className="text-indigo-800 text-lg float-right items-center px-[10px]">{MaxTweetLength - props.value.length}</span>
                                </div>):(<div className="inline-block float-right">
                                <button
                                className="px-[10px] float-right"
                                onClick={()=>{}}>
                                    <FiSend className="text-red-800" size={30} />
                                </button>
                                <span className="text-red-800 text-lg float-right items-center px-[10px]">{MaxTweetLength - props.value.length}</span>
                                </div>)}
                        </div>
                    ):(
                        <div className="block pt-[5px] min-h-[35px]"></div>
                    )}

        </div>
    )
}
