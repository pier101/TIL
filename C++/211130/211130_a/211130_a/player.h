#pragma once
#include <iostream>

using namespace std;
class player
{
private:
	int _hp;
	int _atk;

public:
	player();
	~player();

	//�÷��̾�Ŭ������ ������Ȳ ����Լ�
	void printPlayer();
	//�����ڴ� ������ ���ٸ� �ϵ��� ���ش�.
	int getPlayerHP() { return _hp; }
	//������ 
	void setplayerHP(int hp) { _hp = hp; }

	//���ݷ¿� ���� ������
	int getPlayerAtk() {return _atk;}
};

