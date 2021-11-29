#pragma once


using namespace std;
class mainGame1
{
public:

	//그냥 아이템 사용하는 함수
	void UseItem();
	//코인의 가스를 소모하는 함수
	void useItem(int gas);

	//오버로딩 : 객체 내에서 같은 이름 함수를 형태가 다르게 재정의 하는것
	//오버라이드 : 상속 관계에서 부모클래스에 있는 같은 이름 함수를 
	//				자식 객체에서 재정의 하는 것.
	mainGame1();
	~mainGame1();
};

