import { useEffect, useState } from "react";
import ItemList from "./itemList";
import Pagination from "./pagination";


const Home: React.FC = () => {
    const [ items, setItems ] = useState([]);
    const [ currentPage, setCurrentPage ] = useState(1);
    const [ postPerPage, setPostPerPage] = useState(8);
    const lastPostIndex = currentPage * postPerPage;
    const firstPostIndex = lastPostIndex - postPerPage;
    const currentPosts = items.slice(firstPostIndex, lastPostIndex);
    
    useEffect(() => {
        const getItems = async () => {
            try {
                const res = await fetch("http://localhost/v1/books");
                if (!res.ok) throw new Error("Network response was not ok");
                const data = await res.json();
                if (!data) throw new Error("No data returned in the response.");
                console.log("Data:", data.books)
                setItems(data.books);
            } catch (err) {
                console.log("Problem occured when calling /v1/books:", err);
                setItems([]);
            }
        };

        getItems();
    }, []);

    return (
        <div className="home">
            <ItemList data={currentPosts}/>
            { items.length > 0 && <Pagination totalPosts={items.length} postsPerPage={postPerPage} setCurrentPage={setCurrentPage}/>}
        </div>
    );
};

export default Home;