//SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.0 <0.9.0;

/*
    ㅁ 인터페이스
        : 추상 계약과 유사하지만 기능이 없다.
        - 다른 계약에서는 상속할 수 없지만 다른 인터페이스에서는 상속할 수 있다.
        - 선언된 함수는 external 함수여야 하며, 생성자를 선언할 수 없다.
        - 상태변수 또한 선언 불가
*/

// interface 선언
interface BaseContract {
    function setX(int _x) external; //external 키워드 설정
}

contract A is BaseContract{
    // 상태변수 선언해준다.
    int public x; 
    int public y = 3;

    // interface의 함수 재정의 방법
    function setX(int _x) public override { //public으로 변경가능(private/internal 불가)
        x = _x;
    } 
}
