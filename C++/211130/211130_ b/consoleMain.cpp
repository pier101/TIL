#include <iostream>
#include "mainGame1.h"

int main()
{
	mainGame1 mg;

	mg.Output("���� ������");
	mg.Output(3.14f);
	mg.Output(7);

	return 0;

	/*
		����
		�迭 ���� ���
		cpp���� �Լ� �ȿ� 
		int num[10000]

		.h(�������)�ȿ�
		int num[10000]
		-----------------
		>>  cpp�Լ� ���� ��� ���ÿ����÷ο� ��Ÿ�� �� ����
		>> .h ���� heap������ �����ؼ� �޸� �Ҵ����� (�����÷ο� �ȳ�)
	*/
}