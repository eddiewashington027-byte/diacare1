package diacare

// Hello returns a greeting message for the given name.
func Hello(name string) string {
    if name == "" {
        name = "world"
    }
    return "hello " + name
}
