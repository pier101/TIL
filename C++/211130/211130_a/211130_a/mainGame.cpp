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
	
	int num = 3;	//C��� ���Կ�����
	int num2(5); //C++ ��ȣ�� ���� �ʱ�ȭ ����
}
// : _hp(hp), _mp(mp) -> ����̴ϼȶ����� : ������ const ���ڸ� �ʱ�ȭ�ϱ� ���� �����ε�
// ���ϴٺ��� ���� �ʱ�ȭ�Ҷ� �� ���� ��
mainGame::mainGame(int hp, int mp)
{
	
}

mainGame::~mainGame()
{
}
