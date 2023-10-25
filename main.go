package main

import (
	"fmt"

	"github.com/douglira/fullcycle-go-expert/core_packages"
	"github.com/douglira/fullcycle-go-expert/foundation"
)

var titleLessonTemplate = `
######################## %s - %s ########################
`

type Lesson struct {
	Chapter string
	Title   string
	run     func()
}

func runAllLessons(lessons ...Lesson) {
	for _, l := range lessons {
		fmt.Printf(titleLessonTemplate, l.Chapter, l.Title)
		l.run()
	}
}

func main() {
	runAllLessons(
		Lesson{"FOUNDATION", "TYPES", foundation.TypesLesson},
		Lesson{"FOUNDATION", "SLICES", foundation.SlicesLesson},
		Lesson{"FOUNDATION", "STRUCTS", foundation.StructsLesson},
		Lesson{"FOUNDATION", "POINTERS", foundation.PointersLesson},
		Lesson{"FOUNDATION", "TYPE ASSERTATION", foundation.TypeAssertationLesson},
		Lesson{"FOUNDATION", "GENERICS", foundation.GenericsLesson},

		Lesson{"CORE PACKAGES", "FILE MANIPULATION", core_packages.FileManipulationLesson},
		Lesson{"CORE PACKAGES", "DEFER", core_packages.DeferLesson},
	)
}
