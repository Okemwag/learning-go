package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name string
	Age  int
}

type hiddenData struct {
	public  string
	private string
}

type Record struct {
	ID    uint32
	Flags uint16
	Code  uint16
}

func main() {
	demoReflectionBasics()
	demoMakeNewValues()
	demoNilWithReflection()
	demoReflectMarshaler()
	demoBuildFunctionWithReflection()
	demoUnsafeBasics()
	demoUnsafeBinaryConversion()
	demoUnsafeUnexportedFieldAccess()

	// Reflection cannot add methods to a type at runtime.
	// Go types are fixed at compile time.
	fmt.Println("see notes.md for why reflection cannot make methods and why cgo is for integration, not speed")
}

func demoReflectionBasics() {
	// Reflection lets you inspect types at runtime.
	user := User{Name: "Gopher", Age: 10}

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	fmt.Println("reflect type:", t.Name())
	fmt.Println("reflect kind:", t.Kind())
	fmt.Println("reflect value field count:", v.NumField())
}

func demoMakeNewValues() {
	// reflect.New creates a new zero value of a type and returns a pointer Value.
	t := reflect.TypeOf(User{})
	ptr := reflect.New(t)

	// Elem gives the pointed-to value.
	elem := ptr.Elem()
	elem.FieldByName("Name").SetString("Created with reflect.New")
	elem.FieldByName("Age").SetInt(42)

	created := ptr.Interface().(*User)
	fmt.Println("new reflected value:", created.Name, created.Age)
}

func demoNilWithReflection() {
	// A plain interface may be non-nil while holding a nil concrete pointer.
	var maybe any = (*User)(nil)
	fmt.Println("interface == nil:", maybe == nil)
	fmt.Println("reflect-based nil check:", isNilInterfaceValue(maybe))
}

func isNilInterfaceValue(value any) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func demoReflectMarshaler() {
	// Reflection can be used to write generic data marshalers.
	// This is educational, but usually slower and more fragile than typed code.
	user := User{Name: "Asha", Age: 28}
	fmt.Println("reflect marshaled:", marshalStruct(user))
}

func marshalStruct(input any) map[string]any {
	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)

	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	out := make(map[string]any, rv.NumField())
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		out[field.Name] = rv.Field(i).Interface()
	}
	return out
}

func demoBuildFunctionWithReflection() {
	// reflect.MakeFunc can create function values dynamically.
	// This is possible, but it is usually much harder to read than normal code.
	fnType := reflect.TypeOf(func(int, int) int { return 0 })
	add := reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		sum := args[0].Int() + args[1].Int()
		return []reflect.Value{reflect.ValueOf(int(sum))}
	})

	typedAdd := add.Interface().(func(int, int) int)
	fmt.Println("reflect-built function:", typedAdd(3, 4))
}

func demoUnsafeBasics() {
	// unsafe is unsafe because it bypasses normal type and memory safety guarantees.
	var sample Record

	fmt.Println("unsafe.Sizeof Record:", unsafe.Sizeof(sample))
	fmt.Println("unsafe.Offsetof Record.ID:", unsafe.Offsetof(sample.ID))
	fmt.Println("unsafe.Offsetof Record.Flags:", unsafe.Offsetof(sample.Flags))
	fmt.Println("unsafe.Offsetof Record.Code:", unsafe.Offsetof(sample.Code))
}

func demoUnsafeBinaryConversion() {
	// Using unsafe to convert external binary data can avoid copying,
	// but it depends on alignment, layout, and endianness assumptions.
	data := [8]byte{0x78, 0x56, 0x34, 0x12, 0xCD, 0xAB, 0xEF, 0xBE}

	// Safer standard-library approach:
	id := binary.LittleEndian.Uint32(data[0:4])
	flags := binary.LittleEndian.Uint16(data[4:6])
	code := binary.LittleEndian.Uint16(data[6:8])
	fmt.Println("binary package parse:", id, flags, code)

	// Unsafe reinterpretation:
	record := *(*Record)(unsafe.Pointer(&data[0]))
	fmt.Println("unsafe reinterpretation:", record.ID, record.Flags, record.Code)
}

func demoUnsafeUnexportedFieldAccess() {
	// Accessing unexported fields through reflect is intentionally restricted.
	data := hiddenData{public: "visible", private: "secret"}

	rv := reflect.ValueOf(&data).Elem()
	privateField := rv.FieldByName("private")
	fmt.Println("private field can set normally via reflect:", privateField.CanSet())

	// unsafe can bypass that restriction. This is powerful and dangerous.
	writable := reflect.NewAt(privateField.Type(), unsafe.Pointer(privateField.UnsafeAddr())).Elem()
	writable.SetString("changed unsafely")

	fmt.Println("unexported field after unsafe write:", data.private)
}
