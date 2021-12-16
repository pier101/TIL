
# 파이썬 예외처리 이해

# 예외 종류
# 문법적으로 에러가 없지만, 코드 실행(런타임)프로세스에서 발생하는 예외 처리도 중요
# linter : 코드 스타일 , 문법 체크, 


# 1. SyntaxError : 잘못된 문법
    # print('Test)

    # if True
    #     pass
    

# 2. NameError : 참조변수 없음
    # a = 10
    # b = 15
    # print(c)

# 3. zeroDivisionError : 0 나누기 에러
    #print(10 /0)

# 4. IndexError : 인덱스 범위 오버
# x = [10,20,30]
# print(x[0])
# print(x[3])

# 5. KeyError
dic = {'name':'Kim','age':33,'city':'seoul'}
# print(dic['hobby'])
print(dic.get('hobby')) # get쓰면 에러처리 안하고 None으로 나타내줌


# 6. AttributeError : 모듈, 클래스에 있는 잘못된 속성 사용시에 예외
import time
print(time.time())
#print(time.month()) #없는 속성 사용


# 7. ValueError : 참조 값이 없을 때 발생
#x = [1, 5, 9]
#x.remove(10)


# 8. FileNotFoundError
#f = open('test.txt','r')


# TypeError
x = [1, 2]
y = (1, 2)
z = 'test'
#print(x + z) #예외
#print(x + y) #예외
print(x + list(y))

##=================================================


# 항상 예외가 발생하지 않을 것으로 가정하고 먼저 코딩
# 그 후 런타임 예외 발생시 예외 처리 코딩 권장(EAFP 코딩 스타일)


##==================================================

# 예외 처리 기본
# try : 에러가 발생할 가능성이 있는 코드 실행
# except : 에러명1
# except : 에러명2
# else : 에러가 발생하지 않았을 경우 실행
# finally : 항상 실행

print()
print('---예제 1---')
# 예제1
name = ['kim','lee','park']

try:
    z = 'kim'
    x = name.index(z)
    print('{} Found it! in name'.format(z, x+1))
except ValueError:
    print('Not found it! - Occurred ValueError!')
else:
    print('Ok! else!')


print()
print('---예제 2---')
# 예제2
try:
    z = 'jin' 
    x = name.index(z)
    print('{} Found it! in name'.format(z, x+1))
except :
    print('Not found it! - Occurred !')
else:
    print('Ok! else!')
    

print()
print('---예제 3---')
# 예제3
try:
    z = 'kim' 
    x = name.index(z)
    print('{} Found it! in name'.format(z, x+1))
except :
    print('Not found it! - Occurred !')
else:
    print('Ok! else!')
finally:
    print('finally OK!')


print()
print('---예제 4---')
# 예제4
# 예외 처리는 하지 않지만, 무조건 수행되는 코딩 패턴
try:
    print('Try')
finally:
    print('Ok Finally!')


print()
print('---예제 5---')
# 예제5
try:
    z = 'jin' 
    x = name.index(z)
    print('{} Found it! in name'.format(z, x+1))
except ValueError as l: # Alias도 줄 수 있음
    print('Not found it! - Occurred ValueError!')
    print(l)
except IndexError: 
    print('Not found it! - Occurred IndexError!')
except Exception: #Exception 디폴트값이라 없어도 됨
    print('Not found it! - Occurred !')
else:
    print('Ok! else!')
finally:
    print('finally OK!')


print()
print('---예제 6---')
# 예제6
# 예외 발생 : raise
# raise 키워드로 예외 직접 발생 시킬 수 있음
try:
    a = 'Kim'
    if a == 'Keim':
        print('Ok 허가')
    else:
        raise ValueError
except ValueError:
    print('문제 발생!')
except Exception as f:
    print(f)
else:
    print('Ok')