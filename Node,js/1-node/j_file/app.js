const fs = require('fs');

/*
    # 모든 api는 3가지 형태로 제공된다
        rename(...(코드들 실행되고나서) , callback(error, data))    콜백함수 전달
        try {renameSync(....)} catch(e){ }        콜백함수 전달 안함 
        promises.rename().then().catch(0)
*/


//-----------동기적 수행 ------------------------
// fs.renameSync('./file.txt','./file-new.txt') //파일 없어 에러 발생
// console.log('hello')

//----------------------------------------------
// >> try로 감싸주자
try {
    fs.renameSync('./file.txt','./file-new.txt') ;
} catch(err){
    console.error(err)
}
console.log('hello') //에러 발생하더라도 찍힌다.


//---------비동기적 수행하는 것들 ------------------
fs.rename('./file-new.txt','./file.txt',error=>console.error(error));
console.log("hello2")

fs.promises.rename('./file2.txt','file2-new.txt')
.then(()=> console.log('done'))
.catch(console.error)
console.log("hello3")
//----------------------------------------------