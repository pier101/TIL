
/*
	������ : ������ ������ �� �ڵ����� ȣ��Ǵ� �Լ�
			>>��� ���� �ʱ�ȭ
	�Ҹ��� : ��ü�� �Ҹ�� �� �ڵ����� ȣ��Ǵ� �Լ�
			>>�޸� ����
*/

#include <iostream>

using namespace std;
//----------------111---------------------
class MyClass {
public:
	MyClass() { //������
		cout << "������ ȣ���" << endl;
	}
	~MyClass() { //�Ҹ���
		cout << "�Ҹ��� ȣ���" << endl;
	}
};

//MyClass globalObj;

void testLocalObj() {
	cout << "testLocalObj�Լ� ����!" << endl;
	MyClass localObj;
	cout << "testLocalObj�Լ� ��!" << endl;
}
//.........................................

//-------------------2222-----------------
class Complex {
public:
	Complex() {
		real = 0;
		imag = 0;
	}
	Complex(double real_, double imag_) {
		real = real_;
		imag = imag_;
	}
	double GetReal() {
		return real;
	}
	void SetReal(double real_) {
		real = real_;
	}
	double GetImag() {
		return imag;
	}
	void SetImag(double imag_) {
		imag = imag_;
	}
private:
	double real;
	double imag;
};
//--------------------------------------

//------------MAIN----------------------
int main() {
	//11111111111111111111111111111111
	cout << "main�Լ� �� ��!" << endl;
	testLocalObj();
	cout << "main�Լ� ��!" << endl;
	//---------------------------------
	
	//----------222-------------------
	Complex c1;
	Complex c2 = Complex(2, 3);
	Complex c3(2, 3);//c2�� ����ȭ

	cout << "c1 = " << c1.GetReal() << ", " << c1.GetImag() << endl;
	cout << "c2 = " << c2.GetReal() << ", " << c2.GetImag() << endl;
	cout << "c3 = " << c3.GetReal() << ", " << c3.GetImag() << endl;
}
