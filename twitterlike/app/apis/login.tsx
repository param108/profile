import axios, { AxiosRequestConfig } from "axios";
import "./interceptors";

type OneTimeData = {
    value: string
}

type OneTime = {
    data: OneTimeData,
    success: boolean
}

type profileData = {
    user_id: string,
    username: string,
    profile: string
}

type Profile = {
    data: profileData,
    success: boolean
}

function getConfig(cfg: any): AxiosRequestConfig {
    return cfg
}

export const getOneTime = async (onetime:string) => {
    const config = {
        retry: 3
    }
    const res = await axios.get<OneTime>(
       `${process.env.NEXT_PUBLIC_BE_URL}/onetime?id=${onetime}`,
        getConfig(config)
    );
    return res;
}

export const getProfile = async (token:string) => {
    const config = {
        headers:{
            "TRIBIST_JWT": token,
        },
        retry: 3
    };

    const res = await axios.get<Profile>(
       `${process.env.NEXT_PUBLIC_BE_URL}/profile`,
        config
    );
    return res;
}
