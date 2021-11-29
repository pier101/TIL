
#include "mainGame.h"
#include <iostream>

void mainGame1::UseItem()
{
	cout << "집행검을 사용했다" << endl;
}

void mainGame1::useItem(int gas)
{
	cout << gas << "를 소모하여 공격을 하였다" << endl;
}

mainGame1::mainGame1()
{
	cout << "생성자!" << endl;
}

mainGame1::~mainGame1()
{
	cout << "소멸자!" << endl;
}

