import { useEffect, useState } from "react"
import { FiImage, FiSend } from "react-icons/fi"
import { formatTweet } from "../strings"
import RingLoader from "react-spinners/RingLoader"
type EditorProps = {
    isLoggedIn: Boolean,
    defaultMessage: string,
    onSendClicked: Function,
    onChange: Function,
    showLoading: boolean,
    value: string
}

export default function Editor(props: EditorProps) {

    function onSendClick() {
        props.onSendClicked(props.value)
    }

    return (
        <div className="bg-sky-200 w-[96%] md:w-[510px] rounded mt-[60px] mb-[10px]">
            {props.isLoggedIn ? (
                <textarea value={props.value} onChange={(e)=> (props.onChange(e.target.value))} placeholder={"What are you thinking about ?"}
                    className="block w-[96%] md:w-[500px] h-[150px] resize-none caret-red-500 mt-[5px] mx-[2%] md:mx-[5px] pl-[10px] pr-[5px] py-[5px] rounded focus:outline-none">
                </textarea>
            ):(
                <div className="w-[96%] md:w-[500px] h-[150px] bg-white
mt-[5px] mx-[2%] md:mx-[5px] p-[5px] rounded focus:outline-none overflow-x-auto text-gray-600">
                {formatTweet(props.defaultMessage)}
                </div>    
            )}
            {props.isLoggedIn ? (
                        <div className="block pt-[5px]">
                            <button className="px-[10px]"><FiImage className="text-indigo-800" size={30} /></button>
                            <div className="inline-block float-right pr-[10px]">
                                <RingLoader className="inline-block" color="#EC4899"
                                    loading={props.showLoading} size={30}/>
                            </div>
                    <button className="px-[10px] float-right" onClick={()=>onSendClick()}><FiSend className="text-indigo-800" size={30} /></button>
                        </div>
                    ):(
                        <div className="block pt-[5px] min-h-[35px]"></div>
                    )}

        </div>
    )
}
