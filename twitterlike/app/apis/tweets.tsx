import axios from "axios"
import "./interceptors";


export type TweetType = {
    tweet: string,
    time: string,
    id: string
}

export type Tweets = {
    data: TweetType[]
    success: boolean
}

export const getTweetsForUser = async (users:string[], tags:string[], offset:number) => {
    const config = {
        params: {
            users: users.join(","),
            tags: tags.join(","),
            offset,
        },
        retry: 3
    };

    const res = await axios.get<Tweets>(
       'https://data.tribist.com/tweets',
        config
    );
    return res;
}
