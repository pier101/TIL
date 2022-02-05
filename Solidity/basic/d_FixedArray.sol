// SPDX-License-Identifier: GPL-3.0

// pragma solidity >=0.5.0 <0.9.0;
pragma solidity ^0.8.0;

contract FixedSizeArray{
    // # 3개의 요소가 있는 uint 유형의 고정 크기 배열 선언
    uint[3] public number = [2,4,5];

    bytes1 public b1;
    bytes2 public b2;
    bytes3 public b3;
    //.. 최대 bytes32

    // # 특정 인덱스의 배열 요소 변경
    function setElement(uint index, uint value) public{
        number[index] = value;
    }

    // # 배열의 원소 개수 반환
    function getLength() public view returns(uint){
        return number.length;
    }

    // # 바이트 배열 설정
    function setBytesArray() public {
        b1 = 'a';  // => 0x61 (ASCII code of `a` in hex)
        b2 = 'ab';  // => 0x6162
        b3 = 'abc'; // => 0x616263
        //b3 = 'z'; // => 0x7a0000
        //b3[0] =  'a' // => 오류 - 단일바이트 수정불가
        
        
    }// bytes는 이전 (버전)코드의 bytes1에 대한 별칭.
}