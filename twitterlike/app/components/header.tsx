import { useState } from "react";
import { FiCloudRain } from "react-icons/fi";

export default function Header() {
    var [showVMenu, setShowVMenu] = useState(false)
    return (
        <div className="fixed bg-black h-[50px] md:p-[5px] mb-[5px] w-full md:items-center">
            <div className="hidden md:block w-full">
                <button className="text-white float-right p-[5px] mr-[50px]">Signup</button>
                <button className="text-white float-right p-[5px] mr-[50px]">Login</button>
                <button className="text-white float-right p-[5px] mr-[50px]">About</button>
            </div>
            <div className="flex flex-col overflow-y-visible md:hidden ">
            <div className="h-[50px] flex items-center">
                <div className="mx-[15px] text-white">
            <FiCloudRain size={30} onClick={()=>{setShowVMenu(!showVMenu)}}/>
                </div>
            </div>
            {((show: Boolean)=> {
                if (show) {
                    return (
                <div className="bg-black text-white max-w-[70%]">
                    <ul className="">
                            <li className="bg-black hover:bg-slate-500 w-full pl-[5px] py-[5px] block">Signup</li>
                            <li className="bg-black hover:bg-slate-500 w-full pl-[5px] py-[5px] block">Login</li>
                            <li className="bg-black hover:bg-slate-500 w-full pl-[5px] py-[5px] block">About</li>
                    </ul>
                </div>
                    )
                }
            })(showVMenu)}
            </div>
        </div>
    );
}
