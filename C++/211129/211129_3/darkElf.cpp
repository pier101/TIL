#include "darkElf.h"

darkElf::darkElf()
{

	_hp = 100;
	_mp = 50;
	_atk = 10;
	_speed = 3.5f;

	cout << "��ũ������ ü�� : "<<_hp <<endl;
	cout << "��ũ������ ���� : "<<_mp <<endl;
	cout << "��ũ������ ���ݷ� : "<<_atk <<endl;
	cout << "��ũ������ ���ǵ� : "<<_speed <<endl;
}

darkElf::~darkElf()
{
}

void darkElf::Qskill()
{
	cout << "�⺻ ������ Q��ų : Ȱ���" << endl;
}

void darkElf::Wskill()
{
	cout << "�⺻ ������ W��ų : ����" << endl;
}

void darkElf::Eskill()
{
	cout << "�⺻ ������ E��ų : �Ŵ޸���" << endl;
}

void darkElf::Rskill()
{
	cout << "�⺻ ������ R��ų : �ϻ�����" << endl;
}
