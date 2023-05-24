export default function Header() {
    return (
        <div className="fixed flex bg-black h-[50px] p-[5px] mb-[5px] w-full items-center justify-end">
            <div>
                <button className="text-white float-right p-[5px] mr-[50px]">Signup</button>
                <button className="text-white float-right p-[5px] mr-[50px]">Login</button>
                <button className="text-white float-right p-[5px] mr-[50px]">About</button>
            </div>
        </div>
    );
}
