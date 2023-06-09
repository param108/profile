import { formatTweet } from "../strings";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context";
import moment from "moment";
type TweetProps = {
    tweet_id: string,
    tweet: string,
    key: Number,
    date: string,
    router: null|AppRouterInstance
}

export default function Tweet(props: TweetProps) {
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
                onClick={()=>{
                    // FIXME: Turning this off until it is implemented
                    // props.router?.push('/tweets/'+props.tweet_id+"/show");
                }}>
                <i className="text-gray-300">{formatDate(props.date)}</i><br/>
                <span className="text-gray-600">{formatTweet(props.tweet)}</span>
            </div>
    );
}
