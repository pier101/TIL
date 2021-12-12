

# 딕셔너리, 집합 자료형

# ㅁ 딕셔너리 : 순서X, 중복,X, 수정O, 삭제O
#   key, value

# 선언
a = {'name': 'Kim','Phone': '010-7777-7777', 'birth': 930312}
b = {'arr' : [1, 2, 3,4,5]}
print(type(a))

#출력
print(a['name'])
# print(a['name1'])
print(a.get('name'))
print(a.get('name1'))
print(b['arr'][1:3])

# 딕셔너리 추가
a['address']='Seoul'
print(a)
a['rank'] = [1,3,4]
a['rank2'] = (1,3,4)
print(a)

#key, values, items
print()
print(a.keys())
print(list(a.keys()))
temp=  list(a.keys())
print(temp[1:3])

print(a.values())
print(list(a.values()))

print(a.items())
print(list(a.items()))
print(2 in b)
print('name2' not in a)

# ㅁ 집합 (순서 x, 중복x)
print()
print('======= 잡합 ========')

a= set()
b= set([1,2,3,4])
c= set([1,4,5,6,6])

print(type(a))
print(c) 

t= tuple(b)
print(t)
l = list(b)
print(l)
print()

s1= set([1,2,3,4,5,6])
s2= set([4,5,6,7,8])

print(s1.intersection(s2)) # 교집합
print(s1 & s2) # 교집합

print(s1.union(s2))# 합집합
print(s1 | s2) # 합집합

print(s1.difference(s2))
print(s1 - s2)

#추가 & 제거
print()
print('----추가&제거')
s3 = set([7,8,10,15])

s3.add(18)
print(s3)

s3.remove(15)
print(s3)