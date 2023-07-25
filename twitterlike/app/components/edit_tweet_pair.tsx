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
    url: string,
    visible: boolean,
    headerMargin?: boolean
}

export default function EditPair( props: EditPairProps) {
    var toplayerStyle = "w-full flex flex-col items-center";
    var headerMargin = true;
    if (props.headerMargin !== undefined) {
        headerMargin = props.headerMargin;
    }

    if (!props.visible) {
        toplayerStyle += " invisible"
    }
    return (
        <div className={toplayerStyle}>
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
            hideClicked={props.hideClicked}
            headerMargin={headerMargin}/>
        ):null}
        {(props.viewing?(
            <div
            className="flex flex-col items-center w-full">
            <Tweet
            url={props.url}
            visible={props.visible}
            tweet_id={props.tweet?.id?props.tweet.id:""}
            tweet={props.value}
            date={props.tweet?.created_at?props.tweet.created_at:""}
            deleteClicked={()=>{props.deleteClicked()}}
            editClicked={()=>{props.editClicked()}}
            externalClicked={null}
            showMenu={props.showMenu}
            onClick={()=>{}}/>
            </div>
        ):null)}
        </div>
    )
}
