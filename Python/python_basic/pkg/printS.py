def prt1():
    print("I'm Niceboy")
def prt2():
    print("I'm Goodboy")
    
# 단위 실행(독립적으로 파일 실행) - 해당함수가 모듈로 써지기 전에 잘 작동 되는지 테스트하기 위함
if __name__ == "__main__":
    prt1()
    prt2()