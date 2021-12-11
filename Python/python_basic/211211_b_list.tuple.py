

name1 ='Lee'
name2 ='Kim'

# 파이썬 데이터 타입(자료형)
# 리스트 튜플

# ㅁ리스트(순서0, 중복0, 수정0,삭제0)

# 선언
a= []
b = list()
c = [1,2,3,4]
d = [10, 100, 'pen', 'banana']
e = [10, 100,['pen', 'banana']]

# 인덱싱
print(d[3])
print(d[-1])
print(d[0] + d[1])
print(e[2][1])

# 슬라이싱
print()
print(d[0:3])
print(e[2][1:3])

# 연산   
print()
print("----연산----")
print(c + d)
print(c*3) 
print(str(c[0])+'hi')

# 리스트 수정, 삭제
print()
print("----리스트 수정,삭제----")
c[0]=77
print(c)
c[1:2]= [200, 300, 400]
print(c)
# c[1]=['a', 'b', 'c'] #배열안의 값만 들어감
c[1:2]=['a', 'b', 'c'] #배열 통째로 들어감
print(c)

# 삭제
print()
del c[1]
print(c)

# 리스트 함수
print()
print("----리스트 함수----")
y=[5, 2, 3, 1, 4]
print(y)
y.append(6)
print(y)
y.sort() # 정렬
print(y)
# y.reverse()
y.insert(2, 7) # 2번 인덱스에 7 넣어줌
print(y)
y.remove(2) # del처럼 인덱스가 아닌 값을 찾아 지워줌
print(y)
y.pop()# 맨 마지막 원소 꺼냄 LIFO(스택)
print(y)
ex = [88, 77]
y.extend(ex) # 배열 자체가 아닌 값으로 넣어줌
# y.append(6) # 배열 자체를 넣어줌
print(y)

# 삭제 : del, remove, pop  

# ㅁ 튜플 (순서0, 중복0, 수정X,삭제X)
# where? 변경되면 실행이 오작동되는 중요값들에 쓰임
print()
print('==========튜 플 ===========')
a = () #소괄호 사용
b = (1,)
c = (1, 2, 3, 4)
d = (10, 100, ('a', 'b', 'c'))
print(c[2])
print(c[3])
print(d[2][2])

print(d[2:])
print(d[2][0:2])

print(c + d)
print(c*3)

# 튜플 함수
print()
print('-------- 튜플 함수--------')
z=(5,2,1,3,4)
print(z)
print(3 in z)
print(z.index(2)) #해당 값의 인덴스 번호 찾기
print(z.count(1)) # 1 몇개 있음?