import { useState } from "react"
import { TweetType } from "../apis/tweets"
import Editor  from "./editor"
import Tweet from "./tweet"
import { GiBoltBomb, GiChemicalDrop } from "react-icons/gi"
type EditPairProps = {
    editting: boolean,
    viewing: boolean,
    tweet: TweetType|null,
    isLoggedIn: boolean,
    defaultMessage: string,
    onChange: Function,
    onSendClicked: Function,
    showLoading: boolean,
    value: string,
    showMenu: boolean,
    editClicked: Function,
    deleteClicked: Function,
    hideClicked: Function,
    editorHideable: boolean,
    url: string
}

export default function EditPair( props: EditPairProps) {

    return (
        <div className="w-full flex flex-col items-center">
        {(props.editting)?(
            <Editor
            isLoggedIn={props.isLoggedIn}
            defaultMessage={props.defaultMessage}
            showLoading={props.showLoading}
            value={props.value}
            onSendClicked={props.onSendClicked}
            onChange={props.onChange}
            hideable={props.editorHideable}
            url={props.url}
            hideClicked={props.hideClicked}/>
        ):null}
        {(props.viewing?(
            <div
            className="flex flex-col items-center w-full">
            <Tweet
            url={props.url}
            tweet_id={props.tweet?.id?props.tweet.id:""}
            tweet={props.value}
            date={props.tweet?.created_at?props.tweet.created_at:""}
            deleteClicked={()=>{props.deleteClicked()}}
            editClicked={()=>{props.editClicked()}}
            showMenu={props.showMenu}
            onClick={()=>{}}/>
            </div>
        ):null)}
        </div>
    )
}
