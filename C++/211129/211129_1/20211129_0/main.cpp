#include <iostream> //input output stream ����� ���̺귯��
#include "mainGame.h" //Ŭ���� �ҷ��ö� ū ����ǥ
using namespace std; //��Ӿ��°� ���ڴ�

//�׻� �����Լ��� ������Ʈ �� 1���� �ʼ���
int main()
{
	std::cout << "���ü�� ȭ����" <<std::endl;
	std::cout << "��� ȭ����" << std::endl;
	cout << "������� ����" << endl;
	
	//�����Ҵ�
	mainGame mg;

	//�����Ҵ�
	mainGame* mg1 = new mainGame;
	delete mg1;

	return 0;
}