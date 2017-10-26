package main

import "errors"
import "fmt"
import "sort"
import "time"
import "math"
import "os"

func main() {

    strs := []string{"c", "a", "b", "d", "в"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)
	
    ints := []int{7, 2, 4, 60000}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)

    //-----------------------------------

    const n = 500000000
    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))

	//-----------------------------------

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	//-----------------------------------

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

	//-----------------------------------

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    //-----------------------------------

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    //-----------------------------------

    ss := make([]string, 3)
    fmt.Println("emp:", s)

    ss = append(ss, "d")
    ss = append(ss, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(ss))
    copy(c, ss)
    fmt.Println("cpy:", c)

    tt := []string{"g", "h", "i"}
    fmt.Println("dcl:", tt)

    var twoDD = make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoDD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoDD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoDD)

    //-----------------------------------

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    delete(m, "k2")

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

	var nn = map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", nn)

    //-----------------------------------

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }

    //-----------------------------------

	var aa, bb = vals()
    fmt.Println(aa)
    fmt.Println(bb)

	_, cc := vals()
    fmt.Println(cc)

    numss := []int{1, 2, 3, 4}
    summ(numss...)

    //-----------------------------------

    nextInt := intSeq()
	fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    //-----------------------------------

    iii := 1
    fmt.Println("initial:", iii)
    zeroval(iii)
    fmt.Println("zeroval:", iii)
    zeroptr(&iii)
    fmt.Println("zeroptr:", iii)
    fmt.Println("pointer:", &iii)

    //-----------------------------------

    fmt.Println(&person{name: "Ann", age: 40})
    sss := person{name: "Sean", age: 50}
    fmt.Println(sss.name)
    sp := &sss
    fmt.Println(sp.age)

    //-----------------------------------

    rr := rectt{width: 10, height: 5}
    rp := &rr
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())

    //-----------------------------------

    var re = rect{width: 3, height: 4}
    var ci = circle{radius: 5}
	measure(re)
    measure(ci)

    //-----------------------------------

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }

    //-----------------------------------

    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)

    //-----------------------------------

    panic("a problem")
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }

    //-----------------------------------

    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
    
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
    //-----------------------------------
}

func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 7
}

func summ(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

type person struct {
    name string
    age  int
}

type rectt struct {
    width, height int
}

func (r *rectt) area() int {
    return r.width * r.height
}

func (r rectt) perim() int {
    return 2*r.width + 2*r.height
}

type geometry interface {
    area() float64
    perim() float64
}
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
func f1(arg int) (int, error) {
    if arg == 42 {
    	        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}
type argError struct {
    arg  int
    prob string
}
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
    if arg == 42 {
    	return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

type ByLength []string
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}
func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
