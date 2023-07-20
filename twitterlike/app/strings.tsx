"use client";
import { ReactElement } from "react";
const showdown = require('showdown');
import './strings.css';


function isFlagOn(tweet: string, flag: string): Boolean {
    return tweet.includes(flag);
}

// searches for hyperlinks in second lines and converts them
// for display as links
export function tagsToHyperlinks(tweet: string, baseURL: string, commandLineExists: boolean): string {
    // Ignore tags on the first line
    let newTweet = tweet.split('\n');
    let firstline = ""
    if (commandLineExists) {
        firstline = newTweet[0]
        newTweet.splice(0,1);
    }
    let remainingTweet=newTweet.join('\n');

    // Extract all the tags
    const tagRegexp = /#[a-z0-9A-Z_]+/g;

    let tags = remainingTweet.matchAll(tagRegexp)

    // Split the url to extract the tags query.
    let url = new URL(baseURL)

    let tagQuery : string[]
    let tagQueryStr = url.searchParams.get("tags")
    tagQuery = tagQueryStr?(tagQueryStr.split(",")):[]

    interface FoundDict {
        [index: string]: boolean
    }

    let tagReplaced : FoundDict = {}


    let fulltagObj
    while (fulltagObj = tags.next()) {
        if (!fulltagObj.value) {
            break
        }
        let fulltag = fulltagObj.value[0]
        let tag = fulltag.slice(1)
        let tagQCopy: string[] = []
        tagQCopy.push(tag)

        let newURL= new URL(baseURL)
        newURL.searchParams.delete("id")
        newURL.searchParams.set("tags", tagQCopy.join(","))

        let newURLStr = newURL.toString()
        if (!tagReplaced[tag]) {
            remainingTweet = remainingTweet.replaceAll(fulltag, `[${fulltag}](${newURLStr})`)
            tagReplaced[tag] = true
        }
    }
    return firstline+"\n"+remainingTweet

}

export function formatTweet(tweet: string, baseURL: string):ReactElement {
    const converter = new showdown.Converter();
    
    // The first line is the command line if it exists
    // Default is for it not to exist.
    let commandLineExists = false;
    let kamalFont = false;

    if (isFlagOn(tweet, "#font:kamal")) {
        commandLineExists = true;
        kamalFont = true;
    }

    tweet = tagsToHyperlinks(tweet, baseURL, commandLineExists)
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
