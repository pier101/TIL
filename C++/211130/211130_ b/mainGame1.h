#pragma once
#include <iostream>

using namespace std;
class mainGame1
{
public:
	//mainGame1();
	//~mainGame1();

	//함수 템플릿의 형태, 유효 범위는 함수 하나에 국한된다.
	template<typename T>
	void Output(T t);
};

template<typename T>
inline void mainGame1::Output(T t)
{
	cout << t << endl;
}
