package main
 
import(
	"fmt"
)
//实现list
type listTag int

const (
    tagNil  listTag = 0
    tagCons  listTag = 1          
)

type List[A any] struct {
    tag  listTag
    head A        
    tail *List[A] 
}

func Nil[A any]() List[A] {
    return List[A]{tag: tagNil}
}

func Cons[A any](x A, xs List[A]) List[A] {
    return List[A]{
        tag:  tagCons,
        head: x,
        tail: &xs, 
    }
}

//实现maybe
type maybeTag int
const (
    tagNone maybeTag = 0
    tagJust maybeTag = 1
)

type Maybe[A any] struct {
    tag   maybeTag
    value A 
}

func None[A any]() Maybe[A] {
    return Maybe[A]{tag: tagNone}
}

func Just[A any](val A) Maybe[A] {
    return Maybe[A]{tag: tagJust, value: val}
}

func (m Maybe[A]) Match(onNone func(), onJust func(val A)) {
    switch m.tag {
    case tagNone:
        onNone() 
    case tagJust:
        onJust(m.value)
	default:
        panic("Unknown Maybe state!")
    }
}

// 实现either
type eitherTag int
const (
    tagLeft  eitherTag = 0
    tagRight eitherTag = 1
)

type Either[A any, B any] struct {
    tag   eitherTag
    left  A 
    right B 
}

func Left[A any, B any](a A) Either[A, B] {
    return Either[A, B]{tag: tagLeft, left: a}
}

func Right[A any, B any](b B) Either[A, B] {
    return Either[A, B]{tag: tagRight, right: b}
}

func (e Either[A,B]) Match (onLeft func(A),onRight func(B)){
	switch e.tag{
	case tagLeft:
		onLeft(e.left)
	case tagRight:
		onRight(e.right)
	default:
        panic("Unknown Either state!")
	}
}

//实现head函数
func Head[A any](l List[A]) Maybe[A] {
    switch l.tag {
    case tagNil:
        return None[A]() 
    case tagCons:
        return Just(l.head)
    default:
        panic("Unknown List state!")
    }
}

//除0 的一个either处理错误的例子 
func SafeDivide(a, b float64) Either[string, float64] {
    if b == 0 {
        return Left[string, float64]("error: division by zero")
    }
    return Right[string, float64](a / b)
}

// maybe的哈希表查找
func Lookup[K comparable, V any](m map[K]V, key K) Maybe[V] {
    val, ok := m[key]
    if !ok {
        return None[V]()
	}
    return Just(val)
}

