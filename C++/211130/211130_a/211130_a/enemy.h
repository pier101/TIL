#pragma once
#include <iostream>

using namespace std;
class enemy
{
private:
	int _hp;
	int _atk;

public:
	enemy();
	~enemy();

	//���ʹ� ���� ������Ȳ �����ִ� �Լ�
	void printEnemy();

	int getEnemyHP() { return _hp; }
	void setEnemyHP(int hp){ _hp = hp; }

	int getEnemyAtk() { return _atk; }
};

