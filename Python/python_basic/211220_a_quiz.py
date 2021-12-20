#1

# class Foo:
#     def __init__(self,name):
#         self.name = name
    
#     def speak(self):
#         print('i am' + self.name)   

# class Bar(Foo):
#     def __init__(self, name):
#         super().__init__(name)  
    
#     def speak(self):
#         print('you are' + self.name)
        
# bar = Bar('wook')
# bar.speak()

#2
# from pkg.foo import Foo
# print(dir(Foo))
# Foo.hi('wook')

#3
# import datetime
# now = datetime.datetime.now()
# print(now)

#==========================\
# class Median:
#     def __init__(self):
#         pass

#     def get_item(self, item):
#         pass

#     def clear(self):
#         pass

#     def show_result(self):
#         pass

# for x in range(10):
#     median.get_item(x)
# median.show_result()

# median.clear()
# for x in [0.5, 6.2, -0.4, 9.6, 0.4]:
#     median.get_item(x)
# median.show_result()

#================
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        print(self.name + ' cannot speak.')

    def move(self):
        print(self.name + ' cannot move.')


class Dog(Animal):
	def move(self):
		print(self.name + ' moves like a jagger.')

class Retriever(Dog):
    def speak(self):
        print(self.name + ' is smart enough to speak.')


dog = Dog('Nancy')
dog.speak()
dog.move()

super_dog = Retriever('Michael')
super_dog.speak()
super_dog.move()

# Nancy cannot speak.
# Nancy moves like a jagger.
# Michael is smart enough to speak.
# Michael moves like a jagger.

#=========================
class Foo:
    bar = 'A'
    def __init__(self,bar='B'):
        self.bar = bar
    class Bar:
        bar = 'C'
        def __init__(self,bar='D'):
            self.bar = bar
		
		
print(Foo.bar)       # A 출력
print(Foo().bar)     # B 출력
print(Foo.Bar.bar)   # C 출력
print(Foo.Bar().bar) # D 출력

