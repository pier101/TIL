const fs = require('fs').promises
const path = require('path');
const process = require('process');

// console.log(path.dirname(__filename))
// console.log(process.argv[1])

const currentFolder = path.dirname(__filename)

console.log("-------")
fs.readdir(currentFolder)
.then(fileList=>{
    fileList.forEach(file=>{
        if(path.extname(file) === 'mp4' || path.extname(file) === 'mov'){
            
        }
    })
})
.catch(console.error)

// fs.mkdir('video')
// .then()
// .catch(console.error)

// fs.mkdir('captured')
// .then()
// .catch(console.error)

// fs.mkdir('duplicated')
// .then()
// .catch(console.error)

/*
    실행

    1)
    파일 확장자들 읽음 
        mp4,mov = video 폴더로
        aae,png,jpg = captured 폴더로
        폴더 존재?*
            없다 => 폴더 생성*
            있다 =>  

*/