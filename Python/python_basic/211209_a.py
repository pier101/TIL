
# ㅁ print 구문 이해
# 참조 : https://www.python-course.eu/python3_formatted_output.php

# 기본출력
print('hello')
print("hello")
print('''hello''')
print("""hello""")

print()

# Separator 옵션 사용
print('T','E','S','T', sep='')
print('2019','02','19',sep='-')
print('niceman','google.com',sep='@')

# end 옵션 사용
# 다음 라인과 줄바꿈 안되고 연결되게
print('Welcome To',end='')
print(' the black paradise',end='')
print(' piano notes')
print('줄바꿈됨')

print()

# format 사용 []. {}, () 중에 중괄호 넣어 사용
print('{} and {}'.format('You','Me'))
print("{0} and {1} and {0}".format('You','Me'))
print("{a} are {b}".format(a='You',b='Me')) #좀 더 직관적으로 사용
# %s:문자, %d:정수, %f:실수
print("%s's favorite number is %d" % ('dongwook',3))
print("test1:%5d, Price:%4.2f" %(776,6534.123))
print("test1:{0: 5d}, price:{1:4.2f}".format(776,6534.123))
print("test1:{a: 5d}, price:{b:4.2f}".format(a=776,b=6534.123))

#------------------------------------
print('----------------------------')
"""
참고 : Escape 코드

\n : 개행
\t : 탭
\\ : 문자
\' : 문자
\" : 문자
\r : 캐리지 리턴
\f : 폼 피드
\a : 벨 소리
\b : 백 스페이스
\000 : 널 문자
"""
print("'you'")
print('\'you\'')
print('"you"')
print('\\you\\ \n')
print('test')
print('\ttest')