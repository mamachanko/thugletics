package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

const exerciseFilename string = "exercises.yaml"

type MuscleGroup struct {
	Name      string
	Exercises []Exercise
}

type Exercise string

func (m *MuscleGroup) randomExercise() Exercise {
	randomIndex := rand.Perm(len(m.Exercises))[0]
	return m.Exercises[randomIndex]
}

func readExerciseFile() []byte {
	filename, _ := filepath.Abs(exerciseFilename)
	yamlfile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return yamlfile
}

func parseExercises(raw map[string][]string) []MuscleGroup {
	result := []MuscleGroup{}
	for muscleGroupName, exerciseNames := range raw {
		exercises := []Exercise{}
		for _, exerciseName := range exerciseNames {
			exercises = append(exercises, Exercise(exerciseName))
		}
		muscleGroup := MuscleGroup{Name: muscleGroupName, Exercises: exercises}
		result = append(result, muscleGroup)
	}
	return result
}

func GetExercises() []MuscleGroup {
	var muscleGroups map[string][]string
	err := yaml.Unmarshal(readExerciseFile(), &muscleGroups)
	if err != nil {
		panic(err)
	}
	return parseExercises(muscleGroups)
}

func GetRandomExercises() []Exercise {
	muscleGroups := GetExercises()
	exercises := []Exercise{}
	for _, muscleGroup := range muscleGroups {
		exercises = append(exercises, muscleGroup.randomExercise())
	}
	return exercises
}

func GetWorkout() []Exercise {
	workout := []Exercise{"10 min cardio"}
	workout = append(workout, GetRandomExercises()...)
	return workout
}

func PrintWorkout() {
	for index, exercise := range GetWorkout() {
		fmt.Printf("%v. %s\n", index+1, exercise)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	PrintWorkout()
}
