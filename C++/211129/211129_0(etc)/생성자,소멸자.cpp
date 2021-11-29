
/*
	생성자 : 객제가 생성될 때 자동으로 호출되는 함수
			>>멤버 변수 초기화
	소멸자 : 객체가 소멸될 때 자동으로 호출되는 함수
			>>메모리 해제
*/

#include <iostream>

using namespace std;
//----------------111---------------------
class MyClass {
public:
	MyClass() { //생성자
		cout << "생성자 호출됨" << endl;
	}
	~MyClass() { //소멸자
		cout << "소멸자 호출됨" << endl;
	}
};

//MyClass globalObj;

void testLocalObj() {
	cout << "testLocalObj함수 시작!" << endl;
	MyClass localObj;
	cout << "testLocalObj함수 끝!" << endl;
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
	cout << "main함수 사 직!" << endl;
	testLocalObj();
	cout << "main함수 끝!" << endl;
	//---------------------------------
	
	//----------222-------------------
	Complex c1;
	Complex c2 = Complex(2, 3);
	Complex c3(2, 3);//c2의 간소화

	cout << "c1 = " << c1.GetReal() << ", " << c1.GetImag() << endl;
	cout << "c2 = " << c2.GetReal() << ", " << c2.GetImag() << endl;
	cout << "c3 = " << c3.GetReal() << ", " << c3.GetImag() << endl;
}
