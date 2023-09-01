package util

type Person struct {
	Name    string
	Age     int
	Address string
}

type Car struct {
	Brand string
	Model string
	Year  int
}

type ExampleStruct struct {
	PeopleList []Person
	CarList    []Car
	Field1     int64
	Field2     string
	Field3     float64
	Field4     bool
	Field5     []byte
	Field6     map[string]int
	Field7     [1000]int
	Field8     []*ExampleStruct
	Field9     interface{}
	Field10    int
}

func getExampleData() ExampleStruct {
	people := []Person{
		{Name: "Ahmet", Age: 30, Address: "İstanbul"},
		{Name: "Mehmet", Age: 25, Address: "Ankara"},
		{Name: "Ayşe", Age: 28, Address: "İzmir"},
		{Name: "Ayşe", Age: 28, Address: "İzmir"},
		{Name: "Ayşe", Age: 28, Address: "İzmir"},
		{Name: "Ayşe", Age: 28, Address: "İzmir"},
		{Name: "Ayşe", Age: 28, Address: "İzmir"},
	}

	cars := []Car{
		{Brand: "Toyota", Model: "Corolla", Year: 2020},
		{Brand: "Honda", Model: "Civic", Year: 2019},
		{Brand: "Honda", Model: "Civic", Year: 2019},
		{Brand: "Honda", Model: "Civic", Year: 2019},
		{Brand: "Honda", Model: "Civic", Year: 2019},
		{Brand: "Honda", Model: "Civic", Year: 2019},
		{Brand: "Honda", Model: "Civic", Year: 2019},
		{Brand: "Honda", Model: "Civic", Year: 2019},
	}

	return ExampleStruct{
		PeopleList: people,
		CarList:    cars,
		Field1:     1234567890,
		Field2:     "This is a huge struct example",
		Field3:     3.141592653589793,
		Field4:     true,
		Field5:     []byte{1, 2, 3, 4, 5},
		Field6:     map[string]int{"a": 1, "b": 2, "c": 3},
		Field7:     [1000]int{}, // Initializing an array with 1000 elements
		Field8:     []*ExampleStruct{},
		Field9:     nil,
		Field10:    1,
	}
}
