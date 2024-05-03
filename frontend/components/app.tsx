import React from "react";
import Form from "./form";
import MainLayout from "./mainLayout";
import Home from "./home";
import ItemView from "./itemView";
import { useRouter } from "next/router";

interface IApp {
    page?: string;
}

const App: React.FC<IApp> = ({ page = "" }) => {
    const { query: { _id } } = useRouter();

    const renderPage = () => {
        if (page === 'home') {
            return <Home />;
        }

        if (page === 'book/create') {
            return <Form />;
        }

        if (page === 'book/view') {
            if (!_id) return <Home />;
            return <ItemView _id={_id as string}/>
        }

        return (
            <h3>No component for navigation value. {page} not found</h3>
        );
    };

    return (
        <MainLayout>
            {renderPage()}
        </MainLayout>
    );
};

export default App;
