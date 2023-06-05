import axios from "axios";

type OneTimeData = {
    value: string
}

type OneTime = {
    data: OneTimeData,
    success: boolean
}

export const getOneTime = async (onetime:string) => {
    const res = await axios.get<OneTime>(
       'https://data.tribist.com/onetime?onetime='+onetime
    );
    return res;
}
