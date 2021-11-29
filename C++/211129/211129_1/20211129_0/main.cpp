#include <iostream> //input output stream 입출력 라이브러리
#include "mainGame.h" //클래스 불러올땐 큰 따옴표
using namespace std; //계속쓰는거 빼겠다

//항상 메인함수는 프로젝트 당 1개는 필수다
int main()
{
	std::cout << "블록체인 화이팅" <<std::endl;
	std::cout << "블록 화이팅" << std::endl;
	cout << "배고프다 벌써" << endl;
	
	//정적할당
	mainGame mg;

	//동적할당
	mainGame* mg1 = new mainGame;
	delete mg1;

	return 0;
}