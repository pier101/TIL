

# 업그레이드 타이핑 게임 제작
# 타이핑 게임 제작 및 기본 완성

import random
import time

import winsound # 사운드 출력 모듈
import sqlite3
import datetime


# db생성 & auto commit 
# 본인 db 경로
conn = sqlite3.connect('C:/python_workspace/python_basic/resource/records.db',isolation_level=None)
# Cursor연결
cursor = conn.cursor()

cursor.execute("CREATE TABLE  IF NOT EXISTS records(id INTEGER PRIMARY KEY AUTOINCREMENT, cor_cnt INTEGER, record text, regdate text )")




words = [] # 영어 단어 리스트(1000개 로드)

n = 1 # 게임 시도 횟수
cor_cnt = 0 # 정답 개수

with open('./resource/word.txt','r') as f:
    for c in f:
        words.append(c.strip())

# print(words) # 단어 리스트 확인

input("Ready? Press Enter Key!") # 게임 스타트

start = time.time()

while n <= 5:
    random.shuffle(words)
    q = random.choice(words)

    print()
    print("*Question # {}".format(n))
    print(q) # 문제 출력
    
    x = input() # 타이핑 입력
    print()
    
    
    if str(q).strip() == str(x).strip(): #입력 확인(공백제거)
        print("Pass!")
        #정답 소리 재생
        winsound.PlaySound('./sound/good.wav',winsound.SND_FILENAME)
        cor_cnt += 1
    else:
        #오답 소리 재생
        winsound.PlaySound('./sound/bad.wav',winsound.SND_FILENAME)
        print("Wrong!")
        
    n+=1

end = time.time() #종료시간

et = end - start
et = format(et,".3f")  

if cor_cnt>=3:
    print("합격")
else:
    print("불합겹")

# 기록 db 삽입
cursor.execute("INSERT INTO records('cor_cnt','record','regdate') VALUES (?,?,?)",(cor_cnt, et, datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')))

#수행 시간 출력
print("게임 시간 : ", et,"초","정답 개수 : {}".format(cor_cnt))
    
# 시작 지점
if __name__=='__main__':
    pass
