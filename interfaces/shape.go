package interfaces

type Shape interface {
	Area() int
}


type Square struct{
	side int
}

func (this Square) Area() int{
	return this.side * this.side
}

type Rectangle struct{
	length int
	breadth int
}

func (this Rectangle) Area() int{
	return this.length * this.breadth
}
type Hybrid struct{
	 s1 Shape
	 s2 Shape
 }

func (this Hybrid) Area() int{
	return this.s1.Area() + this.s2.Area()
}
