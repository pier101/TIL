console.log('logging...');
console.clear();

//log level
console.log('log'); // 개발
console.info('info'); // 정보
console.warn('warn') // 경고
console.error('error') // 에러, 사용자 에러, 시스템 에러

//assert
//첫 번째 인자가 참이 아닌 경우에만 출력
console.assert(2===3, 'not same!');
console.assert(2===2, "same!");

//print object
const student = {name: 'ellie',age:20, company:{name:'AC'}};
console.log(student);
console.table(student);
console.dir(student,{showHidden:true, color: false, depth:0});

// measuring time
// 성능 측정 시
console.time('for loop');
for (let i = 0; i < 10; i++) {
    i++;
}
console.timeEnd('for loop');

// counting
// 함수가 몇 번 호출되었는지
function a(){
    console.count('a function');
}
a();
console.countReset('a function'); // 카운트 초기화
a();

// trace
// 호출 추적 가능, 디버깅 시 유용하다.
function f1(){
    f2();
}
function f2(){
    f3();
}
function f3(){
    console.log('f3');
    console.trace();
}
f1();