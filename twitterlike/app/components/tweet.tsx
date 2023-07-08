import { formatTweet, hasThread, ThreadData, ThreadInfo } from "../strings";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context";
import moment from "moment";
import { useEffect, useState } from "react";
import { GiBoltBomb, GiChemicalDrop, GiScrollUnfurled, GiSewingNeedle, GiTiedScroll } from "react-icons/gi"
import { getThread } from "../apis/threads";
import { object } from "underscore";
type TweetProps = {
    tweet_id: string,
    tweet: string,
    date: string,
    onClick: Function,
    deleteClicked: Function,
    editClicked: Function,
    showMenu: boolean,
    url: string
    visible: boolean,
    token: string
}

export default function Tweet(props: TweetProps) {
    var [ menuVisible, setMenuVisible ] = useState(false)
    var [ threadList, setThreadList ] = useState<ThreadData[]>([])
    var [ threadData, setThreadData ] = useState<{[name:string]:ThreadData}>({})
    var [ visibleThreadList, setVisibleThreadList ] = useState<ThreadData[]>([])
    var [ expanded, setExpanded ] = useState(false);

    useEffect(()=>{
        setThreadList(hasThread(props.tweet).sort((x,y)=>{
            if (x.seq >= y.seq) {
                return 1;
            }
            return -1
        }));
    },[]);

    useEffect(()=>{
        if (threadList.length == 0) {
            return
        }

        let newThreadData:{[name:string]:ThreadData} = {};
        threadList.forEach((thread : ThreadInfo)=> {

            if (thread.id in threadData) {
                return;
            }

            let key = thread.id;
            newThreadData[key] = {
                id: thread.id,
            }

            getThread(props.token, thread.id).
                then((res)=>{
                    let key = res.data.data.id;
                    let data = {...threadData};
                    data[key] = res.data.data;
                    setThreadData(data);
                })
        })

        setThreadData({
            ...threadData,
            ...newThreadData
        })

        setVisibleThreadList([...threadList])

    },[threadList])

    useEffect(()=>{
        let newThreadList: ThreadInfo[] =[];
        visibleThreadList.forEach((x) => {
            if ( x.id in threadData ) {
                if ("name" in threadData[x.id]) {
                    x.name = threadData[x.id]['name']
                }
                newThreadList.push(x)
            }
        })

        setVisibleThreadList([...threadList])
    }, [threadData])

    function formatDate(date:string):String {
        if ((new Date(date)).getTime() > 0) {
            // valid timestamp
            return moment(date).format('llll')
        }

        return date;
    }

    var toplayerStyle = `
            w-full bg-white hover:bg-cyan-50
            min-h-[100px] pl-[15px] pr-[5px] pt-[5px]
            pb-[40px] overflow-x-auto`;
    if (!props.visible) {
        toplayerStyle += " invisible"
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
                        size={35} className="bg-sky-200 rounded my-[5px] text-red-600 p-[5px] mx-[5px] float-right"/>
                <GiChemicalDrop onClick={()=>{
                        props.editClicked()}}
                        size={35} className="bg-sky-200 rounded my-[5px] text-indigo-800 p-[5px] mx-[5px] float-right"/>
                <GiSewingNeedle onClick={()=>{
                        props.editClicked()}}
                        size={35} className="bg-sky-200 rounded my-[5px] text-indigo-800 p-[5px] mx-[5px] float-right"/>
            </span>
            ):null}

            <i className="text-gray-300">{formatDate(props.date)}</i><br/>
            <span className="text-gray-600">{formatTweet(props.tweet, props.url)}</span>
            </div>
            {visibleThreadList.length > 0?(
            <div className="w-full  overflow-auto">
             {(!expanded)?(
                 <span className="float-right" onClick={()=>setExpanded(true)}><GiScrollUnfurled className="m-[5px] p-[2px] rounded text-gray-500 bg-gray-200" size={30}/></span>
            ):(
                <ul>
                     <li className="overflow-auto" onClick={()=>{setExpanded(false)}}><GiTiedScroll className="float-right p-[2px] rounded text-gray-500 bg-gray-200 m-[5px]" size={30}/></li>
                    {
                        visibleThreadList.map((v) => {
                            return (
                                <li className="text-blue-700 px-[15px] mb-[2px]"><GiSewingNeedle className="text-gray-500 inline mx-[5px]"/><i>{v.name}</i></li>
                            );
                        })
                    }
                </ul>
             )}

            </div>
            ):null}
        </div>
    );
}
