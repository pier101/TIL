#include <iostream>

using namespace std;

int main()
{
	int** d;

	const char* star = "*****"

	for (int i = 0; i < length; i++)
	{
		cout << star + i << endl;
	}

	for (int i = 0; i < length; i++)
	{
		cout << star + 4 - i << endl;
	}
	return 0;
}

#include <vector> //�߿�
#include <iostream>
#include "stdafx.h"
using namespace std;
kkkk

int t_main(int argc, _TCHAR* argv[])
{
    vector<int> Vect;

    for (int i = 0; i < 7; i++)
    {
        Vect.push_back(i);//���� �ڿ� �߰�
    }

    for (int i = 0; i < Vect.size(); i++)
    {
        cout << Vect[i] << endl;//���� �迭ó�� ���� ���� ����.
    }

    Vect.clear();//���� ����

    return 0;
}
