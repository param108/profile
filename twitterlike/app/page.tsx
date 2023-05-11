"use client";
import Image from 'next/image'
import Tweet from './components/tweet'
import Editor from './components/editor'
import Header from './components/header'
import { useState } from 'react'
import { useRouter } from 'next/navigation';

export default function Home() {
  const [tweets, setTweets] = useState(Array(10).fill(0));
  const router = useRouter();
  return (
    <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
      <Header></Header>
      <Editor isLoggedIn={false} defaultMessage={
        `#font:kamal
If these tweets **contribute** to you, 

please consider, **contributing** at 

[http://paymenow.com/param108](http://paymenow.com/param108)`}></Editor>
      {
        tweets.map((k,idx)=>{
          return (<Tweet router={router} tweet_id={"a-b-123211-bs"} key={idx} tweet={
`
The world is full of too many of these tweets.

There are too many of these tweets.

There are tweets and too many of them.

There are only tweets and lots of them.            

There is no tweet left to tweet.
`
            } date={
  `13:00 Thu`
}></Tweet>)
        })
      }
    </main>
  )
}
