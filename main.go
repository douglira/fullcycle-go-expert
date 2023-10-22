package main

import (
	"fmt"

	"github.com/douglira/fullcycle-go-expert/foundation"
)

var titleLessonTemplate = `
######################## %s ########################
`

type Lesson struct {
	Title string
	run   func()
}

func runAllLessons(lessons ...Lesson) {
	for _, l := range lessons {
		fmt.Printf(titleLessonTemplate, l.Title)
		l.run()
	}
}

func main() {
	runAllLessons(
		Lesson{"TYPES", foundation.TypesLesson},
		Lesson{"SLICES", foundation.SlicesLesson},
		Lesson{"STRUCTS", foundation.StructsLesson},
		Lesson{"POINTERS", foundation.PointersLesson},
		Lesson{"TYPE ASSERTATION", foundation.TypeAssertationLesson},
		Lesson{"GENERICS", foundation.GenericsLesson},
	)
}
