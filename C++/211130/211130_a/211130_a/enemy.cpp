#include "enemy.h"

enemy::enemy()
{
	_hp = 100;
	_atk = 8;
}

enemy::~enemy()
{
}

void enemy::printEnemy()
{
	cout << "���ʹ� ü�� " <<_hp << endl;
	cout << "���ʹ� ���ݷ� :" << _atk  << endl;
}
