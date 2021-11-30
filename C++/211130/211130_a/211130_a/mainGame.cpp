#include "mainGame.h"

mainGame::mainGame()
{
	_hp = 100;
	_mp = 50;
}

mainGame::mainGame(int hp)
{
	_hp = hp;
	_mp = 50;
	
	int num = 3;	//C언어 대입연산자
	int num2(5); //C++ 괄호로 변수 초기화 가능
}
// : _hp(hp), _mp(mp) -> 멤버이니셜라이즈 : 원래는 const 인자를 초기화하기 위한 수단인데
// 편하다보니 변수 초기화할때 더 많이 씀
mainGame::mainGame(int hp, int mp)
{
	
}

mainGame::~mainGame()
{
}
