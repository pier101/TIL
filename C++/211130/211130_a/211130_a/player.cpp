#include "player.h"

player::player()
{
	_hp = 100;
	_atk = 10;
}

player::~player()
{
}

void player::printPlayer()
{
	cout << "�÷��̼� ü�� : " << _hp <<endl;
	cout << "�÷��̾���ݷ�: " << _atk << endl;
}
