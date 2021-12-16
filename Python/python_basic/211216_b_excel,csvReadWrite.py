

# 파이썬 외부 파일 처리
# 파이썬 Excel, CSV 파일 읽기 및 쓰기

# CSV : MIME - text/csv

import csv # csv 파일을 읽기 위해 필요/기본적으로 파이썬 내부에서 제공됨

print('---예제 1---')
#예제1
with open('./resource/sample1.csv','r') as f:
    reader = csv.reader(f)
    next(reader) # header 스킵
    #확인
    print(reader)
    print(type(reader))
    print(dir(reader))  # iter가 있네? > 반복문 사용가능하니 함 해보자
    print()
    
    for c in reader:
        print(c)


print()
print('---예제 2---')
#예제2
with open('./resource/sample2.csv','r') as f:
    # 현재 프린트 해보니 콤마가 없고 |로 구분된 상태 > 콤마로 구분해주자
    # 구분해주는 옵션 = delimiter
    reader = csv.reader(f, delimiter='|')
    next(reader) # header 스킵
    #확인
    print(reader)
    print(type(reader))
    print(dir(reader))  
    print()
    
    for c in reader:
        print(c)
        
print()
print('---예제 3---')
# 예제3 (Dict변환)
with open('./resource/sample1.csv','r') as f:
    reader = csv.DictReader(f) 
    
    for c in reader:
        for k,v in c.items():
            print(k,v)
        print('----------')


print()
print('---예제 4---')
# 예제4
w = [[1,2,3],[4,5,6],[7,8,9],[10,11,12],[13,14,15],[16,17,18]]

# newline : 줄바꿈 옵션 
# 지금 wt변수 내부에서도 줄바꿈하고, open에서 'w'할 떄도 줄바꿈되기 때문에
# 이중 줄바꿈되어 newline 써서 open시에 줄바꿈 없애줌
with open('./resource/sample3.csv','w',newline='') as f:
    wt = csv.writer(f)
    
    for v in w:
        wt.writerow(v) # 예제4랑 차이 : 데이터별로 검수할 땐 요걸로 써주는게 맞음(ex)if 조건처리 필요시) 
        

print()
print('---예제 5---')
# 예제4
with open('./resource/sample4.csv','w',newline='') as f:
    wt = csv.writer(f)
    wt.writerows(w) # 예제3 과 달리 순회 안하고 한 번에 쓸 때

#================================================================

# XSL, XLSX

# 엑셀을 처리하는 오픈소스들
# openpyxl, xlsxwriter, xlrd, xlwt, xlutils
# pandas 를 주로 사용 (openpyxl, xlrd 를 pandas가 내부적으로 사용함)

# 가상환경에서 설치
# pip install xlrd   설치 필요
# pip install openpyxl 설치 필요
# pip install pandas 설치 필

import pandas as pd

# 옵션 살펴보기
# sheetname='시트명' 또는 숫자 >>해당 시트 가져오기
# header=숫자 >> 몇 번째 열을 header로 지정할지
# skiprow=숫자 >> 몇 번째 행을 skip할지

xlsx = pd.read_excel('./resource/sample.xlsx')

# 상위 데이터 확인
print(xlsx.head())
print()

# 데이터 확인
print(xlsx.tail())
print(xlsx.shape) #행 ,열


# 엑셀 or CSV 다시 쓰기
xlsx.to_excel('./resource/result.xlsx',index=False) #index 열에다 숫자 붙여줌
xlsx.to_csv('./resource/result.csv',index=False)