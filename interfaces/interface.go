package main

import "fmt"

type Area interface{
	Area() float32
}

type Circle struct{
	radius float32
}

type Square struct{
	side float32
}

type Rectangle struct{
	length, height float32
}

func (c Circle)Area() float32{
	return 3.12*c.radius*c.radius
}

func (s Square)Area() float32{
	return s.side*s.side
}

func (r Rectangle)Area() float32{
	return r.length*r.height
}

func main(){
	circleObj:= Circle{radius:10}
	squareObj:= Square{side:10 }
	rectObj:= Rectangle{length:10,height: 5}

	fmt.Printf("the area of a circle is : %f \n", circleObj.Area())
	fmt.Println("the area of a square is : %f \n", squareObj.Area())
	fmt.Println("the area of a rectangle is : %f \n", rectObj.Area())
}