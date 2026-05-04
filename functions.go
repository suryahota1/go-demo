package main
import ("fmt")

func add(a rune, b int32) rune {
	var res = a + b
	fmt.Println("res", res)
	return res
}

func sumWithNamedReturn (x int32, y int32) (res int32) {
	res = x + y
	return
}

// variadic function
func varSum ( nums ...rune ) rune {
	var res rune;
	for _, val := range nums {
		res += val
	}
	return res
}

func greetAll ( prefix string, names ...string) string {
	if len(names) == 0 {
		return "No one to greet"
	}
	res := fmt.Sprintf("%s %s", prefix, names[0])
	for _, val := range(names[1:]) {
		res += "," + val
	}
	return res
}

type User struct {
	Name string
	Age int
	Address string
}

func updateUser ( attrs ...interface{} ) User {
	user := User{Name: "Surya", Age: 32, Address: "Odisha"}
	for i := 0; i < len(attrs); i++ {
		var key string = attrs[i].(string)
		switch(key) {
			case "age": 
				user.Age = attrs[i+1].(int)
				i++
			
			case "address": 
				user.Address = attrs[i+1].(string)
				i++
			
		}
	}
	return user
}

// Multiple return values
func multiReturn(nums []rune)(rune, rune, float64) {
	if len(nums) == 0 {
		return 0, 0, 0.0
	}

	var sum rune
	sum = 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}

	max := nums[0]
	for _, val := range(nums[1:]) {
		if val > max {
			max = val
		}
	}

	average := float64(sum) / float64(len(nums))

	return sum, max, average
}

func main () {
	add(43, 51)
	fmt.Println("named res", sumWithNamedReturn(12, 54))
	fmt.Println("variadic response", varSum(1, 2, 5))
	fmt.Println("variadic response", varSum(7, 2, 5))
	fmt.Println("variadic response", varSum()) // No args, returns 0

	slice1 := []int32{5, 6, 7}
	fmt.Println("variadic response", varSum(slice1...)) // unpacking slice
	fmt.Println("greet response", greetAll("Hello", "Surya", "Sashi"))
	fmt.Println("updated user", updateUser("age", 27, "address", "Bangalore"))

	sum, max, avg := multiReturn([]rune{5, 10, 4, 12, 7})

	fmt.Println("Multi return values", sum, max, avg)
}