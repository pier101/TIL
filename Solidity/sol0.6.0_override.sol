pragma solidity >=0.5.15 <=0.6.0;

/** 
 * Works in Solidity 0.5.15
 * 
 * *

contract A {
    event MyEvent(string _myString);
    function funA() public {
        emit MyEvent("from A");
    }
}

contract B {
    function funA() public {
        //does nothing
    }
}

contract C is B,A {
    function funA() public {
        emit MyEvent("from B");
        super.funA();
    }
}
/**/



/**
 * works in 0.6.0
* */

contract A {
    event MyEvent(string _myString);
    function funA() public virtual {
        emit MyEvent("from A");
    }
}

contract B {
    function funA() public virtual {
        //does nothing
    }
}

contract C is A,B {
    function funA() public override(B,A) {
        emit MyEvent("from B");
        super.funA();
    }
}
/**
 * */