package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"
)

type Exercises map[string][]string

func GetExercises() Exercises {
	filename, _ := filepath.Abs("exercises.yaml")
	yamlfile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var exercises Exercises

	err = yaml.Unmarshal(yamlfile, &exercises)
	if err != nil {
		panic(err)
	}

    return exercises
}

func GetRandomExercises() []string {
    exercises := GetExercises()
     make([]string, 7)
}

func GetWorkout() []string {
    workout := []string{"10 min cardio"}
    workout = append(workout, GetRandomExercises()...)
    return workout
}

func main() {
    // exercises := GetExercises()
    // for musclegroup, exercises := range exercises {
    //     fmt.Printf("%v:\n", musclegroup)
    //     for _, exercise := range exercises {
    //         fmt.Printf("  - %v\n", exercise)
    //     }
    // }
    fmt.Printf("%v\n", GetWorkout())
}
