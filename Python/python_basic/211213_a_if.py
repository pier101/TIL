
# 파이썬 흐름제어(제어문)
# 조건문 실습

print(type(True))
print(type(False))

# ex1
if True:
    print('Yes')
# ex2
if False:
    print('No')    
# ex2
if False:
    print('No')    
else:
    print('Yes2')
    
    
# 관계연산자
# >, >=, <, <=, ==, !=
print()
print('===관곚=연산자===')
a = 10
b = 0
print(a == b)
print(a != b)
print(a >= b)
print(a <= b)
print(a < b)
print(a > b)

# 참 거짓 종류(true,false)
# 참인 경우 : "내용", [내용], (내용), {내용}, 1
# 거짓인 경우 : "", [], (), {}, 0
print()
print('---참 거짓 종류---')
city = ""
if city:
    print('true')
else:
    print('false')
    
    
# 논리연산자
# and or not
aa = 100
bb = 60
cc = 15
print('and : ', aa > bb and bb > cc)
print('or : ',aa > bb or cc > bb)
print('not : ',not aa >bb )


# 산술, 관계, 논리 연산자
# 우선순위 : 산술 > 관계 > 논리 순서로 적용
print('ex1 : ', 5 + 10 > 0 and not 7 + 3 == 10)

score1 = 80
score2 = 'A'

if score1 >=90 and score2 == 'A' : 
    print("합격")
else: 
    print("불합격")
    

# 다중 조건문
num = 70
if num >= 90:
    print("num 등급 A", num)
elif num>=80:
    print("num 등급 B", num)
elif num>=70:
    print("num 등급 C", num)
else:
    print("꽝")

#중첩 조건문
age = 19
height = 158

if age >= 20:
    if height>= 170:
        print("A지망 지원 가능")
    elif height>= 160:
        print("B지망 지원 가능")
    else:
        print("지원 불가")
else:
    print("20세 이상 지원 가능")