import React, {useState} from 'react';
import "./ReviewForm.css"
import FileInput from "./FileInput";

const ReviewForm = () => {
    const [values, setValues]= useState({
        title: '',
        rating: 0,
        content: '',
        imgFile: null,
    })

    // todo: check
    function sanitize(type, value) {
        switch (type) {
            case 'number':
                return Number(value) || 0;

            default:
                return value;
        }
    }

    const handleChange = (name, value) => {
        setValues((prevState) => ({
            ...prevState,
            [name]: value
        }))
    }

    const handleInputChange = (e) => {
        const {name, value, type}= e.target;
        handleChange(name, sanitize(type, value));
    }

    const handleSubmit = (e) => {
        e.preventDefault();
    }

    return (
        <form className="ReviewForm" onSubmit={handleSubmit}>
            <FileInput name="imgFile" value={values.imgFile} onChange={handleChange}/>
            <input name="title" value={values.title} onChange={handleInputChange}/>
            <input type="number" name="rating" value={values.rating} onChange={handleInputChange}/>
            <textarea name="content" value={values.content} onChange={handleInputChange}/>
            <button type="submit">확인</button>
        </form>
    );
};

export default ReviewForm;