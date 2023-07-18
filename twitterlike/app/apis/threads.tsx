import axios from "axios"
import "./interceptors";
import { TweetType } from "./tweets";

export type Thread = {
    id: string,
    user_id: string,
    created_at: string,
    deleted: boolean,
    writer: string,
    name: string
}

export type ThreadData = {
    id: string,
    user_id: string,
    created_at: string,
    deleted: boolean,
    writer: string,
    name: string
    tweets: TweetType[]
}

export type ThreadDataResponse = {
    data: ThreadData,
    success: boolean
}

export const getThread = async (username: string, thread_id: string) => {
    const config = {
        headers:{},
        retry: 3
    }

    const res = await axios.get<ThreadDataResponse>(
        `${process.env.NEXT_PUBLIC_BE_URL}/user/${username}/threads/${thread_id}`,
        config
    );
    return res;
}
