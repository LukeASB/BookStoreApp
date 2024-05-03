import Link from "next/link";
import IItem from "../shared/interfaces/IItem";


const ItemRow: React.FC<{ item: IItem}> = ({ item }: { item: IItem}) => {
    return (
        <tr className={item._id}>
            <td><Link href={{pathname: "/book/view", query: { _id: item._id } }}>{item.title}</Link></td>
            <td>{item.pages}</td>
            <td>{item.published} </td>
            <td>{item.rating}</td>
        </tr>
    );
};

export default ItemRow;