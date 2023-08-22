"use client";
import { getProfile } from "@/app/apis/login";
import { deleteTweet, getATweetForUser, getTweetsForUser, sendTweet, TweetType, updateTweet } from "@/app/apis/tweets";
import Editor from "@/app/components/editor";
import EditPair from "@/app/components/edit_tweet_pair";
import Header from "@/app/components/header";
import Tweet from "@/app/components/tweet";
import { addThread, hasThread, mergeTweets, ThreadInfo } from "@/app/strings";
import { AxiosResponse } from "axios";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { UIEvent, Dispatch, SetStateAction, useCallback, useEffect, useState } from "react";
import { FiEdit3, FiExternalLink, FiZap } from "react-icons/fi";
import ReactModal from "react-modal";
import { RingLoader } from "react-spinners";
import _ from "underscore";
import { createThread, getThread, ThreadData } from "@/app/apis/threads";

const largeEditModalStyle = {
    content: {
        left: "25%",
        right: "25%",
        top: "100px",
        background: "#64748b"
    }
};

const bigEditModalStyle = {
    content: {
        left: "10%",
        right: "10%",
        top: "100px",
        background: "#64748b"
    }
};

const smallEditModalStyle = {
    content: {
        left: "2%",
        right: "2%",
        top: "100px",
        background: "#64748b"
    }
};

const largeDelModalStyle = {
    content: {
        left: "25%",
        right: "25%",
        top: "100px",
        background: "#64748b"
    }
};

const bigDelModalStyle = {
    content: {
        left: "10%",
        right: "10%",
        top: "100px",
        background: "#64748b"
    }
};

const smallDelModalStyle = {
    content: {
        left: "2%",
        right: "2%",
        top: "100px",
        background: "#64748b"
    }
};

    const welcomeTweets = [
  {
    created_at: `At the beginning.`,
    tweet: `
I think in tweets.

*Short paragraphs of thought*

**Shorter the Better**
`},
  {
    created_at: `A little later.`,
    tweet: `
These thoughts could be **independent**

OR

They could be *connected* as **threads** or
*related* through **#tags**
`},
  {
    created_at: `Even later.`,
    tweet: `
I re-read my tweets a lot. Over & Over.

Sometimes **Narcissism** & sometimes to **remind** me

of things I already know.
`},
  {
    created_at: `Even later....er.`,
    tweet: `
At times I want to **explore** them and **discover** new connections,
or new **insights** or wallow in old ones.

I like **high-lighting** and _italics_.
Did I mention, we support **Markdown!**"
`},
  {
    created_at: `Right Here, Right Now.`,
    tweet: `
You can do all this here and you own your data,
download as you wish.

Unlike twitter this is not a **performance**,

this is **recreation**. This is **expression**.

This is **Freedom**!

_Interested ?_

Then [**signup**](${process.env.NEXT_PUBLIC_BE_URL}/users/login?source=twitter&redirect_url=/)!
`}
];

