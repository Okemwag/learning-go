package main

import "fmt"

// Types in Go are explicit declarations that give meaning to data.
// They are not inheritance hierarchies.
type UserID int

// iota is often used for small enumerations.
// It is useful, but not every constant group should be an enum.
type Status int

const (
	StatusDraft Status = iota
	StatusPublished
	StatusArchived
)

// Type declarations are executable documentation:
// these names communicate domain intent more clearly than raw primitives.
type Message string

// Methods attach behavior to types.
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	if c == nil {
		// Code methods for nil instances when nil is a meaningful state.
		return
	}
	c.Value++
}

func (c Counter) Snapshot() int {
	return c.Value
}

// Embedding supports composition.
type AuditInfo struct {
	CreatedBy string
}

type Document struct {
	AuditInfo
	Title  string
	Status Status
}

func (d Document) Summary() string {
	return fmt.Sprintf("%s by %s", d.Title, d.CreatedBy)
}

// A quick lesson on interfaces:
// interfaces describe behavior, not structure.
type Speaker interface {
	Speak() string
}

type Runner interface {
	Run() string
}

// Embedding interfaces combines required behaviors.
type ActiveSpeaker interface {
	Speaker
	Runner
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "hello, I am " + p.Name
}

func (p Person) Run() string {
	return p.Name + " is running"
}

// Interfaces and nil:
// a nil concrete pointer inside a non-nil interface is a classic trap.
type Resource struct {
	Name string
}

func (r *Resource) Speak() string {
	if r == nil {
		return "<nil resource>"
	}
	return "resource: " + r.Name
}

// The empty interface says nothing about behavior.
// any is an alias for interface{} in modern Go.
func describe(value any) {
	// Use type assertions and type switches sparingly.
	// They are sometimes necessary, but they often signal that a cleaner design exists.
	switch v := value.(type) {
	case string:
		fmt.Println("type switch string:", v)
	case int:
		fmt.Println("type switch int:", v)
	default:
		fmt.Println("type switch unknown:", v)
	}
}

// Function types can bridge into interfaces by giving the function type a method.
type Formatter interface {
	Format(string) string
}

type FormatterFunc func(string) string

func (f FormatterFunc) Format(input string) string {
	return f(input)
}

// Accept interfaces, return structs:
// take behavior at the boundary, return concrete values for maximum usability.
type Clock interface {
	Now() string
}

type FixedClock struct {
	Value string
}

func (f FixedClock) Now() string {
	return f.Value
}

type Service struct {
	clock Clock
}

func NewService(clock Clock) Service {
	return Service{clock: clock}
}

func (s Service) Message() string {
	return "service time: " + s.clock.Now()
}

func main() {
	// Type aliases like UserID make code more expressive than bare ints.
	var id UserID = 42
	fmt.Println("typed ID:", id)

	msg := Message("typed message")
	fmt.Println("message:", msg)

	// Methods can use pointer or value receivers.
	counter := &Counter{Value: 1}
	counter.Increment()
	fmt.Println("counter after pointer receiver:", counter.Snapshot())

	// A nil receiver can still be safe if the method handles it.
	var nilCounter *Counter
	nilCounter.Increment()
	fmt.Println("nil counter method call is safe")

	// Methods are functions too:
	// method expressions let you use a method like a normal function.
	incrementFn := (*Counter).Increment
	incrementFn(counter)
	fmt.Println("counter after method expression:", counter.Value)

	snapshotFn := Counter.Snapshot
	fmt.Println("snapshot via function form:", snapshotFn(*counter))

	// Embedding is composition, not inheritance.
	doc := Document{
		AuditInfo: AuditInfo{CreatedBy: "Ada"},
		Title:     "Go Notes",
		Status:    StatusPublished,
	}
	fmt.Println("embedded field promoted:", doc.CreatedBy)
	fmt.Println("document summary:", doc.Summary())

	// Interfaces are type-safe duck typing:
	// if a type has the right methods, it satisfies the interface implicitly.
	var speaker Speaker = Person{Name: "Gopher"}
	fmt.Println(speaker.Speak())

	var active ActiveSpeaker = Person{Name: "Runner"}
	fmt.Println(active.Speak())
	fmt.Println(active.Run())

	// Interfaces and nil:
	// this interface value is not nil because it contains type information.
	var maybe Speaker = (*Resource)(nil)
	fmt.Println("interface with nil concrete pointer == nil:", maybe == nil)
	fmt.Println("calling method on nil concrete pointer:", maybe.Speak())

	// Interfaces are comparable when their dynamic values are comparable.
	var left any = 10
	var right any = 10
	fmt.Println("comparable interface values:", left == right)

	describe("Go")
	describe(99)

	// A direct type assertion should only be used when the type is truly expected.
	if text, ok := any("assertion").(string); ok {
		fmt.Println("type assertion:", text)
	}

	// Function types can satisfy interfaces.
	upper := FormatterFunc(func(input string) string {
		return "formatted: " + input
	})
	fmt.Println(upper.Format("notes"))

	// Implicit interfaces make dependency injection easier.
	// No framework is required to swap implementations.
	service := NewService(FixedClock{Value: "2026-03-03T12:00:00Z"})
	fmt.Println(service.Message())
}
