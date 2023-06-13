import { useState } from "react"
import { TweetType } from "../apis/tweets"
import Editor  from "./editor"
import Tweet from "./tweet"
type EditPairProps = {
    editting: boolean,
    viewing: boolean,
    tweet: TweetType|null,
    isLoggedIn: boolean,
    defaultMessage: string,
    onChange: Function,
    onSendClicked: Function,
    showLoading: boolean,
    value: string
    key: number
}

export default function EditPair( props: EditPairProps) {
    return (
        <div className="flex flex-col w-full items-center">
        {(props.editting)?(
            <Editor
            isLoggedIn={props.isLoggedIn}
            defaultMessage={props.defaultMessage}
            showLoading={props.showLoading}
            value={props.value}
            onSendClicked={props.onSendClicked}
            onChange={props.onChange}/>
        ):null}
        {(props.viewing?(
            <Tweet
            tweet_id={props.tweet?.id?props.tweet.id:""}
            tweet={props.value}
            key={props.key}
            date={props.tweet?.created_at?props.tweet.created_at:""}
            onClick={()=>{}}
            router={null}/>
        ):null)}
        </div>
    )
}
