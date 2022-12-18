import {useEffect, useRef, useState} from "react";

/*
    file input은 반드시 비제어로 만들어야 함
    해커가 js조작을 통해 사용자 파일을 전송하는 것을 막기위해
    value props사용시 사용자 path 숨겨줌

*/

function FileInput( {name, value, onChange }) {
    const inputRef = useRef();

    useEffect(()=>{
        if(inputRef.current)
        console.log(inputRef.current);
    })
    const handleChange = (e) => {
        const nextValue = e.target.files[0];
        onChange(name, nextValue);
    }
    return <input type="file" onChange={handleChange} ref={inputRef}/>
}

export default FileInput;