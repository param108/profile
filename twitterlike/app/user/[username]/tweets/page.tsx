"use client";
import { getProfile } from "@/app/apis/login";
import { getTweetsForUser, sendTweet, TweetType } from "@/app/apis/tweets";
import Editor from "@/app/components/editor";
import EditPair from "@/app/components/edit_tweet_pair";
import Header from "@/app/components/header";
import Tweet from "@/app/components/tweet";
import { AxiosResponse } from "axios";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import _ from "underscore";

export default function ShowTweet() {
    const params = useParams();
    const router = useRouter();
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

    // this has the tweet_id of the tweet which is being editted
    // if any
    var [ showTweetEditor, setShowTweetEditor ] = useState("")
    console.log("rerendering user-tweet-page");

    useEffect(()=>{
        const token = localStorage.getItem('api_token');
        if (token && token.length > 0) {
            setAPIToken(token)
        }
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

    const infiniteScroll = useCallback(() => {
        // End of the document reached?
        if (window.innerHeight + document.documentElement.scrollTop
            === document.documentElement.offsetHeight){
            getTweetsForUser([params.username], [], tweets.length).
                then((res:AxiosResponse)=>{
                    setTweets(mergeTweets(tweets, res.data.data))
                    setUsername(params.username)
                }).
                catch(()=>{
                    setErrorMessage("Failed to get tweets.")
                    setShowError(true)
                });
        }
    }, [params.username, tweets])

    // Once in the beginning
    useEffect(()=>{
        getTweetsForUser([params.username], [], 0).
            then((res:AxiosResponse)=>{
                setTweets(mergeTweets(tweets, res.data.data))
                setUsername(params.username)
            }).
            catch(()=>{
                setErrorMessage("Failed to get tweets.")
                setShowError(true)
            });
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [params.username])

    // Add infinite scroll!
    useEffect(()=> {
        window.addEventListener('scroll', infiniteScroll)
    }, [infiniteScroll])

    useEffect(()=>{
        if (APIToken.length == 0) {
            return
        }
        getProfile(APIToken).
            then((res: AxiosResponse)=>{
                setLoggedIn(res.data.data.username === params.username)
                localStorage.setItem('username', res.data.data.username)
                localStorage.setItem('user_id', res.data.data.user_id)
            }).
            catch(()=>{
                // clear out the api_token
                localStorage.removeItem('api_token')
                setErrorMessage("Login Failure. Please login again.")
                setShowError(true)
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

    return (
        <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
            <Header></Header>
            {loggedIn?(
                <EditPair editting={true} isLoggedIn={true} showLoading={editorLoading}
                    onSendClicked={onSendClicked} value={editorValue} viewing={showEditorTweet}
                    onChange={onChanged} key={1} tweet={{
                        created_at: "Preview",
                        id: 'new',
                    }}
                    defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`           }></EditPair>):(
                <div className="mt-[60px] mb-[10px]">
                    <span className="text-pink-600">{username}</span>
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
                        <Tweet router={router} tweet_id={k.id} key={idx} tweet={k?.tweet}
                            date={k?.created_at}></Tweet>
                    )
                }) : (
                    <Tweet router={router} tweet_id={"1"}  tweet={`
Nothing here **yet**!`} key={1} date="Start of time"/>
                )
            }
        </main>
    )
}
