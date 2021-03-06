# coding: utf-8

### 数组快排
def quicksort2(arr):
    qsort2(arr,0,len(arr)-1)
    return arr


def qsort2(arr,left, right):
    if left < right:
        pivot = partition2(arr,left, right)
        qsort2(arr,left,pivot - 1)
        qsort2(arr,pivot + 1,right)


def partition2(arr,left, right):
    pivot = arr[left]
    i = left
    j = right
    while i<j:
        while i<j and arr[j] >= pivot:
            j-=1
        arr[i]=arr[j]
        while i < j and arr[i] <= pivot:
            i += 1
        arr[j] = arr[i]
    arr[i]=pivot
    return i

arr = [2,12,23,3,5,8,188]
print quicksort2(arr)

### 链表快排
class LNode(object):
    def __init__(self, val):
        self.val = val
        self.next = None
        self.pre = None


class LList(object):
    def __init__(self):
        self.head = LNode(-1)
        self.tail = self.head
        self.len = 0
        self.head.next = None
        self.head.pre = None

    def append(self, val):
        new = LNode(val)
        if self.head is None:
            self.head = LNode(-1)
            self.len = 1
            self.head.next = new
            new.pre = self.head
        else:
            self.len+=1
            lastNode = self.head
            while lastNode.next:
                lastNode = lastNode.next

            lastNode.next = new
            new.pre = lastNode
        self.tail = new

    def show(self):
        if self.head is None:
            return
        cur = self.head.next
        while cur:
            print(cur.val)
            cur = cur.next

    def showreverse(self):
        if self.head is None:
            return
        cur = self.tail
        while cur:
            print(cur.val)
            cur = cur.pre

    def length(self):
        return self.len

    def get(self,index):
        if self.head is None:
            return
        cur = self.head.next
        index-=1
        while cur and index>0:
            cur = cur.next
            index-=1
        if index == 0:
            return cur.val
    def set(self,index, val):
        if self.head is None:
            return
        cur = self.head.next
        index-=1
        while cur and index>0:
            cur = cur.next
            index-=1
        if index == 0:
            cur.val = val

    def swap(self,i,j):
        ival = self.get(i)
        jval = self.get(j)
        self.set(i,jval)
        self.set(j,ival)



lTest = LList()
lTest.append(2)
lTest.append(12)
lTest.append(23)
lTest.append(3)
lTest.append(5)
lTest.append(8)
lTest.append(188)

#print(lTest.get(3))

# print(lTest.length())
# lTest.show()
# lTest.showreverse()


def quicksort(l):
    if l.head is None:
        return l
    qsort(l,1,l.length())
    return l


def qsort(l,left, right):
    if left < right:
        pivot = partition(l,left, right)
        qsort(l,left,pivot - 1)
        qsort(l,pivot + 1,right)


def partition(l,left, right):
    pivot = l.get(left)
    i = left
    j = i+1
    while j<=right:
        while j<=right and l.get(j) >= pivot:
            j+=1
        if j<= right:
            i+=1
            l.swap(i,j)
            j+=1
    l.swap(i,left)
    return i


quicksort(lTest).show()
