import { useEffect, useState } from "react";
import IItem from "../shared/interfaces/IItem";
import Item from "./item";

interface IItemView {
    _id: string
}

const ItemView: React.FC<IItemView> = ({ _id }: { _id: string }) => {
    const initialState: IItem = {
        _id: "",
        title: "",
        published: 0,
        pages: 0,
        genres: [],
        rating: 0,
    }
    const [ item, setItem ] = useState(initialState);
    useEffect(() => {
        const getItem = async () => {
            try {
                const res = await fetch(`http://localhost/v1/books/${_id}`);
                if (!res.ok) throw new Error("Network response was not ok");
                const data = await res.json();
                if (!data) throw new Error("No data returned in the response.");
                console.log("Data:", data.book)
                setItem(data.book);
            } catch (err) {
                console.log("Problem occured when calling /v1/books:", err);
                setItem(initialState);
            }
        };

        getItem();
    }, []);

    return (
        <div className="bookView">
            <Item item={item}/>
        </div>
    );
};

export default ItemView;