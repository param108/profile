import axios from "axios"
import { ReadableByteStreamController } from "node:stream/web";
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

export type ThreadDetails = {
    id: string,
    user_id: string,
    name: string,
    writer: string,
    created_at: string,
    deleted: boolean
}

export type ThreadDetailsResponse = {
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

export const createThread = async (token: string, username: string, name: string) => {
    const config = {
        headers:{
            "TRIBIST_JWT": token,
        },
        retry: 3
    }

    const res = await axios.post<ThreadDetailsResponse>(
        `${process.env.NEXT_PUBLIC_BE_URL}/user/${username}/threads`,
        {
            name: name
        },
        config
    );
    return res;
}
