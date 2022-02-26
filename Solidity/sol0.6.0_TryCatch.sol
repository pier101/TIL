pragma solidity ^0.6.0;

contract ContractA {
    function funARequireFailure() public pure {
        require(false, "This is an error String");
    }
    
    function funBRevertFailure() public pure {
        revert("Error from Contract A");
    }
    
    function funCAssertFailure() public pure {
        assert(false);
    }
}

contract B {
    ContractA instA;
    
    event Error(string _reason);
    event LowLevelError(bytes _reason);
    
    constructor() public {
        instA = new ContractA();
    }
    
    function testRequireTryCatch() public returns(bool) {
        try instA.funCAssertFailure() {
            return true;
        } catch Error(string memory reason) {
            // 이 부분 실행되는 경우
            // getData 내부에서 revert명령문 호출될 때
            // reason이 string type으로 제공 
            emit Error(reason);
            return false;
        } catch (bytes memory lowLevelData) {
            // revert()함수가 사용될 떄 이 부분 실행
            // or there was a failing assertion, division
            // by zero, etc. inside getData.
            emit LowLevelError(lowLevelData);
            return false;
        }
    }
}
