#pragma once
#include <iostream>

using namespace std;
class mainGame1
{
public:
	//mainGame1();
	//~mainGame1();

	//�Լ� ���ø��� ����, ��ȿ ������ �Լ� �ϳ��� ���ѵȴ�.
	template<typename T>
	void Output(T t);
};

template<typename T>
inline void mainGame1::Output(T t)
{
	cout << t << endl;
}
