# adt
#我语文理解能力太差了，没看出来这一题的数据输入和输出的格式，我的大致想法见下面qwq）
注： 代数数据类型 每一版的改动在日志上 代码用过注释标明了用途
#由于不清楚该题目的数据输入是怎样的 我的代码默认了有一条链表，它的每个结点是如下的结构体：
type List[A any] struct {
    tag  listTag
    head A        
    tail *List[A] 
}

其中 type listTag int
不过我的代码里有初始化链表结点的相关函数 func Nil和func Cons
#中间过程，即题目要求的基本要求中，我的落脚点都是Maybe的类型，即None【A】和Just【A】，有时间会在main函数中将Maybe的类型落脚到输出
