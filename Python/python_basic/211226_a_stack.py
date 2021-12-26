# 재귀 함수
def recursive(data):
    if data < 0:
        print ("ended")
    else:
        print(data)
        recursive(data - 1)
        print("returned", data) 
recursive(4)


# push ,pop 기능 구현해보기

stack_list = list()

def push(data):
    stack_list.append(data)

def pop():
    data = stack_list[-1]
    del stack_list[-1]
    return data

push(1)
push(2)

print(stack_list)

pop()
print(stack_list)