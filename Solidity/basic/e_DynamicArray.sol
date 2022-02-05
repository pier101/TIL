// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;

contract DynamicArrays{
    // # uint 유형의 동적 배열
    uint[] public numbers;
    
    // # 길이 반환
    function getLength() public view returns(uint){
        return numbers.length;
    }
    
    // # 새로운 요소 추가
    function addElement(uint item) public{
        numbers.push(item);
    }
    
    // # 인덱스에 있는 요소 반환
    function getElement(uint i) public view returns(uint){
        if(i < numbers.length){
            return numbers[i];
        }
        return 0;
    }
    
    
    // # 배열의 마지막 요소 제거
    function popElement() public{
        numbers.pop();
    }
    
    function f() public{
    // # 메모리 동적 배열 선언
    // 메모리 배열의 크기를 조정할 수 없습니다(push() 및 pop()는 사용할 수 없음).
        uint[] memory y = new uint[](3);
        y[0] = 10;
        y[1] = 20;
        y[2] = 30;
        numbers = y;
    }
 
}