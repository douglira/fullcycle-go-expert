package main

import "github.com/douglira/fullcycle-go-expert/foundation"

var titleLessonTemplate = `
######################## %s ########################
`

func main() {
	foundation.TypesLesson(titleLessonTemplate, "TYPES")
	foundation.SlicesLesson(titleLessonTemplate, "SLICES")
	foundation.StructsLesson(titleLessonTemplate, "STRUCTS")
	foundation.PointersLesson(titleLessonTemplate, "POINTERS")
}
