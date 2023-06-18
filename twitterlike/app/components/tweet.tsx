import { formatTweet } from "../strings";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context";
import moment from "moment";
import { useState } from "react";
import { GiBoltBomb, GiChemicalDrop } from "react-icons/gi"
type TweetProps = {
    tweet_id: string,
    tweet: string,
    date: string,
    onClick: Function,
    deleteClicked: Function,
    editClicked: Function,
    showMenu: boolean
}

export default function Tweet(props: TweetProps) {
    var [ menuVisible, setMenuVisible ] = useState(false)
    function formatDate(date:string):String {
        if ((new Date(date)).getTime() > 0) {
            // valid timestamp
            return moment(date).format('llll')
        }

        return date;
    }
    return (
            <div className="border border-t-slate-100 
            bg-white hover:bg-cyan-50 w-[90%] md:w-[510px]
            min-h-[100px] pl-[15px] pr-[5px] pt-[5px]
            pb-[40px] overflow-x-auto"
            onMouseLeave={()=>setMenuVisible(false)}
            onMouseEnter={()=>setMenuVisible(true)}

            onClick={()=>{
                props.onClick(props.tweet_id)
                // FIXME: Turning this off until it is implemented
                // props.router?.push('/tweets/'+props.tweet_id+"/show");
            }}>
            {(menuVisible && props.showMenu)?(
            <span className="border-x-1 bg-sky-200 mt-[5px] w-[90%] md:w-[510px] rounded-t">
                <GiBoltBomb onClick={()=>{
                        props.deleteClicked()}}
                        size={35} className="bg-sky-200 rounded my-[5px] text-red-600 p-[5px] mx-[5px] float-right"/>
                <GiChemicalDrop onClick={()=>{
                        props.editClicked()}}
                        size={35} className="bg-sky-200 rounded my-[5px] text-indigo-800 p-[5px] mx-[5px] float-right"/>
            </span>
            ):null}

            <i className="text-gray-300">{formatDate(props.date)}</i><br/>
            <span className="text-gray-600">{formatTweet(props.tweet)}</span>
            </div>
    );
}
