
# 파이썬 흐름제어(반복문)
# 반복문 실습

# 코딩의 핵심 -> 조건 해결 중요

# 기본 반복문 : for, while

v1 = 1
while v1 < 11:
    print("v1 is : ", v1)
    v1+=1
print()
for v2 in range(10):
    print("v2 is : ",v2)
print()    
for v3 in range(1,11):
    print("v3 is : ",v3)
print()

# 1~100합
sum1 = 0
cnt1 = 1
while cnt1 <= 100:
    sum1 += cnt1
    cnt1 +=1

print("1~100 : ",sum1)
print("1~100 : ",sum(range(1,101))) #간단한 방법
print('1~100 : ',sum(range(1, 101, 2))) # 짝수만 더하기


#시퀀스 (순서가 있는)자료형 반복
# 문자열, 리스트,튜플 , 집합, 딕셔너리
# iterable 리턴 함수: range, reversed, enumerate, filter, map , zip
names = ["Kim","Lee","Park"]
for name in names:
    print("You are : ",name)
    
word = "dreams"
for s in word:
    print("Word : ",s)
    
my_info ={
    "name":"Kim",
    "age": 29,
    "city":"Suwon",
}
#기본 값은 key  
for key in my_info: ##default : key 호출
    print("my_info",key)
#values
for key in my_info.values():
    print("my_info",key)
#key & values
for k,v in my_info.items():
    print("my_info",k,v)


name = "KennRY"
for n in name:
    if n.isupper():
        print(n.lower())
    else:
        print(n.upper())

#break
numbers=[14,3,4,6,12,24,17,35,2,15,35,39]
# for num in numbers:
#     if num == 33:
#         print("found : 33!")
#         break
#     else:
#         print("not found: 33!")
#for ~else 구문(반복문이 정상 수행된 경우 else 블럭 수행)
for num in numbers:
    if num == 33:
        print("found : 33!")
        break
    else:
        print("not found: 33!")
else:
    print("not found: 33.........")
    
#continue
lt = ["1",2,5, True,4.3,complex(4)]
for v in lt:
    if type(v) is float:
        continue
    print("type : ",type(v))