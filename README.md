import the module
```go
  import "github.com/imshi187/sutil"
```

The following part is a demo of how to use the sutil

```Python
func main() {
  sw := IntSliceWrapper{
  data: []int{10, 20, 30, 40},
}
 
// add
sw.Add(50)
 

// delete element in index == 1
err := sw.Remove(1)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("\nAfter removing index 1:")
}

// modify the element 
err = sw.Update(2, 100)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("\nAfter updating index 2 to 100:")

// 查询索引为 1 的元素
item, err := sw.Get(1)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Printf("\nItem at index 1: %d\n", item)
}

// sort(ascending)
sw.SortAscending()

// sort(descending)
sw.SortDescending()

// filter
sw.Filter(func(x int) bool {
	return x > 30
})
 
}
```




