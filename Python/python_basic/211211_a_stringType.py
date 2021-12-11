
#문자열, 문자열연산 , 슬라이싱

str1= "I am Boy."
str2= "NiceMan"
str3 = ''
str4 = str()
print(len(str1),len(str2),len(str3),len(str4))

escape_str1 = "Do you have a \"big collection\""
print(escape_str1)
escape_str2 = "Tab\tTab\tTab"
print(escape_str2)

# Raw String 
# r붙이면 문자 그대로 출력
raw_s1 = r'C:\Programs\Test\Bin'
print(raw_s1)


print()
# 문자열 연산 
str_o1 = '*'
str_o2 = 'abc'
str_o3 = 'def'
str_o4 = 'Niceman'

print(str_o1*10)
print(str_o2+str_o3)
print('a' in str_o4)
print('z' not in str_o4)

print()

#문자열 형 변환
print(str(77)) #string임
print(str(10.1)) #string임


# 문자열 함수 (많기 떄문에 그때그떄 찾아서 볼것)
a = 'niceman'
b = 'orange'
print(a.islower()) # 소문자가 맞니?
print(a.endswith('e')) #끝 글자가 e로 끝나니?
print(a.capitalize()) # 첫 글자만 대문자로 바꿔줌
print(a.replace('nice', 'good'))
print(list(reversed(b)))

#문자열은 한 번 할당이 되면 반환이 불가능하다. immutable
print(a[0:3])
print(a[0:4])
print(a[0:len(a)])
print(a[0:len(a)-1])
print(a[:4])
print(b[0:4:2]) #세번째 값만큼 정핑
print(b[1:-2]) 
print(b[::-1]) 