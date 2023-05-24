"use client";
import { ReactElement } from "react";
const showdown = require('showdown');
import './strings.css';

function isFlagOn(tweet: string, flag: string): Boolean {
    return tweet.includes(flag);
}

export function formatTweet(tweet: string):ReactElement {
    const converter = new showdown.Converter();
    
    // The first line is the command line if it exists
    // Default is for it not to exist.
    let commandLineExists = false;
    let kamalFont = false;

    if (isFlagOn(tweet, "#font:kamal")) {
        commandLineExists = true;
        kamalFont = true;
    }

    // If the command Line Exists we need to remove it after processing.
    if (commandLineExists) {
        let newTweet = tweet.split('\n');
        newTweet.splice(0,1);
        tweet=newTweet.join('\n');
    }

    let hdata = converter.makeHtml(tweet);
    let classNames = "";
    if (kamalFont) {
        classNames = classNames + " font-kamal"
    }

    return (
        <div>
            <div className={classNames} dangerouslySetInnerHTML={{__html: hdata}}>
            </div>
        </div>
    )
}
