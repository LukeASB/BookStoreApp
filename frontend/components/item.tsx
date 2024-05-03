import IItem from "../shared/interfaces/IItem";

// This will be called when the item is clicked.
const Item: React.FC<{ item: IItem}> = ( { item }: { item: IItem}) => {
    if (item?.genres) {
        item.genres = [item.genres.join(", ")];
    }

    return (
        <div className="book-details">
            <ul>
                <li><strong>Title:</strong> {item.title}</li>
                <li><strong>Published:</strong> {item.published}</li>
                <li><strong>Pages:</strong> {item.pages}</li>
                <li><strong>Genres:</strong> {item.genres} </li>
                <li><strong>Rating:</strong> {item.rating}</li>
            </ul>
        </div>
    );
};

export default Item;