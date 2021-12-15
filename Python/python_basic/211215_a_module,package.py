
# 파이썬 모듈과 패키지

# 패키지 예제
# 상대 경로
# .. : 부모 디렉토리
# .  : 현재 디렉토리

#pkg에 파일(모듈 만듦)


#사용1(클래스)
from pkg.fiboonacci import Fibonacci

Fibonacci.fib(300)

print('ex1 : ', Fibonacci.fib2(400))
print('ex1 : ', Fibonacci().title)

#사용2(클래스) - 권장하지는 않음(사용하지 않는 것들도 다 가져옴)
from pkg.fiboonacci import *
Fibonacci.fib(500)

print('ex2 : ', Fibonacci.fib2(600))
print('ex2 : ', Fibonacci().title)

print()
#사용3(클래스) import해오는 이름 길때 as 써서 대체어 사용 가능(Alias) (권장함)
from pkg.fiboonacci import Fibonacci as fb
fb.fib(1000)

print('ex3 : ', fb.fib2(1200))
print('ex3 : ', fb().title)

print()
print('----사용 4----')
#사용4(함수)
import pkg.calculations as c

print("ex4 : ",c.add(10, 100))
print("ex4 : ",c.mul(10, 100))


print()
print('----사용 5----')
#사용5(함수)
from pkg.calculations import div as d
print("ex5 : ",d(100, 10))



print()
print('----사용 6----')
#사용6
import pkg.printS as p
import builtins #파이썬 내장 함수
p.prt1()
p.prt2()
print(dir(builtins))