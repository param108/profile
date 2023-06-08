"use client";
import Tweet from './components/tweet'
import Header from './components/header'
import { useRouter } from 'next/navigation';
import Editor from './components/editor';
import { MouseEventHandler, useEffect, useState } from 'react';
import { getProfile } from './apis/login';

const loggedInTweets = [

]

const welcomeTweets = [
  {
    time: `At the beginning.`,
    tweet: `
I think in tweets.

*Short paragraphs of thought*

**Shorter the Better**
`},
  {
    time: `A little later.`,
    tweet: `
These thoughts could be **independent**

OR

They could be *connected* as **threads** or
*related* through **#tags**
`},
  {
    time: `Even later.`,
    tweet: `
I re-read my tweets a lot. Over & Over.

Sometimes **Narcissism** & sometimes to **remind** me

of things I already know.
`},
  {
    time: `Even later....er.`,
    tweet: `
At times I want to **explore** them and **discover** new connections,
or new **insights** or wallow in old ones.

I like **high-lighting** and _italics_.
Did I mention, we support **Markdown!**"
`},
  {
    time: `Right Here, Right Now.`,
    tweet: `
You can do all this here and you own your data,
download as you wish.

Unlike twitter this is not a **performance**,

this is **recreation**. This is **expression**.

This is **Freedom**!

_Interested ?_

Then [**signup**](https://ui.tribist.com/users/login?source=twitter&redirect_url=/)!
`}
];

export default function Home() {
  const router = useRouter();
  var [ APIToken, setAPIToken ] = useState("")
  var [ loggedIn, setLoggedIn ] = useState(true)
  var [ username, setUsername ] = useState("")
  useEffect(()=>{
    const token = localStorage.getItem('api_token');
    if (token && token.length > 0) {
      setAPIToken(token)
    }
  }, [])

  useEffect(()=>{
    if (APIToken.length == 0) {
      return
    }

    getProfile(APIToken).
      then((res)=>{
        setLoggedIn(true)
        setUsername(res.data.data.username)
        localStorage.setItem('username', res.data.data.username)
        localStorage.setItem('user_id', res.data.data.user_id)
      }).
      catch(()=>{
        // clear out the api_token
        localStorage.removeItem('api_token')
      })
  }, [APIToken])
  return (
    <main className="flex bg-white min-h-screen flex-col items-center justify-stretch">
      <Header></Header>
      <div className='flex' onClick={(e)=>{
        if (loggedIn) {
          location.href=`/user/${username}/tweets`;
          e.stopPropagation()
        }}}>
      <Editor isLoggedIn={loggedIn} defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`}></Editor>
      </div>
      {
        welcomeTweets.map((k,idx)=>{
          return (<Tweet router={router} tweet_id={idx.toString()} key={idx} tweet={k.tweet}
                  date={k.time}></Tweet>)
        })
      }
    </main>
  )
}
