package airportrobot
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, g Greeter) string {
    return "I can speak " + g.LanguageName() + ": " + g.Greet(name)
}

type GermanGreeter struct{}
func (GermanGreeter) LanguageName() string {
    return "German"
}

func (GermanGreeter) Greet(name string) string {
    return "Hallo " + name + "!"
}

type Italian struct{}
func (Italian) LanguageName() string {
    return "Italian"
}

func (Italian) Greet(name string) string {
    return "Ciao " + name + "!"
}

type Portuguese struct{}
func (Portuguese) LanguageName() string {
    return "Portuguese"
}

func (Portuguese) Greet(name string) string {
    return "Olá " + name + "!"
}

