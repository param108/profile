"use client";
import { ReactElement } from "react";
import { TweetType } from "./apis/tweets";
const showdown = require('showdown');
import './strings.css';


function isFlagOn(tweet: string, flag: string): Boolean {
    return tweet.includes(flag);
}

const threadRegexp=/#thread:([a-z0-9-]{36}):([0-9]{1,})/g;

// merge the retrieved tweets with the existing.
export function mergeTweets(oldTweets: TweetType[],newTweets: TweetType[], reverse: boolean): TweetType[] {
    let found :{[k: string]: boolean} ={}
    let final: TweetType[]= [];

    // newTweets first as they maybe updated
    newTweets.forEach((x)=>{
        found[x.id]=true;
        final.push(x);
    });

    // merge
    oldTweets.forEach((x)=>{
        if (found[x.id]) {
            return
        }
        final.push(x)
    })

    // finally sort
    final.sort((x,y)=>{
        let dx = new Date(x.created_at);
        let dy = new Date(y.created_at);

        if (dx > dy) {

            if (reverse) {
                // reversed case
                // x is later means x should be later in the list
                return 1;
            }
            // x is later means x should be before in the list
            return -1;
        }

        if (dx < dy) {
            if (reverse) {
                // reversed case
                // x is earlier means x should be before in the list
                return -1;
            }
            // y is later so y should be before in the list
            return 1;
        }

        return 0
    })
    return final;
}


// searches for hyperlinks in second lines and converts them
// for display as links
export function tagsToHyperlinks(tweet: string, baseURL: string): string {
    // Ignore tags on the first line
    let remainingTweet=tweet;

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
    return remainingTweet

}

export type CommandLineData = {
    exists: boolean,
    fontKamal: boolean,
    hasThread: boolean,
    threads: ThreadInfo[]
}

export function parseCommandLine(flags: string):CommandLineData {
    // TODO Need to move this to a common place.
    // The first line is the command line if it exists
    // Default is for it not to exist.
    let commandLineExists = false;

    let ret: CommandLineData = {
        exists: false,
        fontKamal: false,
        hasThread: false,
        threads: []
    }

    if (isFlagOn(flags, "#font:kamal")) {
        commandLineExists = true;
        ret.fontKamal = true;
    }

    ret.exists = false
    let threads = hasThread(flags)
    if (threads.length > 0) {
        commandLineExists = true;
        ret.hasThread = true;
        ret.threads = threads;
    }


    if (commandLineExists) {
        ret.exists = true
    }

    return ret
}

// Add thread to the commandline of a tweet
export function addThread(flags: string, thread_id: string):string {
    let cmdLine = parseCommandLine(flags);

    let cmd = "";
    if (cmdLine.exists) {
        let parts = flags.split("\n");
        cmd = parts[0];
        parts.splice(0,1);
        flags = parts.join("\n");
        cmd = cmd + ` #thread:${thread_id}:0\n`;
    } else {
        cmd = `#thread:${thread_id}:0\n`;
    }

    return cmd + flags;
}

// sort the thread of tweets in order of sequence number
export function sortThreadTweets(tweets: TweetType[], thread_id: string):TweetType[] {
    let foundTweets:{[key: string]:number} = {}

    return tweets.sort((x,y) => {
        let seqX = 0;
        let seqY = 0;
        if (x.id in foundTweets) {
            seqX = foundTweets[x.id]
        } else {
            let threadsInfoX = hasThread(x.flags);
            let threadInfos = threadsInfoX.filter((t)=>(t.id === thread_id));
            if (threadInfos.length > 0) {
                seqX =threadInfos[0].seq
            };
        }

        if (y.id in foundTweets) {
            seqY = foundTweets[y.id]
        } else {
            let threadsInfoY = hasThread(y.flags);
            let threadInfos = threadsInfoY.filter((t)=>(t.id === thread_id));
            if (threadInfos.length > 0) {
                seqY =threadInfos[0].seq
            };

        }

        if (seqX > seqY) {
            return 1;
        }

        if (seqX < seqY) {
            return -1;
        }

        return 0
    })
}

export function formatTweet(tweet: string, baseURL: string, flags: string):ReactElement {
    const converter = new showdown.Converter();
    
    // The first line is the command line if it exists
    // Default is for it not to exist.
    let cmdLine = parseCommandLine(flags);

    tweet = tagsToHyperlinks(tweet, baseURL)

    let hdata = converter.makeHtml(tweet);
    let classNames = "";
    if (cmdLine.fontKamal) {
        classNames = classNames + " font-kamal"
    }

    return (
        <div>
            <div className={classNames} dangerouslySetInnerHTML={{__html: hdata}}>
            </div>
        </div>
    )
}

export type ThreadInfo = {
    id: string
    seq: number
    name: string
}

// hasThread: Returns the threadID and seq for each tweet in a thread
// returns the tweets in seq order
export function hasThread(flags: string): ThreadInfo[] {
    const firstLine=flags.split("\n")[0];
    var data: ThreadInfo[] = []

    var matches = firstLine.matchAll(threadRegexp)
    var match = matches.next();

    while (!match.done) {
        var d = {
            id: match.value[1],
            seq: parseInt(match.value[2]),
            name: "" // this will be filled when the api returns
        };
        data.push(d);
        match = matches.next();
    }
    return data.sort((x,y)=>{
            if (x.seq >= y.seq) {
                return 1;
            }
            return -1
        })
}
