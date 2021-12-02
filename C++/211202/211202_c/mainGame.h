#pragma once
#include "exchange.h"
#include "wallet.h"


enum LOCATION
{
	LOCATION_EX = 1, //�ŷ���
	LOCATION_WALLET //����
};
class mainGame
{
private:
	exChange* _ex; //�ŷ��� Ŭ����
	wallet* _wal; //���� Ŭ����
	LOCATION _location;

	int _money; //������
	int _selectNum; //�޴� ���� �������� ���� ����

public:
	mainGame();
		~mainGame();

		void setLocation(int num);
		void setMainPage();
		void setExPage();
		void setWalletPage();
};

