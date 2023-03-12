# 数据结构基础

集合结构：数据元素的有限集合。数据元素之间除了“属于同一个集合”的关系之外没有其他关系。

线性结构：数据元素的有序集合。数据元素之间形成一对一的关系。

树型结构：树是层次数据结构，树中数据元素之间存在一对多的关系。

图状结构：图中数据元素之间的关系是多对多的。

## 线性数据结构

### 栈

![栈存储结构示意图](.\md_img\1I0526392-0.png)

向栈中添加元素，此过程被称为"进栈"（入栈或压栈）；
从栈中提取出指定元素，此过程被称为"出栈"（或弹栈）；

#### 数组模拟：

```cpp
int Stack[1005], top = 0;
// 创建一个栈
Stack[++top] = x;
// 添加元素 x
x = Stack[top];
// 读取栈顶元素
top;
// 栈的长度
top--;
// 删除元素
top == 0;
// 判断栈是否为空
```

#### stl 实现

```cpp
#include <stack>

stack<int> st;
// 声明一个栈
st.push(x);
// 添加元素 x
st.top();
// 读取栈顶元素
st.size();
// 栈的长度
st.pop();
// 删除元素
st.empty();
// 判断栈是否为空
```

**空栈取栈顶会报错**

### 队列

只允许在表前端进行删除 在表后端进行插入。

#### 数组实现

```cpp
int Queue[1005],head=1,tail=0;
// 创建
Queue[++tail] = x;
// 添加元素 x
x = Queue[head];
// 读取队首元素
tail - head + 1;
// 长度
head++;
// 删除元素
head > tail;
// 判空
```

#### stl 实现

```cpp
#include <queue>
queue<int>qu;
st.push(x);
// 添加元素 x
st.front();
// 读取队首元素
st.size();
// 长度
st.pop();
// 删除元素
st.empty();
// 判空
```

#### 优先队列

```cpp
priority_queue<int> p_qu;
// 声明 （默认从大到小，与下面等效）
priority_queue<<int>,vector<int>,less<int>>p_l_qu;

priority_queue<<int>,vector<int>,greater<int>>p_g_qu;
// 声明一个 从小到大的优先队列
```

### 图
