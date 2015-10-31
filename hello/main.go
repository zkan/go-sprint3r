package main

//import "errors"
import "fmt"

//import "strconv"

type Num int
type NumTest int

// phone cannot be used in other packages, but the others can be
type Customer struct {
	FirstName string
	LastName  string
	Age       int
	phone     string
}

type AnotherCustomer struct {
	firstName string
	lastName  string
	age       int
}

func (c Customer) Name() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func (c *AnotherCustomer) SetFirstName(firstName string) {
	c.firstName = firstName
}

func (c AnotherCustomer) Name() string {
	return fmt.Sprintf("%s %s", c.firstName, c.lastName)
}

func (n Num) String() string {
	return fmt.Sprintf("Num: %d", int(n))
}

func main() {
	// int, float32, float64, string, bool, byte
	var a int
	a = 11
	var hello string
	hello = "Hello, World!"
	var b byte
	b = 10
	fmt.Println(hello)
	fmt.Printf("Number: %d\n", a)
	fmt.Println("Byte:", b)

	// run-time error
	//n := 0
	//num := 100 / n
	//fmt.Println(num)

	// compile-time error
	//num := 100 / 0
	//fmt.Println(num)

	//name := getName()
	//fmt.Println("Hello", name)

	//name = getName1()
	//fmt.Println("Hello", name)

	for i := 0; i < 10; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()

	i := 0
	switch i {
	default:
		fmt.Println("Default")
	case 0, 1:
		fmt.Println("Zero or One")
		// The line below is used when we want the code to fall through the case below
		//fallthrough
	case 2:
		fmt.Println("Two")
	}

	// infinite loop
	//for {
	//fmt.Println("Hello")
	//}

	for i := 1; i <= 100; i++ {
		fmt.Print(fizzbuzz(i))
		fmt.Print(" ")
	}
	fmt.Println()

	var nums [5]int
	fmt.Println(nums)

	var x int = 10
	var y int = 20
	x, y = swap(x, y)
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

	//var nums1 []int
	nums1 := []int{}
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 9)
	nums1 = append(nums1, 2, 3)
	fmt.Println(nums1)

	nums2 := [5]int{1, 2, 3, 4, 5}
	var slice []int
	slice = nums2[:3]
	slice[1] = 999
	fmt.Println(nums2)
	fmt.Println(slice)

	nums3 := []int{1, 2, 3, 4, 5}
	var slice1 []int
	slice1 = append(slice1, nums3...)
	fmt.Println(nums3)
	fmt.Println(slice1)

	//m := map[string]string{}
	m := make(map[string]string)
	m["name"] = "Kan"
	m["company"] = "Pronto Marketing"
	fmt.Println(m)

	//m1 := map[string]map[string]string{
	//"p1": map[string]string{
	//"name": "Kan",
	//},
	//}
	m1 := map[string]map[string]string{
		"p1": {
			"name": "Kan",
		},
	}
	fmt.Println(m1)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	fmt.Println(arr[:])
	results := []int{}
	results = even(arr[:])
	fmt.Println(results)
	fmt.Println(arr)

	//var add func(int, int) int = func(a, b int) int {
	var add = func(a, b int) int {
		return a + b
	}
	fmt.Println(add(9, 9))

	fmt.Println(filter(func(n int) bool {
		return n%2 == 0
	}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(filter(isEven, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	num, err := div(100, 0)
	if err != nil {
		fmt.Println("Error:", err)
		//return
	} else {
		fmt.Println(num)
	}

	n := Num(10)
	n.printHello()
	//fmt.Println(n.String())
	//fmt.Println(n)

	// If n does not implement String(), it will have error during the compile time
	var s fmt.Stringer
	s = n
	fmt.Println(s.String())

	var s1 interface{}
	s1 = n

	//ss := s1.(fmt.Stringer)
	//fmt.Println(ss.String())

	// Check if the interface can be converted into Stringer interface
	if ss, ok := s1.(fmt.Stringer); ok {
		fmt.Println(ss.String())
	}

	// The other way to check interface
	// s.(type) can be used only with switch/case
	n2 := NumTest(30)
	var s2 interface{}
	s2 = n2
	switch ss1 := s2.(type) {
	case fmt.Stringer:
		fmt.Println(ss1.String())
	default:
		fmt.Printf("Not fmt.Stringer but is %T\n", s2)
	}

	c := Customer{
		FirstName: "Kan",
		LastName:  "Ouivirach",
		Age:       30,
	}
	fmt.Println(c)
	fmt.Println(c.Name())

	c1 := AnotherCustomer{
		firstName: "Kan",
		lastName:  "Ouivirach",
		age:       30,
	}
	c1.SetFirstName("Pronto")
	fmt.Println(c1.Name())
}

// Create a method for our defined type
func (n Num) printHello() {
	for i := 1; i <= int(n); i++ {
		fmt.Println("Hello")
	}
}

func isEven(n int) bool {
	return n%2 == 0
}

func filter(fn func(int) bool, nums []int) []int {
	results := make([]int, 0, len(nums))
	for _, item := range nums {
		if fn(item) == true {
			results = append(results, item)
		}
	}
	return results
}

func even(nums []int) []int {
	results := make([]int, 0, len(nums))
	for _, item := range nums {
		if item%2 == 0 {
			results = append(results, item)
		}
	}
	return results
}

func swap(x, y int) (int, int) {
	return y, x
}

func div(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		// The line below requires the package named "errors"
		//return -1, errors.New("Divided by Zero")
		return 0, fmt.Errorf("Divided by Zero")
	}
	return (num1 / num2), nil
}

func getName() string {
	//var name string
	name := ""
	fmt.Print("Enter Name: ")
	fmt.Scanf("%s", &name)
	return name
}

// Specify the variable to return
func getName1() (name string) {
	fmt.Print("Enter Name: ")
	fmt.Scanf("%s", &name)
	return
}

func fizzbuzz(num int) string {
	//if num%3 == 0 && num%5 == 0 {
	//return "FizzBuzz"
	//} else if num%3 == 0 {
	//return "Fizz"
	//} else if num%5 == 0 {
	//return "Buzz"
	//} else {
	//return strconv.Itoa(num)
	//}

	switch {
	case num%3 == 0 && num%5 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		//return strconv.Itoa(num)
		return fmt.Sprint(num)
	}
}
