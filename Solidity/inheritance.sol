//SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.6.0 <0.9.0;

contract BaseContract {
    int public x;
    address public owner;

    constructor(){
        x = 5;
        owner = msg.sender;
    }

    function setX(int _x) public {
        x = _x;
    }
}
/*
 ㅁ 상속
  - BaseContract에서 파생된 계약
  - BaseContract가 A에 복사됨을 의미
  - 생성자 또한 암시적으로 호출됨
  - 자체 기능 추가가 가능하다
  - 새로운 단일 인스턴스가 생성된다고 생각하면 된다.(기본 계약은 배포되지 않았다.)
*/
contract A is BaseContract{
    int public y = 3;

}