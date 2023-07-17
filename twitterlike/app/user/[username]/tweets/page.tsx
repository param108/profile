"use client";
import { getProfile } from "@/app/apis/login";
import { deleteTweet, getTweetsForUser, sendTweet, TweetType, updateTweet } from "@/app/apis/tweets";
import Editor from "@/app/components/editor";
import EditPair from "@/app/components/edit_tweet_pair";
import Header from "@/app/components/header";
import Tweet from "@/app/components/tweet";
import { AxiosResponse } from "axios";
import { useParams, useSearchParams } from "next/navigation";
import { Dispatch, SetStateAction, useCallback, useEffect, useState } from "react";
import { FiZap } from "react-icons/fi";
import ReactModal from "react-modal";
import { RingLoader } from "react-spinners";
import _ from "underscore";

const largeEditModalStyle = {
    content: {
        left: "25%",
        right: "25%",
        top: "100px"
    }
};

const bigEditModalStyle = {
    content: {
        left: "10%",
        right: "10%",
        top: "100px"
    }
};

const smallEditModalStyle = {
    content: {
        left: "2%",
        right: "2%",
        top: "100px"
    }
};

const largeDelModalStyle = {
    content: {
        left: "25%",
        right: "25%",
        top: "100px"
    }
};

const bigDelModalStyle = {
    content: {
        left: "10%",
        right: "10%",
        top: "100px"
    }
};

const smallDelModalStyle = {
    content: {
        left: "2%",
        right: "2%",
        top: "100px"
    }
};

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
    // should we show the tweet for the editor?
    var [ showEditorTweet, setShowEditorTweet ] = useState(false)
    var [ edittableTweet, setEdittableTweet ] = useState("")
    var [ editTweetValue, setEditTweetValue ] = useState("")
    var [ editTweetLoading, setEditTweetLoading ] = useState(false)
    var [ editTweetErrorMessage, setEditTweetErrorMessage ] = useState("")
    var [ editTweetShowError, setEditTweetShowError ] = useState(false)
    var [ delTweetLoading, setDelTweetLoading ] = useState(false)
    var [ delTweetErrorMessage, setDelTweetErrorMessage ] = useState("")
    var [ delTweetShowError, setDelTweetShowError ] = useState(false)
    var [ queryTags, setQueryTags ] :[ string[], Function ] = useState([])
    var [ pageLoading, setPageLoading ] = useState(false)
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
                    onClick={()=>setEditTweetShowError(false)}>
                        {editTweetErrorMessage}
                    </div>):null
                }

                <span className="text-black mb-[10px]">Are you sure you want to delete this tweet ?</span>
                <Tweet
                tweet_id={chosenTweet?.id}
                tweet={chosenTweet?.tweet}
                date={chosenTweet?.created_at}
                deleteClicked={()=>{}}
                editClicked={()=>{}}
                showMenu={false}
                onClick={()=>{}}
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

    // merge the retrieved tweets with the existing.
    function mergeTweets(oldTweets: TweetType[],newTweets: TweetType[]): TweetType[] {
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
                // x is later means x should be before in the list
                return -1;
            }

            if (dy < dx) {
                // y is later so y should be before in the list
                return 1;
            }

            return 0
        })
        return final;
    }


    // Once in the beginning
    useEffect(()=>{
        const tagStr = searchParams.get("tags")?.trim()
        let tags:string[] = []

        if (tagStr && tagStr.length > 0) {
            let newTags = tagStr.split(",")
            newTags.forEach((x)=>tags.push(x.trim()))
        }

        setQueryTags(tags)

        setPageLoading(true);
        getTweetsForUser([params.username], tags, 0).
            then((res:AxiosResponse)=>{
                console.log(res.data.data)
                setTweets(mergeTweets(tweets, res.data.data))
                setUsername(params.username)
                setPageLoading(false)
            }).
            catch(()=>{
                setErrorMessage("Failed to get tweets.")
                setShowError(true)
                setPageLoading(false)
            });
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [params.username])

    // Add infinite scroll!
    useEffect(()=> {
        const infiniteScroll = () => {
            // End of the document reached?
            if (window.innerHeight + document.documentElement.scrollTop
                >= document.documentElement.offsetHeight){
                setPageLoading(true)
                getTweetsForUser([params.username], queryTags, tweets.length).
                    then((res:AxiosResponse)=>{
                        setTweets(mergeTweets(tweets, res.data.data))
                        setUsername(params.username)
                        setPageLoading(false)
                    }).
                    catch(()=>{
                        setErrorMessage("Failed to get tweets.")
                        setShowError(true)
                        setPageLoading(false)
                    });
            }
        }

        window.removeEventListener('scroll', infiniteScroll);
        window.addEventListener('scroll', infiniteScroll, { passive: true });
        return () => window.removeEventListener('scroll', infiniteScroll);
    }, [params.username, tweets, queryTags])

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
                    getTweetsForUser([params.username], [], 0).
                        then((res:AxiosResponse)=>{
                            setTweets(mergeTweets(tweets, res.data.data))
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
            setTweets(mergeTweets(tweets,[res.data.data]))
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
                setDelTweetLoading(false)
            }).
            catch(()=>{
                setDelTweetErrorMessage("Something went wrong")
                setDelTweetShowError(true)
                setDelTweetLoading(false)
            })
    }

    return (
        <main className="flex bg-white min-h-screen w-full flex-col items-center justify-stretch">
            <Header showSpinner={pageLoading}></Header>
            {loggedIn?(
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
                    editClicked={()=>{}}
                    deleteClicked={()=>{}}
                    editorHideable={false}
                    hideClicked={()=>{}}
                    url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                                   ></EditPair>):(
                <div className="mt-[60px] mb-[10px]">
                    <span className="text-pink-600 cursor-pointer" onClick={()=>(location.href=`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets`)}>{username}</span>
                </div>
            )}
            { showError?(
                <div
                    className="p-[5px] bg-red-200 rounded mb-[5px]"
                    onClick={()=>setShowError(false)}>
                    {errorMessage}
                </div>):null
            }
            { tweets.length > 0 ?
                tweets.map((k: TweetType ,idx : number)=>{
                    return (
                        <Tweet
                        key={idx}
                        tweet_id={k.id}
                        tweet={k.tweet}
                        date={k.created_at}
                        showMenu={loggedIn}
                        onClick={()=>{}}
                        editClicked={()=>{onEditTweetClicked(k)()}}
                        deleteClicked={()=>{onDeleteTweetClicked(k)()}}
                        url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                        ></Tweet>
                    )
                }) : (
                    <Tweet tweet_id={"1"}  tweet={`
Nothing here **yet**!`} key={1} date="Start of time"
                    onClick={()=>{}}
                    editClicked={()=>{}}
                    deleteClicked={()=>{}}
                    showMenu={false}
                    url={`${process.env.NEXT_PUBLIC_HOST}/user/${username}/tweets?`+searchParams.toString()}
                        />
                )
            }
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

        </main>
    )
}
