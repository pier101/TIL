
# import this
import sys

# ㅁ 파이썬 기본 인코딩
print(sys.stdin.encoding)
print(sys.stdout.encoding)

# ㅁ 변수 선언
myName = 'Goodboy'

# ㅁ 조건문
if myName == 'Goodboy':
    print('Ok')
    
# ㅁ 반복문
# for i in range(1,10):
#     for j in range(1,10):
#         print('%d * %d = ' % (i,j), i*j)
        
# 함수 선언
def hello():
    print("안녕하세요.")
hello()

# 클래스 
class Cookie:
    pass
# 객체 생성(클래스 생성)
cookie = Cookie()
print(id(cookie))
print(dir(cookie))