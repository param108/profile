"use client";
import { getProfile } from "@/app/apis/login";
import { getTweetsForUser, TweetType } from "@/app/apis/tweets";
import Editor from "@/app/components/editor";
import Header from "@/app/components/header";
import Tweet from "@/app/components/tweet";
import { AxiosResponse } from "axios";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";


export default function ShowTweet() {
    const params = useParams();
    const router = useRouter();
    var [ APIToken, setAPIToken ] = useState("")
    var [ loggedIn, setLoggedIn ] = useState(false)
    var [ username, setUsername ] = useState("")
    var [ tweets, setTweets ] = useState<TweetType[]>([])
    var [ startOffset, setStartOffset ] = useState(0)
    var [ errorMessage, setErrorMessage ] = useState("");
    var [ showError, setShowError ] = useState(false);
    useEffect(()=>{
        const token = localStorage.getItem('api_token');
        if (token && token.length > 0) {
            setAPIToken(token)
        }
    }, [])

    useEffect(()=> {
        // users, tags
        getTweetsForUser([params.username], [], startOffset).
            then((res:AxiosResponse)=>{
                setTweets(res.data.data)
                setStartOffset(startOffset+res.data.data.length)
            }).
            catch(()=>{
                setErrorMessage("Failed to get tweets.")
                setShowError(true)
            });
    }, [])

    useEffect(()=>{
        if (APIToken.length == 0) {
            return
        }
        getProfile(APIToken).
            then((res: AxiosResponse)=>{
                setLoggedIn(res.data.data.username === params.username)
                setUsername(res.data.data.username)
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
    return (
        <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
            <Header></Header>
            {loggedIn?(
                <Editor isLoggedIn={true} defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`           }></Editor>):(
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
                            date={k?.time}></Tweet>
                    )
                }) : (
                    <Tweet router={router} tweet_id={"1"}  tweet={`
Nothing here **yet**!`} key={1} date="Start of time"/>
                )
            }
        </main>
    )
}
