#pragma once
class mainGame
{
private:
	int _hp;
	int _mp;


public:
	mainGame();
	//생성자 오버로딩 가능
	mainGame(int hp);
	mainGame(int hp, int mp);
	~mainGame();

};

