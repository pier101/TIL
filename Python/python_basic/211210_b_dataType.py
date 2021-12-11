# 데이터 타입
v_str1= "Niceman"
v_bool= True
v_str2= "Goodboy"
v_float= 10.3
v_int= 7
v_dict={
    "name":"Kim",
    "age": 29
}
v_list = [1,2,3]
v_tuple= 3, 5, 7
v_set = {1,2,3}

print(type(v_str1))
print(type(v_bool))
print(type(v_str2))
print(type(v_bool))
print(type(v_float))
print(type(v_int))
print(type(v_dict))
print(type(v_tuple))
print(type(v_set))


print()

i1 = 39
i2 = 939
big_int = 999999999999999999999999999999999999
big_int2 = 88888888888888888888888888888888888
f1 = 1.234
f2 = 2.333
f3 = .5 
f4 =  10.

print(i1*i2)
print(big_int*big_int2)
print(f1 ** f2)
print(f3 + i2)
result = f3 + i2
print(result, type(result))

print()

a = 5.
b = 4
c = 10
print(type(a),type(b))
result2 = a+b
print(result2, type(result2))
#형변환
#int, float, complex(복소수)
print(int(result2))
print(float(c))
print(complex(3))
print(int(True))
print(int(False))
print(int('3'))
print(complex(False))

print()
# 단항연산자
y = 100
y +=100
print(y)

print()
# 수치 연산 함수
# https://docs.python.org/3/library/math.html
print(abs(-7))
n, m =divmod(100, 8) # 100을 8로 나눈값-n, 나머지=m
print(n , m)

import math

print(math.ceil(5.1)) # 해당값보다 큰것중 가장 작은 정수
print(math.floor(3.874)) # 해당값보다 작은것중 가장 큰 정수