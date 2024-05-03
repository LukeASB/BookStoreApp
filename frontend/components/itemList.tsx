import IItem from "../shared/interfaces/IItem";
import ItemRow from "./itemRow";

const ItemList: React.FC<{ data: IItem[]}> = ({ data }: { data: IItem[]}) => {
    return (
        <div className="ItemList">
            <table>
                <thead>
                    <tr>
                        <th>Title</th>
                        <th>Pages</th>
                        <th>Published</th>
                        <th>Rating</th>
                    </tr>
                </thead>
            <tbody>
                {data.map((el: IItem) => (
                    <ItemRow key={el._id} item={el} />
                ))}

            </tbody>
            </table>
            { data.length <= 0 && <p className="text-center">No Data Found!</p>}
        </div>
    );
};

export default ItemList;