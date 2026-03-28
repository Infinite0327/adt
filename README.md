# adt
# 代数数据类型 每一版的改动在日志上 代码用过注释标明了用途
#由于不清楚该题目的数据输入是怎样的 我的代码默认了有一条链表，它的每个结点是如下的结构体：
type List[A any] struct {
    tag  listTag
    head A        
    tail *List[A] 
}

其中 type listTag int
不过我的代码里有初始化链表结点的相关函数 func Nil和func Cons
