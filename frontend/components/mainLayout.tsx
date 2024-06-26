import { ReactNode } from "react";
import Banner from "./banner";
import Navbar from "./navbar";
import Link from "next/link";

interface IMainLayout {
    children?: ReactNode
}

const MainLayout: React.FC<IMainLayout> = ( { children } ) => {
    return (
        <>
        <header>
            <Banner><h1><Link href="/">Reading List</Link></h1></Banner>
        </header>
        <Navbar />
        <main>
            { children }
        </main>
        </>
    );
};

export default MainLayout;