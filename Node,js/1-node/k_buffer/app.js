/*
    # 버퍼
        고정된 사이즈의 메모리 덩어ㅣ
        숫자의 배열 형태
        데이터에 있는 byte 그 자체를 가르킨다.
*/

const buf = Buffer.from('Hi');
console.log(buf); //<Buffer 48 69> 유니코드 번호로 나온다.
console.log(buf.length); //2
console.log(buf[0]); // 여기선 아스키코드  번호로 나온다.
console.log(buf[1]);
console.log(buf.toString());

// create
const buf2 = Buffer.alloc(2); // >> 사이즈가 2개인 버퍼를 만들어라!
const buf3 = Buffer.allocUnsafe(2); // >> 메모리 초기화 안 하고 바로 보여줘! (더 빠름,메모리 공간 없을시 재대로 출력 안될 수 있음)
buf2[0] = 72;
buf2[1] = 105;
buf2.copy(buf3);

console.log(buf2.toString())
console.log(buf3)

// concat
const  newBuf = Buffer.concat([buf,buf2,buf3])
console.log(newBuf.toString())