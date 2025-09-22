package privatelib

import "fmt"

type Worker struct {
}

func (*Worker) Work() {
	fmt.Printf("Worker is working!")
}
