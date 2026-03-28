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

func MatchList[A any, R any](l List[A], onNil func() R, onCons func(head A, tail List[A]) R) R {
	switch l.tag {
	case tagNil:
		return onNil()
	case tagCons:
		return onCons(l.head, *l.tail)
	default:
		panic("Unknown List state")
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

func MatchMaybe[A any, R any](m Maybe[A], onNone func() R, onJust func(A) R) R {
	switch m.tag {
	case tagNone:
		return onNone()
	case tagJust:
		return onJust(m.value)
	default:
		panic("Unknown Maybe state")
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

func MatchEither[A any, B any, R any](e Either[A, B], onLeft func(A) R, onRight func(B) R) R {
	switch e.tag {
	case tagLeft:
		return onLeft(e.left)
	case tagRight:
		return onRight(e.right)
	default:
		panic("Unknown Either state")
	}
}

//实现head函数 对于nil list返回的是0 
func Head[A any](l List[A]) Maybe[A] {
	return MatchList(l,
		func() Maybe[A] { return None[A]() },
		func(h A, t List[A]) Maybe[A] { return Just(h) },
	)
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

//Maybe 的 fmap
func MapMaybe[A any, B any](f func(A) B, ma Maybe[A]) Maybe[B] {
	return MatchMaybe(ma,
		func() Maybe[B] { return None[B]() },
		func(val A) Maybe[B] { return Just(f(val)) },
	)
}

//List 的 fmap 它返回的是一个List的类型 要么是Cons要么是None
func MapList[A any, B any](f func(A) B, la List[A]) List[B] {
	onNil := func() List[B] { 
        return Nil[B]() 
    }
	onCons := func(h A, t List[A]) List[B] {
        newHead := f(h)           
        newTail := MapList(f, t)  
        return Cons(newHead, newTail) 
    }
	return MatchList(la, onNil, onCons)
}

//maybe 的 Monad
func ReturnMaybe[A any](a A) Maybe[A] {
    return Just(a)
}

func BindMaybe[A any, B any](m Maybe[A], f func(A) Maybe[B]) Maybe[B] {
    return MatchMaybe(m,
        func() Maybe[B] { 
            return None[B]() 
        },
        func(a A) Maybe[B] {
            return f(a) 
        },
    )
}