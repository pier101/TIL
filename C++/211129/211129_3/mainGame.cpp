#include "mainGame.h"


mainGame2::mainGame2()
{
	_elfWarrior = new darkElfWarrior;
	_elfWarrior->Qskill();
}

mainGame2::~mainGame2()

{
	//�����Ҵ����� delete�� �޸𸮻��� ������
	delete _elfWarrior;
}
