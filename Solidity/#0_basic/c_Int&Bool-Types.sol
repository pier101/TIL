// SPDX-License-Identifier: GPL-3.0

// pragma solidity >=0.5.0 <0.9.0;
pragma solidity ^0.8.0;

contract Property{
    // # 1. Boolean Type
    bool public sold;

    // # 2. Int Type
    uint8 public x = 255;
    int8 public y = -10;

    function f1() public {
        unchecked {x += 1;}
    }

}