export default function ShowTweet() {
    const params = useParams();
    const searchParams = useSearchParams();
    var [ APIToken, setAPIToken ] = useState("")
    var [ loggedIn, setLoggedIn ] = useState(false)
    var [ editorLoading, setEditorLoading ] = useState(false)
    var [ editorValue, setEditorValue ] = useState("")
    var [ username, setUsername ] = useState("")
    var [ tweets, setTweets ] = useState<TweetType[]>([])
    var [ errorMessage, setErrorMessage ] = useState("");
    var [ showError, setShowError ] = useState(false);
    var [ threadCatalog, setThreadCatalog ] = useState<{[name:string]:(ThreadData|null)}>({})
    // should we show the tweet for the editor?
    var [ showEditorTweet, setShowEditorTweet ] = useState(false)
    var [ editTweetValue, setEditTweetValue ] = useState("")
    var [ editTweetLoading, setEditTweetLoading ] = useState(false)
    var [ editTweetErrorMessage, setEditTweetErrorMessage ] = useState("")
    var [ editTweetShowError, setEditTweetShowError ] = useState(false)
    var [ delTweetLoading, setDelTweetLoading ] = useState(false)
    var [ delTweetErrorMessage, setDelTweetErrorMessage ] = useState("")
    var [ delTweetShowError, setDelTweetShowError ] = useState(false)
    var [ queryTags, setQueryTags ] :[ string[], Function ] = useState([])
    // thread control
    var [ threadVisible, setThreadVisible ] = useState(false)
    var [ threadData, setThreadData ] = useState<ThreadData|null>(null)
    var [ pageLoading, setPageLoading ] = useState(false)
    var [ reverseFlag, setReverseFlag ] = useState(false)
    var [ darkMode, setDarkMode ] = useState("dark")

    useEffect(()=>{
        setDarkMode((localStorage.getItem('dark_mode')=="dark")?"dark":"")
    },[])

    console.log("rerendering user-tweet-page");
    // Which modal is open
    var [openModal, setOpenModal] = useState("")

    // holds the tweet id of the tweet that is to be editted or deleted.
    var [chosenTweet, setChosenTweet]:[TweetType, Dispatch<SetStateAction<TweetType>>] = useState({
        tweet:"",
        id:"",
        created_at:""
    })

    const editTweetDiv = ()=>{
        return (
            <div className="flex flex-col items-center w-full">
            { editTweetShowError?(
                <div
                    className="p-[5px] bg-red-200 rounded mb-[5px]"
                    onClick={()=>setEditTweetShowError(false)}>
                    {editTweetErrorMessage}
                </div>):null
            }
            <EditPair editting={true} isLoggedIn={true} showLoading={editTweetLoading}
            onSendClicked={onEditTweetSendClicked} value={editTweetValue}
            viewing={true}
            onChange={onEditTweetChanged} key={1} tweet={chosenTweet}
            showMenu={false}
            defaultMessage={`
Unknown Tweet`}
            editClicked={onEditTweetSendClicked}
            deleteClicked={()=>{}}
            editorHideable={false}
            hideClicked={()=>{}}
            visible={true}
            url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                />
            </div>)

    }

    const deleteTweetDiv = ()=>{
        return (
            <div className="flex flex-col items-center w-full">
                { delTweetShowError?(
                    <div
                    className="p-[5px] bg-red-200 rounded mb-[5px]"
                    onClick={()=>setDelTweetShowError(false)}>
                        {delTweetErrorMessage}
                    </div>):null
                }

                <span className="text-black mb-[10px]">Are you sure you want to delete this tweet ?</span>
                <Tweet
                visible={true}
                tweet_id={chosenTweet?.id}
                tweet={chosenTweet?.tweet}
                date={chosenTweet?.created_at}
                deleteClicked={()=>{}}
                editClicked={()=>{}}
                showMenu={false}
                onClick={()=>{}}
                externalClicked={null}
                threadList={[]}
                viewThread={null}
                url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                />
                <div className="w-[90%] md:w-[510px] mt-[10px]">
                <div className="inline-block float-right pr-[10px]">
                    <RingLoader className="inline-block" color="#EC4899"
                        loading={delTweetLoading} size={30}/>
                </div>
                <button onClick={delTweetClicked} className="rounded p-[5px] bg-red-800 text-white float-left">Delete</button>
                </div>
            </div>
        )
    }

    const router = useRouter();
    var [createThreadName, setCreateThreadName ] = useState("")
    var [createThreadLoading, setCreateThreadLoading] = useState(false)
    var [createThreadErrorMessage, setCreateThreadErrorMessage] = useState("")
    var [createThreadShowError, setCreateThreadShowError] = useState(false)
    const createThreadClicked = () => {
        setCreateThreadLoading(true)
        createThread(APIToken, username, createThreadName).
            then((res) => {
                let thread_id = res.data.data.id

                // update the tweet with the thread hash
                chosenTweet.tweet = addThread(chosenTweet.tweet, thread_id);
                updateTweet(APIToken, chosenTweet.tweet, chosenTweet.id).
                    then((res)=>{
                        setTweets(mergeTweets(tweets,[res.data.data], reverseFlag))
                        setOpenModal("closed")
                    }).catch(()=>{
                        setCreateThreadErrorMessage("Something went wrong")
                        setCreateThreadShowError(true)
                    }).finally(()=>{
                        setCreateThreadLoading(false)
                    })
            }).
            catch(() => {
                setCreateThreadErrorMessage("Something went wrong")
                setCreateThreadShowError(true)
            }).
            finally(() => {
                setCreateThreadLoading(false)
            })
    }

    const createThreadDiv = ()=>{
        return (
            <div className="flex flex-col items-center w-full">
                { createThreadShowError?(
                    <div
                    className="p-[5px] bg-red-200 rounded mb-[5px]"
                    onClick={()=>setCreateThreadShowError(false)}>
                        {createThreadErrorMessage}
                    </div>):null
                }
                <span className="text-black mb-[10px]">Create a new thread</span>
                <Tweet
                visible={true}
                tweet_id={chosenTweet?.id}
                tweet={chosenTweet?.tweet}
                date={chosenTweet?.created_at}
                deleteClicked={()=>{}}
                editClicked={()=>{}}
                showMenu={false}
                onClick={()=>{}}
                externalClicked={null}
                viewThread={null}
                threadList={[]}
                url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                />
                <input type="text" className="rounded text-black p-[5px] mt-[5px] border border-slate-200 w-[96%] md:w-[510px]"
                    placeholder="New thread name..." value={createThreadName}
                    onChange={(t)=>{setCreateThreadName(t.target.value)}}/>
                <div className="w-[90%] md:w-[510px] mt-[10px]">
                <div className="inline-block float-right pr-[10px]">
                    <RingLoader className="inline-block" color="#EC4899"
                        loading={createThreadLoading} size={30}/>
                </div>
                <button
                    onClick={createThreadClicked}
                    className={(createThreadName.length > 3)?
                        "rounded p-[5px] bg-sky-600 text-white float-left":
                        "rounded p-[5px] bg-sky-100 text-white float-left"}>Create</button>
                </div>
            </div>
        )
    }

    const [largeScreen, setLargeScreen] = useState(
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 1024px)").matches
    )
    const [bigScreen, setBigScreen] = useState(
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 768px)").matches
    )

    useEffect(() => {
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 768px)").addEventListener('change', (e) => {
            setBigScreen( e.matches );
        });
        (typeof window === "undefined")?true:window.matchMedia("(min-width: 1024px)").addEventListener('change', (e) => {
            setLargeScreen( e.matches );
        });
    }, []);

    useEffect(()=>{
        const token = localStorage.getItem('api_token');
        if (token && token.length > 0) {
            setAPIToken(token)
        }
        ReactModal.setAppElement('body')
    }, [])


    // Once in the beginning
    useEffect(()=>{
        const tagStr = searchParams.get("tags")?.trim()
        const reverseStr = searchParams.get("reverse")?.trim()
        const idStr = searchParams.get("id")?.trim()

        let tags:string[] = []

        if (tagStr && tagStr.length > 0) {
            let newTags = tagStr.split(",")
            newTags.forEach((x)=>tags.push(x.trim()))
        }

        let reverse = false;
        if (reverseStr && reverseStr === "1") {
            reverse = true;
        }
        setReverseFlag(reverse)

        setQueryTags(tags)

        setPageLoading(true);

        if (idStr) {
            getATweetForUser(params.username, idStr).
                then((res:AxiosResponse)=>{
                console.log(res.data.data)
                setTweets(mergeTweets(tweets, res.data.data, reverse))
                setUsername(params.username)
                setPageLoading(false)
            }).
            catch(()=>{
                setErrorMessage("Failed to get tweets.")
                setShowError(true)
                setPageLoading(false)
            });
        } else {
        getTweetsForUser([params.username], tags, 0, reverse).
            then((res:AxiosResponse)=>{
                console.log(res.data.data)
                setTweets(mergeTweets(tweets, res.data.data, reverse))
                setUsername(params.username)
                setPageLoading(false)
            }).
            catch(()=>{
                setErrorMessage("Failed to get tweets.")
                setShowError(true)
                setPageLoading(false)
            });
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [params.username])

    const getThreadCatalog = () => {
        return threadCatalog;
    }

    useEffect(()=>{
        var seen:{ [name:string]:boolean } = {}
        var newThreads: { [name:string]:(ThreadData|null) } = {}
        tweets.forEach((tweet: TweetType)=>{
            let ts = hasThread(tweet.tweet);
            ts.forEach((t: ThreadInfo)=>{
                if (t.id in seen) {
                    return
                }
                seen[t.id] = true;

                if (!(t.id in threadCatalog)) {
                    // launch a request for a thread
                    getThread(username, t.id).then(
                        (res)=>{
                            let key = res.data.data.id;
                            setThreadCatalog((t)=> {
                                let data = {...t};
                                data[key] = res.data.data;
                                return data;
                            });

                        }
                    )
                    newThreads[t.id] = null;
                }
            })
        })

        setThreadCatalog((t)=>({
            ...t,
            ...newThreads
        }))

        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [username, tweets])

    // We need to go one layer deeper to populate the threadview
    useEffect(()=>{
        var seen:{ [name:string]:boolean } = {}
        var newThreads: { [name:string]:(ThreadData|null) } = {}
        let keys = Object.keys(threadCatalog);
        keys.forEach((threadKey: string)=>{
            if (threadCatalog && threadCatalog[threadKey] === null) {
                return;
            }
            let thread = threadCatalog[threadKey];
            if (thread === null) {
                return
            }

            thread.tweets.forEach((tweet: TweetType) => {
                let ts = hasThread(tweet.tweet);
                ts.forEach((t: ThreadInfo)=>{
                    if (t.id in seen) {
                        return
                    }
                    seen[t.id] = true;

                    if (!(t.id in threadCatalog)) {
                        // launch a request for a thread
                        getThread(username, t.id).then(
                            (res) => {
                                let key = res.data.data.id;
                                setThreadCatalog((t) => {
                                    let data = { ...t };
                                    data[key] = res.data.data;
                                    return data;
                                });

                            }
                        )
                        newThreads[t.id] = null;
                    }
                })
            })
        })

        if (Object.keys(newThreads).length > 0) {
            setThreadCatalog((t)=>({
                ...t,
                ...newThreads
            }))
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [username, threadCatalog])

    useEffect(()=>{
        if (APIToken.length == 0) {
            return
        }
        setPageLoading(true);
        getProfile(APIToken).
            then((res: AxiosResponse)=>{
                setLoggedIn(res.data.data.username === params.username)
                localStorage.setItem('username', res.data.data.username)
                localStorage.setItem('user_id', res.data.data.user_id)
                setPageLoading(false)
            }).
            catch(()=>{
                // clear out the api_token
                localStorage.removeItem('api_token')
                setErrorMessage("Login Failure. Please login again.")
                setShowError(true)
                setPageLoading(false)
            });
                }, [APIToken, params.username])

    useEffect(()=>{
        let seenThreads: {[key: string]:Boolean} = {};
        let effectiveTweetLength = 0;
        if (tweets.length < 20) {
            // If tweets are less than 20 long already then
            // probably we don't have enough tweets.
            return;
        }

        tweets.forEach((k: TweetType)=>{
            let threadsInfo = hasThread(k.tweet)
            let tweetVisible = false;

            // a tweet without a thread is always visible
            if (threadsInfo.length == 0) {
                tweetVisible = true;
            }
            // If a tweet is part of threads and each thread has already been represented
            // by a previous tweet then don't render this tweet
            threadsInfo.forEach((x)=>{
                if (x.id in seenThreads) {
                    return;
                }
                seenThreads[x.id] = true;
                tweetVisible = true;
                return;
            })

            if (tweetVisible) {
                effectiveTweetLength++;
            }

        })

        if (effectiveTweetLength < 20) {
            console.log("effectiveTweetLength", effectiveTweetLength);
            setPageLoading(true)
            getTweetsForUser([params.username], queryTags, tweets.length, reverseFlag).
                then((res:AxiosResponse)=>{
                    if (res.data.data.length > 0) {
                        setTweets(mergeTweets(tweets, res.data.data, reverseFlag))
                        setUsername(params.username)
                    }
                    setPageLoading(false)
                }).
                catch(()=>{
                    setErrorMessage("Failed to get tweets.")
                    setShowError(true)
                    setPageLoading(false)
                })
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [tweets])

    const onSendClicked= (tweet: string) => {
        if (loggedIn) {
            setEditorLoading(true)

            // save the value returned so that we don't lose it
            // if there is some error.
            setEditorValue(tweet)
            sendTweet(APIToken, tweet).
                then(()=>{
                    setEditorValue("")
                    setShowEditorTweet(false)
                    setEditorLoading(false)
                    getTweetsForUser([params.username], [], 0, reverseFlag).
                        then((res:AxiosResponse)=>{
                            setTweets(mergeTweets(tweets, res.data.data, reverseFlag))
                            setUsername(params.username)
                        }).
                        catch(()=>{
                            setErrorMessage("Failed to get tweets.")
                            setShowError(true)
                        });
                }).
                catch(()=>{
                    setShowError(true)
                    setErrorMessage("Failed to upload tweet")
                    setEditorLoading(false)
                })
        }
    }

    const onChanged= (newValue: string) => {
        setShowEditorTweet(newValue.length > 0)
        setEditorValue(newValue)
    }

    const onEditTweetChanged=(newValue: string) => {
        setEditTweetValue(newValue)
    }

    const onEditTweetClicked=(tweet: TweetType)=> {
        return () => {
            setOpenModal("edit_tweet")
            setChosenTweet(tweet)
            setEditTweetValue(tweet.tweet)
        }
    }

    const onEditTweetSendClicked=()=>{
        setEditTweetShowError(false)
        setEditTweetLoading(true)
        updateTweet(APIToken, editTweetValue, chosenTweet.id).
        then((res)=>{
            setEditTweetLoading(false)
            console.log(res.data)
            setTweets(mergeTweets(tweets,[res.data.data], reverseFlag))
            setOpenModal("closed")
        }).catch(()=>{
            setEditTweetErrorMessage("Something went wrong")
            setEditTweetShowError(true)
            setEditTweetLoading(false)
        })
    }

    const onDeleteTweetClicked=(tweet: TweetType) => {
        return () => {
            setChosenTweet(tweet)
            setOpenModal("del_tweet")
        }
    }

    const delTweetClicked=() => {
        setDelTweetLoading(true)
        deleteTweet(APIToken, chosenTweet.id).
            then((res)=>{
                setTweets(tweets.filter((x)=>(x.id != res.data.data.id)))
                setOpenModal("closed")
            }).
            catch(()=>{
                setDelTweetErrorMessage("Something went wrong")
                setDelTweetShowError(true)
            }).finally(()=>{
                setDelTweetLoading(false)
            })
    }

    const onScroll=(e: UIEvent<HTMLElement>)=>{
        let scrollTop = e.currentTarget.scrollTop;
        let scrollHeight = e.currentTarget.scrollHeight;
        let clientHeight = e.currentTarget.clientHeight;

        if (scrollTop + clientHeight > (0.9*scrollHeight)) {
            if (!pageLoading) {
                setPageLoading(true)
                getTweetsForUser([params.username], queryTags, tweets.length, reverseFlag).
                    then((res:AxiosResponse)=>{
                        setTweets(mergeTweets(tweets, res.data.data, reverseFlag))
                        setUsername(params.username)
                        setPageLoading(false)
                    }).
                    catch(()=>{
                        setErrorMessage("Failed to get tweets.")
                        setShowError(true)
                        setPageLoading(false)
                    })
            }
        }
    }

    return (
        <main onScroll={onScroll} className={
            darkMode + " flex bg-white dark:bg-slate-900 min-h-screen max-h-screen  w-full flex-col items-center justify-stretch" +
            (threadVisible?"":" overflow-y-scroll")}>
            <Header changeDarkMode={(mode:string)=>{
                setDarkMode(mode)
                document.body.className = ((mode=="dark")?"dark":"")
            }} showSpinner={pageLoading}></Header>
            <div className="mt-[60px]"/>
            <div className="flex flex-col items-center w-full">
                {showError ? (
                    <div
                        className="p-[5px] bg-red-200 rounded mb-[5px]"
                        onClick={() => setShowError(false)}>
                        {errorMessage}
                    </div>) : null
                }
            </div>

            <div className={"flex flex-row" +
                (threadVisible?" overflow-hidden":"")}>
            <div onScroll={onScroll} className={"max-h-full" +
                (threadVisible?" overflow-y-scroll": "")}>
                <div className="flex flex-col items-center">
                {loggedIn ? (
                    <EditPair editting={true} isLoggedIn={true} showLoading={editorLoading}
                    onSendClicked={onSendClicked} value={editorValue} viewing={showEditorTweet}
                    onChange={onChanged} key={10000} tweet={{
                        created_at: "Preview",
                        id: 'new',
                        tweet: ''
                    }}
                    showMenu={false}
                    defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`}
                    editClicked={() => { }}
                    deleteClicked={() => { }}
                    editorHideable={false}
                    hideClicked={() => { }}
                    visible={true}
                    url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?` + searchParams.toString()}
                        ></EditPair>) : (
                    <div className="mb-[10px]">
                        <span className="text-pink-600 cursor-pointer" onClick={() => (location.href = `${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets`)}>{username}</span>
                    </div>
                )}
                { (()=> {
                    let foundThread:{[key:string]:boolean} = {};
                    if (tweets.length > 0 ) {
                        return tweets.map((k: TweetType ,idx : number)=>{
                        let threadsInfo = hasThread(k.tweet)
                        let threads = threadsInfo.map((x:ThreadInfo)=>{
                            if (x.id in threadCatalog) {
                                return threadCatalog[x.id];
                            }
                            return null
                        })

                        // If a tweet is part of threads and each thread has already been represented
                        // by a previous tweet then don't render this tweet
                        if (threadsInfo.length > 0 && threadsInfo.every((x)=>{
                            if (x.id in foundThread) {
                                return true;
                            }
                            foundThread[x.id] = true;
                            return false;
                        })) {
                            return null
                        }

                        return (
                            <Tweet
                                visible={true}
                                key={idx}
                                tweet_id={k.id}
                                tweet={k.tweet}
                                date={k.created_at}
                                showMenu={loggedIn}
                                onClick={() => {
                                    //setThreadVisible(!threadVisible);
                                    //setThreadData(welcomeTweets);
                                }}
                                editClicked={() => { onEditTweetClicked(k)() }}
                                deleteClicked={() => { onDeleteTweetClicked(k)() }}
                                threadList={threads}
                                viewThread={(threadID: string) => {
                                    if (bigScreen || largeScreen) {
                                        setThreadData(threadCatalog[threadID])
                                        setThreadVisible(true)
                                    } else {
                                        location.href = `${process.env.NEXT_PUBLIC_HOST}/user/${username}/threads/${threadID}/`;
                                    }
                                }}
                                externalClicked={(tweet_id: string) => {
                                    location.href = `${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?id=${tweet_id}`;
                                }}
                                createThreadClicked={() => {
                                    setChosenTweet(k)
                                    setCreateThreadName("")
                                    setOpenModal("create_thread")
                                }}
                                url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?` + searchParams.toString()}
                            ></Tweet>
                        )
                    })}

                    // no tweets found return the placeholder.
                    return (
                            <Tweet tweet_id={"1"}  tweet={`
Nothing here **yet**!`} key={1} date="Start of time"
                            onClick={()=>{}}
                            editClicked={()=>{}}
                            deleteClicked={()=>{}}
                            visible={true}
                            showMenu={false}
                            threadList={[]}
                            externalClicked={null}
                            viewThread={null}
                            url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}/>
                        )
                })()}
            </div>
            </div>
            {(threadVisible && threadData)?(
            <div className="max-h-full ml-[10px] float-right overflow-y-scroll">
            <div className="w-[90%] md:w-[510px]">
                <span className="text-black dark:text-white text-xl">{">> "}<b>{threadData.name}</b></span>
                <FiZap
                onClick={()=>{setThreadVisible(false)}}
                className="cursor-pointer ml-[10px] text-pink-600 float-right" size={20}/>
                <FiExternalLink
                onClick={()=>{router.push(
                    `${process.env.NEXT_PUBLIC_HOST}/user/${username}/threads/${threadData?.id}/`)}}
                className="cursor-pointer ml-[10px] text-pink-600 float-right" size={20}/>
                {loggedIn?(
                    <FiEdit3
                    onClick={()=>{router.push(
                        `${process.env.NEXT_PUBLIC_HOST}/user/${username}/threads/${threadData?.id}/`)}}
                    className="cursor-pointer ml-[10px] text-pink-600 float-right" size={20}/>
                ):null}
            </div>
            <div className="flex flex-col items-center">
            {(()=>{
                if (threadData && threadData.tweets.length > 0) {
                    return threadData.tweets.map((k: TweetType ,idx : number)=>{
                        let threadsInfo = hasThread(k.tweet)
                        let threads = threadsInfo.map((x:ThreadInfo)=>{
                            if (x.id === threadData?.id) {
                                return null
                            }
                            if (x.id in threadCatalog) {
                                return threadCatalog[x.id];
                            }
                            return null
                        }).filter((x)=>(x !== null));

                        return (
                            <Tweet
                            visible={true}
                            key={idx}
                            tweet_id={k.id}
                            tweet={k.tweet}
                            date={k.created_at}
                            showMenu={loggedIn}
                            onClick={()=>{}}
                            editClicked={()=>{onEditTweetClicked(k)()}}
                            deleteClicked={()=>{onDeleteTweetClicked(k)()}}
                            url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                            externalClicked={(tweet_id:string)=>{
                                location.href = `${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?id=${tweet_id}`;
                            }}
                            viewThread={(threadID: string) => {
                                location.href = `${process.env.NEXT_PUBLIC_HOST}/user/${username}/threads/${threadID}/`;
                            }}

                            threadList={threads}
                                ></Tweet>
                        )
                    })} else {
                        return (
                            <Tweet tweet_id={"1"}  tweet={`
Nothing here **yet**!`} key={1} date="Start of time"
                            onClick={()=>{}}
                            editClicked={()=>{}}
                            deleteClicked={()=>{}}
                            externalClicked={null}
                            showMenu={false}
                            visible={true}
                            threadList={[]}
                            viewThread={null}
                            url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                                />
                        )
                    }})()}
            </div>
            </div>
            ):null }
            </div>
        { /* editTweetModal */ }
            <ReactModal
                style={largeScreen?largeEditModalStyle:(bigScreen)?bigEditModalStyle:smallEditModalStyle}
                isOpen={openModal=="edit_tweet"}>
                <FiZap
                    size={30}
                    className="text-pink-600 float-right"
                    onClick={()=>{setOpenModal("closed")}}/>
                    {editTweetDiv()}
            </ReactModal>
            { /* deleteTweetModal */ }
            <ReactModal
                style={largeScreen?largeDelModalStyle:(bigScreen)?bigDelModalStyle:smallDelModalStyle}
                isOpen={openModal=="del_tweet"}>
                <FiZap
                    size={30}
                    className="text-pink-600 float-right"
                    onClick={()=>{setOpenModal("closed")}}/>
                    {deleteTweetDiv()}
            </ReactModal>
            { /* CreateThreadModal */ }
            <ReactModal
                style={largeScreen?largeDelModalStyle:(bigScreen)?bigDelModalStyle:smallDelModalStyle}
                isOpen={openModal=="create_thread"}>
                <FiZap
                    size={30}
                    className="text-pink-600 float-right"
                    onClick={()=>{setOpenModal("closed")}}/>
                    {createThreadDiv()}
            </ReactModal>

        </main>
    )
}
