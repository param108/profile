"use client";
import Tweet from './components/tweet'
import Header from './components/header'
import { useRouter } from 'next/navigation';
import { useRouter as useNextRouter } from 'next/router'
import Editor from './components/editor';
import { MouseEventHandler, useEffect, useState } from 'react';
import { getProfile } from './apis/login';

const loggedInTweets = [

]

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

export default function Home() {
  const router = useRouter()
  //const nextRouter = useNextRouter()
  var [ APIToken, setAPIToken ] = useState("")
  var [ loggedIn, setLoggedIn ] = useState(false)
  var [ username, setUsername ] = useState("")
  var [ darkMode, setDarkMode ] = useState("dark")

  useEffect(()=>{
    setDarkMode((localStorage.getItem('dark_mode')=="dark")?"dark":"")
  },[])


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
    <main className={darkMode +" flex bg-white dark:bg-slate-900 min-h-screen flex-col items-center justify-stretch"}>
      <Header changeDarkMode={(mode:string)=>{
        setDarkMode(mode)
        document.body.className = ((mode=="dark")?"dark":"")
      }} showSpinner={false}></Header>
      <div className='mt-[60px]'/>
      <div className='flex flex-col w-full items-center' onClick={(e)=>{
        if (loggedIn) {
          location.href=`/user/${username}/tweets`;
          e.stopPropagation()
        }}}>
      <Editor isLoggedIn={loggedIn}
        onChange={()=>{}}
        onFlagsChange={()=>{}}
        onSendClicked={()=>{}}
        showLoading={false}
        value={""}
        flags={""}
        hideable={false}
        hideClicked={()=>{}}
        defaultMessage={`
This is a blog. A **blog** of _tweets_.
Used to be called **micro-blogging** until twitter
**Hijacked** the space.
`}
        url={`${process.env.NEXT_PUBLIC_HOST}/`}
      ></Editor>
      </div>
      {
        welcomeTweets.map((k,idx)=>{
          return (<Tweet tweet_id={idx.toString()} key={idx} tweet={k.tweet}
                  date={k.created_at} onClick={()=>{}}
                  deleteClicked={()=>{}} editClicked={()=>{}}
                  flags={''}
                  showMenu={false}
                  visible={true}
                  externalClicked={null}
                  viewThread={null}
                  threadList={[]}
                  url={`${process.env.NEXT_PUBLIC_HOST}/`}
            ></Tweet>)
        })
      }
    </main>
  )
}
