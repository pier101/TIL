const { count } = require('console');
const fs = require('fs');
const data = [];

const readStream = fs.createReadStream('./file.txt',{
    // highWaterMark: 8, // 기본 = 64 kbytes
    // encoding: 'utf-8'
});

readStream.on('data',(chunk)=>{
    // console.log(chunk);
    data.push(chunk)
    console.count('data');
});


readStream.on('end',()=>{
    console.log(data.join(''))
})
readStream.on('error',(error)=>{
    console.log(error);
})

//const readStream 변수 없애고 체이닝도 가능하다.