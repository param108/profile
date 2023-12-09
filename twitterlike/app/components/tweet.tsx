import { formatTweet, hasThread, ThreadInfo } from "../strings";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context";
import moment from "moment";
import { useEffect, useState } from "react";
import { GiBoltBomb, GiChemicalDrop, GiScrollUnfurled, GiSewingNeedle, GiTiedScroll } from "react-icons/gi"
import { getThread, ThreadData } from "../apis/threads";
import { object } from "underscore";
import { FiExternalLink } from "react-icons/fi";
type TweetProps = {
    tweet_id: string,
    tweet: string,
    flags: string,
    date: string,
    onClick: Function,
    deleteClicked: Function,
    editClicked: Function,
    externalClicked: Function|null,
    createThreadClicked?: Function,
    showMenu: boolean,
    url: string
    visible: boolean,
    threadList: (ThreadData|null)[],
    viewThread: Function|null,
    shownThread?: string,
    imageSource?: string|null,
}

export default function Tweet(props: TweetProps) {
    var [ menuVisible, setMenuVisible ] = useState(false)
    var [ expanded, setExpanded ] = useState(false);

    function formatDate(date:string):String {
        if ((new Date(date)).getTime() > 0) {
            // valid timestamp
            return moment(date).format('llll')
        }

        return date;
    }

    var toplayerStyle = `
            w-full bg-white dark:bg-slate-700 hover:bg-cyan-50
            min-h-[100px] pl-[15px] pr-[5px] pt-[5px]
            pb-[40px] overflow-x-auto m-auto`;
    if (!props.visible) {
        toplayerStyle += " invisible"
    }

    // if atleast one thread in the threadlist
    // doesnt match the shownthread
    // return true
    const atleastOneThread= ()=>{
        if (props.threadList && props.threadList.length) {
            return props.threadList.some((v)=>{
                if (props.shownThread !== undefined) {
                    return v?.id != props.shownThread
                }
                return true
            })
        }
        return false
    }

    return (
        <div className="border w-[90%] md:w-[510px]">
            <div className={toplayerStyle}
            onMouseLeave={()=>setMenuVisible(false)}
            onMouseEnter={()=>setMenuVisible(true)}

            onClick={()=>{
                props.onClick(props.tweet_id)
                // FIXME: Turning this off until it is implemented
                // props.router?.push('/tweets/'+props.tweet_id+"/show");
            }}>
            {(menuVisible && props.showMenu)?(
            <span className="bg-sky-200 mt-[5px] w-[90%] md:w-[510px] rounded-t">
                <GiBoltBomb onClick={()=>{
                        props.deleteClicked()}}
                        size={35} className="cursor-pointer bg-sky-200 rounded my-[5px] text-red-600 p-[5px] mx-[5px] float-right"/>
                <GiChemicalDrop onClick={()=>{
                        props.editClicked()}}
                        size={35} className="cursor-pointer bg-sky-200 rounded my-[5px] text-indigo-800 p-[5px] mx-[5px] float-right"/>
                <GiSewingNeedle onClick={()=>{
                    if (props.createThreadClicked) {
                        props.createThreadClicked()
                    }}}
                        size={35} className="cursor-pointer bg-sky-200 rounded my-[5px] text-indigo-800 p-[5px] mx-[5px] float-right"/>
            </span>
            ):null}

            <div className="flex flex-row">
            <i className="text-gray-500">{formatDate(props.date)}</i>
            {(menuVisible && props.showMenu)?
                (<FiExternalLink className="ml-[10px] cursor-pointer text-sky-800 dark:text-sky-500"size={20} onClick={()=>{
                    if (props.externalClicked) {
                        props.externalClicked(props.tweet_id);
                    }}}/>):null}<br/>
            </div>
            <span className="text-gray-600 dark:text-slate-100">{formatTweet(props.tweet, props.url, props.flags)}</span>
            </div>
            {atleastOneThread()?(
            <div className={(menuVisible && props.showMenu)?"w-full overflow-auto bg-cyan-50 dark:bg-slate-700":
                "w-full overflow-auto bg-gray-50 dark:bg-slate-700"}>
             {(!expanded)?(
                 <span className="float-right" onClick={()=>setExpanded(true)}><GiScrollUnfurled className="cursor-pointer m-[5px] p-[5px] rounded text-indigo-800 bg-sky-200" size={30}/></span>
            ):(
                <ul>
                     <li className="overflow-auto" onClick={()=>{setExpanded(false)}}><GiTiedScroll className="cursor-pointer float-right p-[5px] rounded text-indigo-800 bg-sky-200 m-[5px]" size={30}/></li>
                    {
                        props.threadList.map((v) => {
                            if (v && (props.shownThread != v.id)) {
                                return (
                                    <li
                                    className="cursor-pointer select-none text-blue-700 dark:text-cyan-300 px-[15px] mb-[2px]"
                                    key={v.id}
                                    onClick={()=>{ if(props.viewThread){
                                        props.viewThread(v.id);
                                    }}}><GiSewingNeedle className="cursor-pointer text-gray-500 inline mx-[5px]"/><i>{v.name}</i></li>
                                );
                            }
                        })
                    }
                </ul>
             )}
            </div>
            ):null}
            {/* image */}
            {props.imageSource?(
            <img src={props.imageSource} className=" w-full md:w-[510px]"/>
            ):null
            }
        </div>
    );
}
