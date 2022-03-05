//SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.0 <0.9.0;

/*
    ㅁ 추상계약
        : 구현되지 않은 기능이 하나 이상 있는 계약
        - 모든 기능이 구현되더라도 계약을 추상적으로 표시할 수 있다.
        - 추상계약은 배포할 수 없다.
        - 다른 계약이 포함된 계약으로만 사용할 수 있다.
        - 기본 계약에서 이미 선언된 상태 변수를 재정의하는 것은 허용되지 않습니다.
        ex) 설계자가 이러이러한 기능들 있어야 된다 표시할 떄 유용

*/
// abstaract 키워드 사용하여 추상계약 선언
abstract contract BaseContract {
    int public x;
    address public owner;

    constructor(){
        x = 5;
        owner = msg.sender;
    }

    // @ 제대로 구현한 기능
    // function setX(int _x) public { 
    //     x = _x;
    // }

    // @ 제대로 구현하지 않은 기능
    function setX(int _x) public virtual; //virtual 키워드 설정
}

contract A is BaseContract{
    int public y = 3;

    // 추상계약에서 제대로 구현하지 않은 기능은 그대로 가져와서 
    // override하여 재정의하거나 현재 계약 또한 abstract설정 하면 된다.
    function setX(int _x) public override {
        x = _x;
    } 
}