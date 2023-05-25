import { FiImage, FiSend } from "react-icons/fi"
import { formatTweet } from "../strings"

type EditorProps = {
    isLoggedIn: Boolean,
    defaultMessage: string
}

export default function Editor(props: EditorProps) {
    return (
        <div className="bg-sky-200 rounded mt-[60px] mb-[10px]">
            {props.isLoggedIn ? (
                <textarea className="block w-[90%] md:w-[500px] h-[150px] resize-none caret-red-500 mt-[5px] mx-[5px] p-[5px] rounded focus:outline-none">
                    What are you thinking about ?
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
                            <button className="px-[10px] float-right"><FiSend className="text-indigo-800" size={30} /></button>
                        </div>
                    ):(
                        <div className="block pt-[5px] min-h-[35px]"></div>
                    )}

        </div>
    )
}
