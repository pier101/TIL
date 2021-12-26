# 노드 구현

# 인자 하나만 쓰는 경우
# class Node:
#     def __init__(self,data):
#         self.data = data
#         self.next = None
# 개션된 클래스 형태 _ 인자 두개 쓸 경우
class Node:
    def  __init__(self,data, next=None):
        self.data = data
        self.next = next

# node와 node 연결하기
node1 = Node(1)
node2 = Node(2)
node1.next = node2
head = node1

# 링크드 리스트로 데이터 추가하기
class Node:
    def  __init__(self,data, next=None):
        self.data = data
        self.next = next
        
def add(data):
    node = head
    while node.next: #node의 넥스트가 있으면 실행
        node = node.next
    node.next = Node(data)

node1 = Node(1)
head = node1
print(head)
for index in range(2,10): # 추가
    add(index)

node = head
while node.next: # 출력
    print(node.data)
    node = node.next
print(node.data)

print()
# 중간 데이터 삽입
node = head
node3 = Node(1.5)
search = True   
while search:
    if node.data == 1:
        search = False
    else:
        node = node.next
node_next = node.next
node.next = node3
node3.next = node_next

node = head
while node.next:
    print(node.data)
    node = node.next
print(node.data)

print("==============")
# 파이썬 객체지향 프로그래밍으로 링크드 리스트 구현
class Node:
    def __init__(self,data, next=None):
        self.data = data
        self.next = next

class NodeMg:
    def __init__(self,data):
        self.head = Node(data)
    
    def add(self, data):
        if self.head == "":
            self.head = Node(data)
        node = self.head
        while node.next:
            node = node.next
        node.next = Node(data)
    
    def desc(self):
        node = self.head
        while  node:
            print(node.data)
            node = node.next 

    def delete(self,data):  # 특정 노드를 삭제할 떄 
        if self.head == '':
            print ("해당 값을 가진 노드가 없습니다.")
            return
        
        if self.head.data == data:
            temp = self.head
            self.head = self.head.next
            del temp
        else:
            node = self.head
            while node.next:
                if node.next.data == data:
                    temp = node.next
                    node.next = node.next.next
                    del temp
                    return
                else:
                    node = node.next
    def search_node(self,data):
        node = self.head
        while node:
            if node.data == data:
                print(node.data)
                return node
            else:
                node = node.next
linkedlist1 = NodeMg(1)
# linkedlist1.desc()

for data in range(1,10):
    linkedlist1.add(data)
linkedlist1.desc()

# 특정 노드를 삭제할 떄 
linkedlist1 = NodeMg(0)
linkedlist1.desc()
print(linkedlist1.head)
linkedlist1.delete(0)
print(linkedlist1.head)

linkedlist2 = NodeMg(0)
for data in range(1,10):
    linkedlist2.add(data)
print(linkedlist2.desc())
print()
print(linkedlist2.search_node(4))


# 더블링크드리스트
class Node:
    def __init__(self, data, prev=None, next=None):
        self.prev = prev
        self.data = data
        self.next = next

class NodeMgmt:
    def __init__(self, data):
        self.head = Node(data)
        self.tail = self.head

    def insert(self, data):
        if self.head == None:
            self.head = Node(data)
            self.tail = self.head
        else:
            node = self.head
            while node.next:
                node = node.next
            new = Node(data)
            node.next = new
            new.prev = node
            self.tail = new

    def desc(self):
        node = self.head
        while node:
            print (node.data)
            node = node.next
    
    def search_from_head(self, data):
        if self.head == None:
            return False
    
        node = self.head
        while node:
            if node.data == data:
                return node
            else:
                node = node.next
        return False
    
    def search_from_tail(self, data):
        if self.head == None:
            return False
    
        node = self.tail
        while node:
            if node.data == data:
                return node
            else:
                node = node.prev
        return False
    
    def insert_before(self, data, before_data):
        if self.head == None:
            self.head = Node(data)
            return True
        else:
            node = self.tail
            while node.data != before_data:
                node = node.prev
                if node == None:
                    return False
            new = Node(data)
            before_new = node.prev
            before_new.next = new
            new.prev = before_new
            new.next = node
            node.prev = new
            return True
        
double_linked_list = NodeMgmt(0)
for data in range(1, 10):
    double_linked_list.insert(data)
double_linked_list.desc()
node_3 = double_linked_list.search_from_tail(3)
node_3.data
double_linked_list.insert_before(1.5, 2)
double_linked_list.desc()