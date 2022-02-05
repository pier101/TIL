// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;


// # 계약 외부에 구조체 유형 선언
// 이 파일의 모든 계약 선언에서 사용할 수 있다.
struct  Instructor{
    uint age;
    string name;
    address addr;
}

contract Academy{
    // # Instructor 유형의 상태 변수 선언
    Instructor public academyInstructor;

    // # 새로운 enum 타입 선언
    enum State {Open,Closed,Unknown}
    // State 유형의 새 상태 변수 선언 및 초기화
    State public academyState = State.Open;

    // # 생성자에서 구조체 초기화
    constructor(uint _age, string memory _name){
        academyInstructor.age = _age;
        academyInstructor.name = _name;
        academyInstructor.addr = msg.sender;
    }

    // # 구조체 상태 변수 변경
    function changeInstructor(uint _age,string memory _name,address _addr) public{
        if(academyState == State.Open){
            Instructor memory myInstructor = Instructor({
                age: _age,
                name: _name,
                addr: _addr
            });
            academyInstructor= myInstructor;
        }
    }

}

// 구조체는 이 파일에 선언된 모든 계약에서 사용할 수 있다.
contract School{
    Instructor public academyInstructor;
}