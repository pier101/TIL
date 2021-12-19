

# 파이썬 데ㅔ이터베이스 연동(SQLite)
# 테이블 생성 및 삽입

import sqlite3
import datetime

# 삽입 날짜 생성
now = datetime.datetime.now()
print('now : ',now)

nowDatetime = now.strftime('%Y-%m-%d %H:%M:%S')
print(nowDatetime)



#sqlite3
print('sqlite3.version : ', sqlite3.version)
print('sqlite3.sqlite_version : ', sqlite3.sqlite_version)


# DB 생성 & auto commit(rollback) 
conn = sqlite3.connect('C:/python_workspace/python_basic/resource/database.db',isolation_level=None)

# Cursor
c = conn.cursor()
print('Cursor Type : ', type(c))

# 테이블 생성(Date Type : TEXT, NUMERIC, INTEGER, REAL , BLOB)
c.execute("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY,username text, email text, \
        phone text, website text, regdate text)")

# 데이터 삽입
c.execute("INSERT INTO users VALUES (1 ,'Kim','Kim@naver.com', '010-0000-0000', 'Kim.com', ?)", (nowDatetime,))
c.execute("INSERT INTO users(id, username, email,phone,website,regdate) VALUES(?,?,?,?,?,?)", \
    (2,'Park','Park@daum.net','010-1111-1111','park,com',nowDatetime))

# Many 삽입(튜플, 리스트)
userList = (
    (3, 'Lee', 'Lee@naver.com', '010-2222-2222', 'Lee.com', nowDatetime),
    (4, 'Cho', 'Cho@naver.com', '010-3333-3333', 'Cho.com', nowDatetime),
    (5, 'Yoo', 'Yoo@naver.com', '010-4444-4444', 'Yoo.com', nowDatetime)
)

c.executemany("INSERT INTO users(id,username,email,phone,website,regdate) VALUES(?,?,?,?,?,?)",userList)

# 테이블 데이터 삭제
# conn.execute("DELETE FROM users")
# print("user db deleted : ",conn.execute("DELETE FROM users").rowcount)

# 커밋 : isolation_level = None일 경우 자동 반영(오토 커밋)
#conn.commit()

#롤백
# conn.rollback()

# 접속 해제
# conn.close()