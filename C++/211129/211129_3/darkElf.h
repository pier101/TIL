#pragma once
#include <iostream>

using namespace std;
class darkElf
{

protected:
	int _hp;
	int _mp;
	int _atk;
	float _speed;

public:
	darkElf();
	~darkElf();
	virtual void Qskill();
	virtual void Wskill();
	virtual void Eskill();
	virtual void Rskill();
};

