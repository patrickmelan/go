class Node:

    def __init__(self, data):
        self.data = data
        self.next = None
    
    def setNext(self, nextNode):
        self.next = nextNode

    

a = Node("a")
b = Node("b")
c = Node("c")

head = a
a.next = b
b.next = c

print(head.next.next.data)