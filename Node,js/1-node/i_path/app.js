const path = require('path');

console.log(__dirname)
console.log(__filename)

console.log(path.sep);
console.log(path.delimiter);

//basename 
console.log(path.basename(__filename))
console.log(path.basename(__filename,'.js'))

//dirname
console.log(path.dirname(__filename))

//extension
console.log(path.extname(__filename))

//parse
const parsed = path.parse(__filename);
console.log(parsed);

const str = path.format(parsed);
console.log(str)

//isAbsolute (절대경로? 인지 여부)
console.log('isAbsolute?', path.isAbsolute(__dirname)); //절대경로
console.log('isAbsolute?', path.isAbsolute('../')); //상대경로

// normalize (경로 올바르게 수정)
console.log(path.normalize('./folder///sub'))

// join
console.log(__dirname+path.sep+"image");
console.log(path.join(__dirname,'image'))// 더 간편히 만들 수 있음

