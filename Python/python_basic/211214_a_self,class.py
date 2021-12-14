
# Self, class , 인스턴스 변수

# 클래스, 인스턴스 차이 중요
# 네임스페이스 : 객체를 인스턴스화 할 때 저장된 공간
# 클래스 변수 : 직접 사용가능, 객체보다 먼저 생성
# 인스턴스 변수 : 객체마다 별도로 존재, 인스턴스 생성후 사용

print("----예제 1----")
# 예제1
class UserInfo:
    #클래스에 들어가는것 : 속성, 메소드
    def __init__(self, name): # 초기화를 위한 함수,첫번째 인수로 self를 지정해야함
        self.name = name
    def user_info_p(self):
        print("name : ",self.name)
        
user1 = UserInfo("Kim") # 인스턴스화된 변수,인스턴스 생성
user1.user_info_p()
user2 = UserInfo("Lee") # 인스턴스화된 변수,인스턴스 생성
user2.user_info_p()

print(id(user1))
print(id(user2))
print(user1.__dict__)
print(user2.__dict__)


print()
print("----예제 2----")
#예제2 - self의 이해
class SelfTest:
    def function1():   # 클래스 메소드
        print('함수1 호출!')
    def function2(self): # 인스턴스 메소드
        print(id(self))
        print('함수 2 호출!')
self_test = SelfTest() #인스턴스 생성
#self_test.function1() # self인자가 없기 떄문에 누구의 함수인지 모름
# 그래서 아래와 같이 호출해줘야 사용 가능하다
SelfTest.function1()
self_test.function2()

print(id(self_test))
SelfTest.function2(self_test) # 인스턴스 변수는 클래스종속으로 호출시 에러뜬다


print()
print("----예제 3----")
#예제3 - 클래스 변수, 인스턴스 변수

class WareHouse:
    #클래스 변수
    stock_num = 0
    def __init__(self,name):
        self.name = name
        WareHouse.stock_num += 1
    def __del__(self): # 인스턴스가 삭제될 때 호출되는 함수
        WareHouse.stock_num -= 1

user1 = WareHouse('Kim')
user2 = WareHouse('Park')
user3 = WareHouse('Lee')

print(user1.__dict__)
print(user2.__dict__)
print(user3.__dict__)
print(WareHouse.__dict__) # 클래스 네임스페이스, 클래스 변수(공유)

print(user1.name)
print(user2.name)
print(user3.name)

print(user1.stock_num) #본인 네임스페이스에 없으면 클래스네이스페이스로 넘어가서 찾음
print(user2.stock_num)
print(user3.stock_num)

del user1
print(user2.stock_num)
print(user3.stock_num)
