/*
    break point 잡기
    debug
    F5 : break point 시점만 확인
    F10 : break point 시점과 관련된 단계 확인 (but,함수 내부까진 안 들어감)
    F11 : break point 시점과 관련된 단계 확인 (but,함수 내부까지 들어감)
    shift + F11 : 함수 내부로 들어왔을 경우 나갈때

    # 조사식 : 원하는 특정값 or 특정식을 넣어주면 그 값은 고정적으로 볼 수 있다.
    # 호출 스택 : 시점별 스택 정보 보여줌
    # 중단점 편집 : 
        활용ex) for loop 같은 식에서 원하는 결과일 때만 break 걸기 (ex i===3)

    # launch.json 활용
        활용 방안) 디버깅 중 코드 편집때 마다 자동 디버깅 되게!
        1. "launch.json 파일 만들기" 클릭
        2. "program" 밑에 "runtimeExecutable":"nodemon", "restart":true 입력
        * wsl에서 사용시 안된다면 npm i -g nodemon 으로 전역설치
        (권한 이슈 생긴다면 sudo chown -R $(whoami) $(npm config get prefix)/{lib/node_modules,bin,share} 입력
        참조 : https://stackoverflow.com/questions/47252451/permission-denied-when-installing-npm-modules-in-osx/47252840 )
        (그래도 안되면 최후 sudo npm i -g nodemon으로 설치)
        * npm 패키지 설치시 sudo로 설치하면 보안에 안전하지 않기 떄문에 최대한 피해야한다는걸 인지하고 있자
*/

function sayHello(){
    console.log('hello~');
    console.log('hellowssssss~');
}

function calculate(x,y){
    console.log('calculating...');
    const result = x + y;
    console.log('result : ',result);
    sayHello();
    return result;
}

calculate(1,2);

const stop = 4;
console.log('..... >> loopings >>.....');
for (let i = 0; i < 10; i++) {
    console.log('count', i)
    if (i === stop) {
        break;
    }
    
}

