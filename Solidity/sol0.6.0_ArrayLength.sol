pragma solidity >=0.5.15 <=0.6.0;



/*
  # works in 0.5.15

contract MyContract {
    uint[] public myUintArray;
    
    function add(uint _num) public {
        myUintArray.length++;
        myUintArray[myUintArray.length - 1] = _num;
    }
    
    function removeElement() public {
        myUintArray.length--;
    }
}
 */

/*
  #works in 0.5.15 AND 0.6.0
*/
contract MyContract {
    uint[] public myUintArray;
    
    function add(uint _num) public {
        myUintArray.push(_num);
    }
    
    function removeElement() public {
        myUintArray.pop();
    }
}

