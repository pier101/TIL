#include <iostream>
#include "mainGame1.h"

int main()
{
	mainGame1 mg;

	mg.Output("점심 뭐먹지");
	mg.Output(3.14f);
	mg.Output(7);

	return 0;

	/*
		코테
		배열 만개 써라
		cpp파일 함수 안에 
		int num[10000]

		.h(헤더파일)안에
		int num[10000]
		-----------------
		>>  cpp함수 안의 경우 스택오버플로우 나타날 수 있음
		>> .h 안은 heap영역에 존재해서 메모리 할당해줌 (오버플로우 안남)
	*/
}