/*
    # 계획    
    1. 사용자가 원하는 폴더의 이름을 받아온다.
    2. 그 폴더안에 video, captured, duplicated 폴더를 만든다
    3. 폴더 안에 있는 파일들을 다 돌면서 해당하는 mp4|mov, png|aae, IMG_1234(IMG_E1234)
*/
const path = require('path');
const folder = process.argv[2];
if (!folder) {
    console.error('픽쳐 안에 있는 폴더 이름을 입력해라')
}

const workingDir = path.join(os.homedir(), 'Pictures',filder);
console.log(workingDir);
console.log(process.argv[2]);