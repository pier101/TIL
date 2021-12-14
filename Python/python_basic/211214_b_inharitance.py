
# 상속, 다중 상속



print('----상속----')
# 예제 1 - 상속 기본
# 슈퍼클래스(부모) 및 서브클래스(자식) ->모든 속성, 메소드 사용 가능

#   라면 -> 속성(종류, 회사, 맛, 면 종류, 이름)  : 부모

class Car:
    """Parent Class"""
    def __init__(self,tp, color):
        self.tp = tp
        self.color = color
    
    def show(self):
        return 'Car Class "Show Method!"'

class BMWCar(Car):
    """Sub Class"""
    def __init__(self,car_name, tp, color):
        super().__init__(tp, color) #부모로부터 상속 받아옴
        self.car_name = car_name
        
    def show_model(self):
        return "your Car Name : %s" % self.car_name

class BenzCar(Car):
    """Sub Class"""
    def __init__(self,car_name, tp, color):
        super().__init__(tp, color) #부모로부터 상속 받아옴
        self.car_name = car_name
        
    def show_model(self):
        return "your Car Name : %s" % self.car_name
    def show(self): # 부모한테 있던 메소드명 내용은 다르게(메소드 오버라이딩일어남)
        print(super().show())
        return 'Car Info : %s %s %s' % (self.car_name, self.tp, self.color)
    
# 일반 사용
model1 = BMWCar('520d', 'sedan', 'red')

print(model1.color)        #super
print(model1.tp)           #super
print(model1.car_name)     #sub
print(model1.show())       #super
print(model1.show_model()) #sub
print(model1.__dict__)

print()
# Method Overriding(오버라이딩)
model2 = BenzCar('Eclass', 'suv', "black")
print(model2.show())

# Parent Method Call
model3 = BenzCar('Cclass', 'sedan', 'silver')
print(model3.show())

print()
# Inheritance Info(상속 정보)
print(BMWCar.mro())
print(BenzCar.mro())


print()
print('----다중 상속----')
# 예제2 - 다중 상속
class X():
    pass

class Y():
    pass

class Z():
    pass

class A(X, Y):
    pass
class B(Y, Z):
    pass
class M(B, A, Z):
    pass

print(M.mro())
print(A.mro())