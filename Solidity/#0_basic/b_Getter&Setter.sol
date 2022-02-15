// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;

contract Property{

     // # 컨트랙트 저장소에 저장된 상태 변수 선언
    int public price; // 기본적으로 비공개 
    string public location = "London";

    // immutable 선언 시 생성자에서만 초기화할 수 있다.
    address immutable public owner;

    // # 상수 선언
    int immutable area = 100;

    // # 생성자 선언 (컨트랙트 배포 시 한 번만 실행된다.)
    constructor(int _price, string memory _location){
        price = _price;
        location = _location;
        owner = msg.sender; // 계약을 배포하는 계정 주소로 초기화  
    }

    // # setter 함수, (상태 변수 설정) 
    // 트랜잭션 생성 -> 가스 비용 발생
    function setPrice(int _price) public{
         int a; // 스택에 저장된 지역 변수 
        a = 10;
        price = _price;
        price = _price;
    }
    
    // # getter 함수, (상태 변수 반환)
    // 'view'로 선언된 함수는 블록체인을 변경하지 않음.
    function getPrice() public view returns(int){
        return price;
    }

    function setLocation(string memory _location) public{ //문자열 유형은 메모리 또는 저장소로 선언되어야 한다.
        location = _location;
    }
}