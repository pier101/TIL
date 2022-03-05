//SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.6.0 <0.9.0;

contract CrowdFunding{
    
    // == 사용자 관련 ==
    mapping(address => uint) public contributors;// 기부자 => 기부금액
    address public admin; // 관리자
    uint public  numOfContributors; // 기부자 수

    // == 기부 캠페인 관련 ==
    uint public minimumContribution; // 인당 최소 기부금
    uint public deadline; // 마감기한 timestamp
    uint public goal; // 기부 목적
    uint public raisedAmount; // 현재까지 모금금액

    // == 지출 요청(기부자는 투표) 관련 ==
    struct Request{
        string description; 
        address payable recipient; // 기부자 주소
        uint value; 
        bool completed; // 완료 여부
        uint numOfVoters; // 투표자 수
        mapping(address => bool) voters; // 투표 여부
    }
    // (관리자는 둘 이상의 지출 요청을 생성할 수 있다.)
    // 지출 요청 목록
    mapping(uint=>Request) public request;
    uint public numRequests;

    event ContributeEvent(address _sender, uint _value);
    event CreateRequestEvent(string _description, address _recipient, uint _value);
    event MakePaymentEvent(address _recipient, uint _value);


    // 캠페인 목표, 기한을 설정하는 생성자
    constructor(uint _goal, uint _deadline)  {
        goal = _goal;
        deadline = block.timestamp + _deadline;
        minimumContribution = 100 wei;
        admin = msg.sender;
    }

    // @ 기부 기능
    function contribute() public payable{
        // 마감기한 전에만 기부 가능
        require(block.timestamp < deadline,"Deadline has passed!");
        require(msg.value >= minimumContribution, "Minimum Contribution not met");

        if(contributors[msg.sender] == 0) { // 기부자가 처음 기부할 경우
            numOfContributors++; //기부수 증가
        }
        contributors[msg.sender] += msg.value;
        raisedAmount += msg.value;

         emit ContributeEvent(msg.sender, msg.value);
    }

    receive() external payable {
        contribute(); // 명시해줘야 calldata호출시 contract balance에 반영됨
    }

    function getBalance() public view returns(uint){
        return address(this).balance;
    }

    // @ 환불 기능
    function getRefund() public {
        // 마감기한 지났거나 목표금액 도달 못했을 경우
        require(block.timestamp > deadline && raisedAmount < goal); 
        // 기부자일 경우에만
        require(contributors[msg.sender] > 0);

        // 기부자에게 환불
        address payable recipient = payable(msg.sender);
        uint value = contributors[msg.sender];
        recipient.transfer(value);
        // payable(msg.sender).transfer(contributors[msg.sender]); 랑 동일

        contributors[msg.sender] = 0;
    }

    modifier onlyAdmin(){
        require(msg.sender == admin,"Only admin can call this function!");
        _;
    }

    //@지출요청 생성 함수
    function createRequest(string memory _description, address payable _recipient, uint _value) public onlyAdmin {
        // * 매핑의 값이 앞에 나오고 키가 뒤에 나오는 형식도 됨 **
        Request storage newRequest = request[numRequests];
        numRequests++;

        newRequest.description = _description;
        newRequest.recipient = _recipient;
        newRequest.value = _value;
        newRequest.completed = false;
        newRequest.numOfVoters = 0;  

         emit CreateRequestEvent(_description, _recipient, _value);
    }

    // @지출요청에 대한 투표 함수
    function voteRequest(uint _requestNo) public{
        // 기부자만 투표 참여 가능
        require(contributors[msg.sender] > 0, "You must be a contributor to vote!");
        Request storage thisRequest = request[_requestNo];

        require(thisRequest.voters[msg.sender] == false,"You are already voted!");
        thisRequest.voters[msg.sender] = true;
        thisRequest.numOfVoters++;
    }

    // @ 과반 투표 결과 처리
    function makePayment(uint _requestNo) public onlyAdmin {
        require(raisedAmount >= goal);
        Request storage thisRequest = request[_requestNo];

        require(thisRequest.completed == false, "The request has been already completed");
        require(thisRequest.numOfVoters > numOfContributors / 2); // 50% 이상 투표시 통과

        thisRequest.recipient.transfer(thisRequest.value);
        thisRequest.completed = true;

         emit MakePaymentEvent(thisRequest.recipient, thisRequest.value);
    }
}

/*
    멤버변수 설정
    생성자 설정
    @ 기부 함수
    이더 수신 함수
    @ 컨트랙트 balance 조회 함수
    ---- test ----
    @ 환불 함수
    ---- test ----
    지출요청 변수 설정
    onlyAdmin 함수 수정자 생성
    @지출요청 함수 생성
    ---- test ----
    @지출요청에 대한 투표 함수
    ---- test ----
    @과반 투표 결과 처리
    ---- test ----
*/