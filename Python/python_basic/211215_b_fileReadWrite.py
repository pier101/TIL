
# 파일 읽기, 쓰기

# 읽기 모드 r, 쓰기 모드(기존 파일 삭제) w, 추가 모드(파일 생성 또는 추가) a
# 기타 : https://docs.python.org/3.7/library/functions.html#open
# 상대 경로('../', './'), 절대 경로 확인('C:\...')


# 파일 읽기

# 예제1
f = open('./resource/review.txt','r')
content = f.read()
print(content)
print(dir(f))
# 반드시 close로 리소스 반환
f.close()

print()
print('---예제 2---')
# 예제2
with open('./resource/review.txt','r') as f: #with문을 쓰면 close 안 해도 자동 닫아줌
    c = f.read()
    print(c)
    print(list(c))
    print(iter(c))
    
print()
print('---예제 3---')
# 예제3
with open('./resource/review.txt','r') as f:
    for c in f:
        print(c.strip())
        
    
print()
print('---예제 4---')
# 예제4
with open('./resource/review.txt','r') as f:
    content = f.read()
    print(">", content)
    content = f.read() # 내용 없음
    print(">", content)
    
    
print()
print('---예제 5---')
# 예제5
with open('./resource/review.txt','r') as f:
    line = f.readline() # 라인별 읽어옴
    # print(line)
    while line:
        print(line,end='&&&&&&')
        line = f.readline()
        
print()
print('---예제 6---')
# 예제6
with open('./resource/review.txt','r') as f:
    contents = f.readlines()
    print(contents)
    for c in contents:
        print(c, end=' **** ')
        
print()
print()
print('---예제 6---')
# 예제6
score = []
with open('./resource/score.txt','r') as f:
    for line in f:
        score.append(int(line))
    print(score)
print('Average : {:6.3}'.format(sum(score)/len(score)))


#-----------------------------------
# 파일 쓰기
print()
print('---예제 1---')
# 예제1
with open('./resource/text1.txt', 'w') as f:
    f.write('Niceman!\n')
    
print()
print('---예제 2---')
# 예제2
with open('./resource/text1.txt', 'a') as f:
    f.write('Goodman!\n')

print()
print('---예제 3---')
# 예제3
from random import randint #특정 패키지 가져옴

with open('./resource/text2.txt', 'a') as f:
    for cnt in range(6):
        f.write(str(randint(0,50)))
        f.write('\n')
        
print()
print('---예제 4---')
# 예제4
# writelines : 리스트 > 파일로 저장하는 메소드
with open('./resource/text3.txt', 'a') as f:
    list= ['Kim\n','Lee\n','Park\n']
    f.writelines(list)
    
print()
print('---예제 5---')
# 예제5
with open('./resource/text4.txt', 'a') as f:
    print('Test Contests1!', file=f)
    print('Test Contests2!', file=f)