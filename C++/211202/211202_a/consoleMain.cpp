#include <iostream>

using namespace std;

//struct == ����ü == ��ü(����Ʈ ���� ���� �����ڰ� public�ΰ� Ư¡)
struct tagUnit
{
	const char* name; //������ �̸�
	int hp;			   //������ ü��
	int atk;		   //������ ���ݷ�
};

//�Ʒ� �Լ� ����� call by value Ÿ���̱� ������ ü���� �ٲ��� �ʴ´�
//������Ű�� �Լ�(����� �̸�, ü��, ������ �̸�, ���ݷ�
//void battle(const char* name, int hp, const char* name2, int atk)
//{
//	cout << name2 << "��(��) " << name << "�� �����ߴ�" << endl;
//	hp -= atk;
//
//	cout << name << "��(��) ü���� " << hp << " ���Ҵ�" << endl;
//	cout << endl << endl;
//}

//������Ű�� �Լ�(����� �̸�, ü��, ������ �̸�, ���ݷ�

//������Ű�� �Լ�
void battle(const char* name, int *hp, const char* name2, int atk)
{
	cout << name2 << "��(��)" << name << "�� �����ߴ�" << endl;
	*hp -= atk;

	cout << name << "��(��)ü���� " << *hp << "���Ҵ�" << endl;

	cout << endl << endl;


};

int main() {
	tagUnit dragoon;
	dragoon.name = "���";
	 dragoon.hp = 100;
	 dragoon.atk = 30;
	 
	 tagUnit vulture;
	 vulture.name = "��ó";
	 vulture.hp = 80;
	 vulture.atk = 10;
	 /* ü���� ��ȭ�� ����
 battle(vulture.name, vulture.hp, dragoon.name, dragoon.atk);
 battle(dragoon.name, dragoon.hp, vulture.name, vulture.atk);

 battle(vulture.name, vulture.hp, dragoon.name, dragoon.atk);
 battle(dragoon.name, dragoon.hp, vulture.name, vulture.atk);
 */

 //ü���� ��ȭ�� �����
	 battle(vulture.name, &vulture.hp, dragoon.name, dragoon.atk);
	 battle(dragoon.name, &dragoon.hp, vulture.name, vulture.atk);
	 battle(vulture.name, &vulture.hp, dragoon.name, dragoon.atk);
	 battle(dragoon.name, &dragoon.hp, vulture.name, vulture.atk);


	 //const char* star = "*****";

	 //for (char* i = 0; i < star.; i++)
	 //{

	 //}





	return 0;
}