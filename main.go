//package main
//
//import "fmt"

//func main() {
//	i , j:=15,23
//	pointerToI:=&i
//
//	fmt.Println(*pointerToI)
//	*pointerToI=54
//	fmt.Println(i)
//
//	pointerToI=&j
//	*pointerToI=78
//	fmt.Print(j)
//
//
//}

//type Vertex struct {
//	X int
//	Y int
//}
//
//func main() {
//	fmt.Println(Vertex{1, 2})
//}
//package main
//
//import "fmt"
//
//func main() {
//	var a[2]string
//	a[0]="I am Bot"
//	a[1]=",Hello"
//	fmt.Println(a[0],a[1])
//	numbers:=[2]int{1,5}
//	fmt.Print(numbers)
//}
//package main
//
//import "fmt"
//
//func main() {
//	q := []int{2, 3, 5, 7, 11, 13}
//	fmt.Println(q)
//
//	r := []bool{true, false, true, true, false, true}
//	fmt.Println(r)
//
//	s := []struct {
//		i int
//		b bool
//	}{
//		{2, true},
//		{3, false},
//		{5, true},
//		{7, true},
//		{11, false},
//		{13, true},
//	}
//	fmt.Println(s)
//}
//package main
//
//func main() {
//
//}
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Abser interface {
//	Abs() float64
//}
//
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	var a Abser
//	f := MyFloat(-math.Sqrt2)
//	v := Vertex{3, 4}
//
//	a = f  // a MyFloat implements Abser
//	a = &v // a *Vertex implements Abser
//
//	// In the following line, v is a Vertex (not *Vertex)
//	// and does NOT implement Abser.
//
//	fmt.Println(a.Abs())
//}

//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//}
//
//func (t *T) M() {
//	fmt.Println(t.S)
//}
//
//type F float64
//
//func (f F) M() {
//	fmt.Println(f)
//}
//
//func main() {
//	var i I
//
//	i = &T{"Hello"}
//	describe(i)
//	i.M()
//
//	i = F(math.Pi)
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}

package main

type Car struct {
	Brand, Model, Type string
	Engine             Engine
}
type Runnable interface {
	Launch()
}
type Engine struct {
	Volume float64
	Type string
	Cylinders int
}

func (e *Car)Launch()  {
	print("The ")
	print(e.Engine.Type)
	print(e.Engine.Cylinders)
	print(" volume")
	print();
	print(" engine started")
}
func (car *Car)toString()string{
	return car.Brand+"/"+car.Model+"/"+car.Type+"/"
}


func main() {
	engine:=Engine{6.2,"V",8};

	car := Car{"Chevrolet", "Corvette","Sedan",engine};
	print(car.toString())
	car.Launch()



	}