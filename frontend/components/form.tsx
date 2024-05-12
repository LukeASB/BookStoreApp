import { useRouter } from "next/router";
import { ChangeEvent, FormEvent, useState } from "react";

interface Book {
    title: string;
    pages: number;
    published: number;
    genres: string;
    rating: number;
    [key: string]: string | number;
}

interface BookErr {
    title: string;
    pages: string;
    published: string;
    genres: string;
    rating: string;
    [key: string]: string
}

const Form = () => {
    const form: Book = {
        title: "",
        pages: 0, 
        published: 0,
        genres: "",
        rating: 0,
    }

    const formErr: BookErr = {
        title: "",
        pages: "", 
        published: "",
        genres: "",
        rating: "",
    }

    const router = useRouter();

    const [formData, setFormData] = useState(form)

    const [errors, setErrors] = useState(formErr);

    const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;

        const parsedValue = () => {
            if (name === 'pages') {
                return parseInt(value);
            }

            if (name === 'published') {
                return parseInt(value);
            }

            if (name === 'genres') {
                return [value];
            }

            if (name === 'rating') {
                return parseInt(value);
            }

            return value;
        }

        setFormData({
            ...formData,
            [name]: parsedValue()
        } as Book);
    }

    const validateForm = () => {
        let isValid = true;
        const newErrors = { ...errors };

        const isRequired = () => {
            const required = "Required"
            for (let item of Object.keys(formData)) {
                const formItem = formData[item];
                console.log(typeof formItem)
                if (typeof formItem === 'string') {
                    console.log(formItem);
                    !formItem.trim() ? newErrors[item] = required : newErrors[item] = "";
                }
                if (typeof formItem === 'number') {
                    console.log(formItem);
                    !formItem ? newErrors[item] = required : newErrors[item] = "";
                }
                if (typeof formItem === 'object') {
                    console.log(formItem);
                    !formItem[0] ? newErrors[item] = required : newErrors[item] = "";
                }
            }
        };
        
        isRequired();
        console.log();
        if (Object.values(newErrors).filter(el => el.length > 0).length > 0) {
             isValid = false;
        }
        setErrors(newErrors);

        return isValid;
    }

    const postForm = async (data = {}) => {
        try {
            await fetch("http://localhost/v1/books", {
                method: "POST",
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });
        } catch (err) {
            console.log("Error occured when posting the form:", err);
        }
    }

    const handleSubmit  = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        if (validateForm()) {
            postForm(formData);
            router.push("/");
        }
    }

    return (
        <form onSubmit={handleSubmit} method='Post'>
            <div className="row mb-3">
                <label htmlFor="title" className="col-sm-3 col-form-label">Title:</label>
                <div className="col-sm-9">
                    <input type="text" className="form-control" id="title" name="title" value={formData.title} onChange={handleInputChange} />
                    <div className="error text-danger">{errors.title}</div>
                </div>
            </div>
            <div className="row mb-3">
                <label htmlFor="pages" className="col-sm-3 col-form-label">Pages:</label>
                <div className="col-sm-9">
                    <input type="number" className="form-control" id="pages" name="pages" value={formData.pages ? formData.pages : ""} onChange={handleInputChange} />
                    <div className="error text-danger">{errors.pages}</div>
                </div>
            </div>
            <div className="row mb-3">
                <label htmlFor="published" className="col-sm-3 col-form-label">Published:</label>
                <div className="col-sm-9">
                    <input type="number" className="form-control" id="published" name="published" value={formData.published ? formData.published : ""} onChange={handleInputChange} />
                    <div className="error text-danger">{errors.published}</div>
                </div>
            </div>
            <div className="row mb-3">
                <label htmlFor="genres" className="col-sm-3 col-form-label">Genres:</label>
                <div className="col-sm-9">
                    <input type="text" className="form-control" id="genres" name="genres" value={formData.genres} onChange={handleInputChange} />
                    <div className="error text-danger">{errors.genres}</div>
                </div>
            </div>
            <div className="row mb-3">
                <label htmlFor="rating" className="col-sm-3 col-form-label">Rating:</label>
                <div className="col-sm-9">
                    <input type="number" step="0.1" className="form-control" id="rating" name="rating" value={formData.rating ? formData.rating : ""} onChange={handleInputChange} />
                    <div className="error text-danger">{errors.rating}</div>
                </div>
            </div>
            <div className="row">
                <div className="col-sm-10 offset-sm-2">
                    <button type="submit" className="btn btn-primary">Submit</button>
                </div>
            </div>
        </form>
    );
};

export default Form;