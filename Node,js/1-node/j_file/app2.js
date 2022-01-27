const fs= require('fs').promises;


// read a file
fs.readFile('./file.txt','utf-8')
.then(data=>console.log(data))
.catch(console.error)

// wriging a file
fs.writeFile('./file3.txt','hello dongwook!')
.catch(console.error);

fs.appendFile('./file3.txt','bye dongwook!') //내용 이어지게
.then(()=>fs.copyFile('./file3.txt','./file4.txt').catch(console.error))
.catch(console.error);

// copy
// fs.copyFile('./file3.txt','./file4.txt')
// .catch(console.error)

// folder
fs.mkdir('sub-folder')
.catch(console.error);

fs.readdir('./')
.then(console.log)
.catch(console.error);