#include <iostream>
#include "mapTest.h"

using namespace std;
mapTest::mapTest()
{
	//pair<자료형1, 자료형2>(실제 넣을 키 값, 실제 넣을 값)
	//make_pair(실제넣을키값, 실제넣을값) 알아서 자료형 찾아서 넣어준다.
	_mapTest.insert(pair<const char*,int>("석훈",3));
	_mapTest.insert(pair<const char*, int>("준혁", 5));
	_mapTest.insert(pair<const char*, int>("해민", 8));
	_mapTest.insert(pair<const char*, int>("해린", 7));

	_mi = _mapTest.find("석훈");

	//맵 컨테이너 안에 키 값이 존재하면
	if (_mi != _mapTest.end())
	{
		cout << _mi->first << "가 좋아하는 숫자는 " << _mi->second << "입니다." << endl;
	}
	else
	{
		cout << "해당값이 존재하지 않습니다." << endl;
	}
}

mapTest::~mapTest()
{
}
