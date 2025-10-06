package main

import "fmt"

type Logger interface {
    Info(string)
    Error(string)
    Debug(string)
}

type ConsoleLogger struct{}
func (c ConsoleLogger) Info(s string)  { 
    fmt.Println("INFO", s) 
}
func (c ConsoleLogger) Error(s string) { 
    fmt.Println("ERROR", s) 
}
func (c ConsoleLogger) Debug(s string) { 
    fmt.Println("DEBUG", s) 
}

type FileLogger struct{}
func (f FileLogger) Info(s string)  { 
    fmt.Print("FILE: INFO ") 
}
func (f FileLogger) Error(s string) { 
    fmt.Print("FILE: ERROR ") 
}
func (f FileLogger) Debug(s string) { 
    fmt.Print("FILE: DEBUG ") 
}

func main() {
    var x Logger = ConsoleLogger{}
    x.Info("информация")
    x = FileLogger{}
    x.Error("ERROR")
}
