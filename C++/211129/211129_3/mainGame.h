#pragma once
#include "darkElfWarrior.h"
class mainGame2
{
private:
	//동적할당을 하기 위해서 객체를 선언만 해놓은 상태 (메모리는 null값이다.)
	darkElfWarrior* _elfWarrior;

public:
	mainGame2();
	~mainGame2();
};

