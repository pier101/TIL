/*
    ### Number obj
        - 숫자 처리를 위한 obj
        - 숫자 처리를 위한 함수와 프로퍼티가 포함되어 있음
        - 함수를 호출하여 숫자 처리를 하게 됩니다.

    ### property list
        - new Number() : 인스턴스 생성
        - Number 함수
            - Numbdr() : 숫자 값으로 변환하여 반환
        - Number.prototype
            - constructor : 생성자
            - toString() : 숫자 값을 문자열로 변환
            - toLocaleString() : 숫자 값을 지역화 문자로 변환
            - ValueOf() : 프리미티브 값 반환(인스턴스에 설정된 값)
            - toExponential() : 지수 표기로 변환
            - toFixed() : 고정 소숫점 표기로 변환
            - toPrecision() : 고정 소숫점 또는 지수 표기로 변환
*/
let a = 123
console.log(a.toExponential())
console.log(typeof a.toString())
console.log(a.valueOf())
console.log(a.toFixed(2))

/*
    ** 인스턴스 생성시 prototype 안의 property만 가져옴

    #### 프리미티브 값
        - 언어에 있어 가장 낮은 단계의 값
        - ex) var book = "책"  ==> "책"은 더 이상 분해, 전개 불가
        
*/