import axios from "axios";

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

export const getOneTime = async (onetime:string) => {
    const res = await axios.get<OneTime>(
       'https://data.tribist.com/onetime?id='+onetime
    );
    return res;
}

export const getProfile = async (token:string) => {
    const config = {
        headers:{
            "TRIBIST_JWT": token,
        }
    };

    const res = await axios.get<Profile>(
       'https://data.tribist.com/profile',
        config
    );
    return res;
}
