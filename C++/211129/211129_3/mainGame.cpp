#include "mainGame.h"


mainGame2::mainGame2()
{
	_elfWarrior = new darkElfWarrior;
	_elfWarrior->Qskill();
}

mainGame2::~mainGame2()

{
	//동적할당했음 delete로 메모리삭제 해주자
	delete _elfWarrior;
}
