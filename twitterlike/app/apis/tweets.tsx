import axios from "axios"
import "./interceptors";


export type TweetType = {
    tweet: string,
    created_at: string,
    id: string,
    image: string,
}

export type SignedURLRequest = {
    tweet_id: string,
    // local filename
    filename: string
}

export type SignedURLResponse = {
    data: {
        url: string,
        headers: {[key:string]:string}
    },
    success: boolean
}

export type UpdateTweetResponse = {
    data: TweetType
    success: boolean
}

export type DeleteTweetResponse = {
    data: TweetType
    success: boolean
}

export type Tweets = {
    data: TweetType[]
    success: boolean
}

export const getTweetsForUser = async (users:string[], tags:string[], offset:number, reverse:boolean) => {
    const config = {
        params: {
            users: users.join(","),
            tags: tags.join(","),
            offset,
            reverse
        },
        retry: 3
    };

    const res = await axios.get<Tweets>(
       `${process.env.NEXT_PUBLIC_BE_URL}/tweets`,
        config
    );
    return res;
}

export const getATweetForUser = async (user:string, tweet_id:string) => {
    const config = {
        params: {
            user: user,
        },
        retry: 3
    };

    const res = await axios.get<Tweets>(
       `${process.env.NEXT_PUBLIC_BE_URL}/tweet/${tweet_id}`,
        config
    );
    return res;
}

export const sendTweet = async (token: string, tweet: string, image: string) => {
    const config = {
        headers:{
            "TRIBIST_JWT": token,
        },
        retry: 3
    };

    const res = await axios.post<TweetType>(
       `${process.env.NEXT_PUBLIC_BE_URL}/tweets`,
        {
            tweet: tweet,
            image: image
        },
        config
    );
    return res;

}

export const updateTweet = async (token: string, tweet: string, tweet_id: string) => {
    const config = {
        headers:{
            "TRIBIST_JWT": token,
        },
        retry: 3
    };

    const res = await axios.put<UpdateTweetResponse>(
       `${process.env.NEXT_PUBLIC_BE_URL}/tweets`,
        {
           tweet: tweet,
           tweet_id: tweet_id
        },
        config
    );
    return res;
}

export const deleteTweet = async (token: string, tweet_id: string) => {
    const config = {
        headers:{
            "TRIBIST_JWT": token,
        },
        retry: 3
    };

    const res = await axios.post<DeleteTweetResponse>(
       `${process.env.NEXT_PUBLIC_BE_URL}/tweets/delete`,
        {
           tweet_id: tweet_id
        },
        config
    );
    return res;
}

export const signedURL = async (token: string, filename: string) => {
    const config = {
        headers:{
            "TRIBIST_JWT": token,
        },
        params: {
            suffix: filename.split(".").reverse().pop(),
        },
        retry: 3
    };

    const res = await axios.get<SignedURLResponse>(
       `${process.env.NEXT_PUBLIC_BE_URL}/signed_image_url`,
        config
    );
    return res;
}

export const uploadPhoto = async (url: string, headers: {[key:string]:string}, fileData: FormData) => {
    const config = {
        headers:{
            ...headers
        },
        retry: 3
    };

    const res = await axios.put<SignedURLResponse>(
       url,
        fileData,
        config
    );
    return res;
}
