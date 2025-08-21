package main

import (
	"fmt"
	"time"
)

type Workout interface{
	Duration()				time.Duration
	CaloriesBurned()	float64
	RecordStats()			
	GetType()					string
}

// CardioWorkout struct
type CardioWorkout struct {
	duration     time.Duration
	distance     float64
	avgHeartRate float64
}

func (c CardioWorkout) Duration() time.Duration {
	return c.duration
}

func (c CardioWorkout) CaloriesBurned() float64 {
	calories := c.Duration().Minutes() * 10 * (c.avgHeartRate / 100)
	return calories
}

func (c CardioWorkout) RecordStats() {
	fmt.Printf("\tDuration: %v\n\tDistance: %.2f\n\tHeart rate(avg): %.2f\n", c.duration, c.distance, c.avgHeartRate) 
}

func (c CardioWorkout) GetType() string {
	return "Cardio"
}

// StrengthWorkout struct
type StrengthWorkout struct {
	duration time.Duration
	weight   int
	reps     int
	sets     int
}

func (s StrengthWorkout) Duration() time.Duration {
	return s.duration
}

func (s StrengthWorkout) CaloriesBurned() float64 {
	calories := s.Duration().Minutes() * 5 * (float64(s.weight) / 10)
	return calories
}

func (s StrengthWorkout) RecordStats() {
	fmt.Printf("\tDuration: %v\n\tWeights used: %d\n\tReps: %d\n\tSets: %d\n", s.duration, s.weight, s.reps, s.sets)
}

func (s StrengthWorkout) GetType() string {
	return "Strength"
}

func summarizeWorkouts(workouts []Workout) {
	for _, workout := range workouts {
		fmt.Printf("Current Workout: %v\n\tCalories Burned: %.2f\nStats:\n", workout.GetType(),
																																			 workout.CaloriesBurned())
		workout.RecordStats()
	}
}

func main() {
	w := []Workout {
		CardioWorkout{
			duration: 30 * time.Minute,
			distance: 5,
			avgHeartRate: 106,
		},
		StrengthWorkout{
			duration: 30 * time.Minute,
			weight: 150,
			reps: 10,
			sets: 6,
		},
	}

	fmt.Println("Workout Summary:")
	summarizeWorkouts(w)
}