package main
import "fmt"
func todo(){
	// arr :=[]int{1,2,4,6,7,8}
	arr1:=[]string{"mithilesh","kumar","raghav"}
	arr1=append(arr1,"Varanasi")
    fmt.Println(arr1)
}
func swap(m1,m2 *int){

	var temp int
	temp=*m1
	*m1=*m2
	*m2=temp

}
func (c Car)  GetName()string{
	return c.Name
}
func (c Car) Driving() {
	fmt.Println("Driving the Car ................")
	
}
type Car1 interface{
	Drive()
	stop()
}
type lambo struct{
	Model string
}
type Audi struct{
	Model string
}

func (l *lambo) Drive() {
	fmt.Println("Driving the lambo")
	fmt.Println("moving the model Number",l.Model)
	
}
func (a *Audi) Drive() {
	fmt.Println("Driving the Audi")
	fmt.Println("moving the model Number",a.Model)
	
}
type Car struct{
	Name string
	Age int
	ModelNo int

}

func main() {
	makedict:=make(map[string] interface{})
	makedict["Name"]="Mithilesh"
	makedict["Age"]=23
	makedict["Car"]="Swift vdi"
	fmt.Println(makedict)
	l:=lambo{"pro"}
	a:=Audi{"A8"}
	l.Drive()
	a.Drive()
	

	m1:=2
	m2:=3
    c:=Car {
		Name:   "Audi",
		Age: 2,
		ModelNo: 2020,
	}
	fmt.Println(c)
	fmt.Println(c.GetName())
	c.Driving()
	// fmt.Println(m1,m2)
	// todo()
	swap(&m1,&m2)
	fmt.Println(m1,m2)

}