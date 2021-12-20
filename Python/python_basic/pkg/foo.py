class Foo:
    def __init__(self,name):
        self.name = name
    
    def speak(self):
        print('i am' + self.name)  
    def hi(name):
        print(name) 

class Bar(Foo):
    def __init__(self, name):
        super().__init__(name)  
    
    def speak(self):
        print('you are' + self.name)