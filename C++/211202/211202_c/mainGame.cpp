#include "mainGame.h"

mainGame::mainGame()
{
	//�ŷ��� Ŭ���� ���� �Ҵ� �� ���� ����
	_ex = new exChange;
	_ex->coinSetup();

	_wal = new wallet;

	//�ܼ� ���� �ʱ�ȭ
	_money = 1000000;
	_selectNum = 0;

	//��ó�� ����� ����ȭ�� �����Բ� ����
	setMainPage();
}

mainGame::~mainGame()
{
	delete _ex;
	delete _wal;
}

void mainGame::setLocation(int num)
{
	switch (num)
	{
	case LOCATION_EX:

		setExPage();

		break;
	case LOCATION_WALLET:

		setWalletPage();

		break;
	default:
		setMainPage();
		break;
	}
}

void mainGame::setMainPage()
{
	cout << "============= ��� 3�� ==============" << endl;
	cout << "1. �ŷ���          2. �� ����" << endl;
	cout << "====================================" << endl;

	cin >> _selectNum;
	system("cls");	//<- ȭ�� �����ϰ� �о��ִ� ��ɾ�
	setLocation(_selectNum);
}

void mainGame::setExPage()
{
	//�ŷ��� �ȿ����� ����ؼ� � ������ ���״ϱ� while������..
	while (true)
	{
		cout << "============== ���3�� �ŷ��� ==========================" << endl;
		cout << "1. ��Ʈ���ΰ迭, 2.�̴�����迭, 3.���ð迭, 4.��Ʈ����" << endl;
		cout << "=======================================================" << endl;
		cout << "����ȭ������ ������ 0�� ��" << endl;

		cin >> _selectNum;
		system("cls");

		if (_selectNum == 0) //����ȭ������ ������ �Ѵٸ�
		{
			//������������ ȣ�����ְ� �극��ũ�� ���Ϲ� Ż��
			setMainPage();
			break;
		}

		_ex->coinOutput(_selectNum);

		cout << "�ٸ� ������ ������ �ƹ�Ű�� ��������" << endl;
		cin >> _selectNum;
		system("cls");
	}
}

void mainGame::setWalletPage()
{
	cout << "================== �� �� �� ====================" << endl;
	_ex->getCoinVector();
}