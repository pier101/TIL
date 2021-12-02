#pragma once
#include <iostream>
#include <string>
#include <vector>

using namespace std;

enum COIN
{
	COIN_BIT, 
	COIN_ETH,
	COIN_RIP,
	COIN_ALT,
	COIN_END
};

class tagCoinInfo
{

public:
	COIN coinKind;
	string name;
	string description;
	int price;
	float change;

};

class exChange
{
private:

	//코인정보가 담긴 클래스를 담을 벡터 컨테이너
	vector<tagCoinInfo> _vCoin;
	vector<tagCoinInfo>::iterator _viCoin;

public:
	void coinSetup();
	void coinOutput(int coinNum);

	vector<tagCoinInfo> getCoinVector() { return _vCoin; }

	exChange();
	~exChange();
};

