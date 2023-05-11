import { NextRouter, useRouter } from "next/router";
import { formatTweet } from "../strings";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context";

type TweetProps = {
    tweet_id: string,
    tweet: string,
    key: Number,
    date: string,
    router: AppRouterInstance
}

export default function Tweet(props: TweetProps) {
    function formatDate(date:String):String {
        return date;
    }
    return (
            <div className="border border-t-slate-100 
            bg-white hover:bg-cyan-50 w-[510px] 
            min-h-[100px] px-[5px] pt-[5px] 
            pb-[40px] overflow-x-auto"
                onClick={()=>{
                    props.router.push('/tweets/'+props.tweet_id+"/show");
                }}>
                <b>{formatDate(props.date)}</b><br/>
                {formatTweet(props.tweet)}
            </div>
    );
}