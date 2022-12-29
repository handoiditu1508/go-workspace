package main

import (
	"log"
	"net/http"
	"testmodule/muxsample"

	"github.com/gorilla/mux"
)

// func main() {
// 	// fmt.Println(smallestAppearTimes.InvokesmallestAppearTimes1("00-44  48 5555 8361"))
// 	// fmt.Println(exercise2.InvokeExercise1("0 - 22 1985--324"))
// 	// fmt.Println(exercise2.InvokeExercise1("555372654"))

// 	// fmt.Println(exercise2.InvokeExercise2([]int{3, 4, 5, 4}))
// 	// fmt.Println(exercise2.InvokeExercise2([]int{5, 5, 2, 3, 4}))
// 	// fmt.Println(exercise2.InvokeExercise2([]int{1, 2, 3, 4, 5, 6, 7}))

// 	// fmt.Println(exercise2.InvokeExercise39([]int{1, 3, 2, 4}, []int{4, 1, 3, 2}))       // true
// 	// fmt.Println(exercise2.InvokeExercise39([]int{1, 2, 3, 4}, []int{2, 1, 4, 3}))       // false
// 	// fmt.Println(exercise2.InvokeExercise39([]int{3, 1, 2}, []int{2, 3, 1}))             // true
// 	// fmt.Println(exercise2.InvokeExercise39([]int{1, 2, 1}, []int{2, 3, 3}))             // false
// 	// fmt.Println(exercise2.InvokeExercise39([]int{1, 2, 3, 4}, []int{2, 1, 4, 4}))       // false
// 	// fmt.Println(exercise2.InvokeExercise39([]int{1, 2, 2, 3, 3}, []int{2, 3, 3, 4, 5})) // false
// 	// fmt.Println(exercise2.InvokeExercise39([]int{1, 3, 4, 2}, []int{3, 4, 2, 3}))       // false

// 	// fmt.Println(exercise2.InvokeExercise44(5, 8))
// 	// fmt.Println(exercise2.InvokeExercise44(4, 10))
// 	// fmt.Println(exercise2.InvokeExercise44(1, 2))
// 	// fmt.Println(exercise2.InvokeExercise44(5, 11))
// 	// fmt.Println(exercise2.InvokeExercise44(10, 44))
// 	// fmt.Println(exercise2.InvokeExercise44(50, 1000))
// 	// fmt.Println(exercise2.InvokeExercise44(50, 567))

// 	// fmt.Println(exercise2.InvokeExercise45("abaaba"))
// 	// fmt.Println(exercise2.InvokeExercise45("zyzyzyz"))
// 	// fmt.Println(exercise2.InvokeExercise45("aabbbabaaa"))
// 	// fmt.Println(exercise2.InvokeExercise45("kqkkqqqqkkkkqkqqkqkqq"))
// 	// fmt.Println(exercise2.InvokeExercise45("kkqq"))

// 	// fmt.Println(exercise2.InvokeExercise3("Sun 10:00-20:00\n" + "Fri 05:00-10:00\n" + "Fri 16:30-23:50\n" + "Sat 10:00-24:00\n" + "Sun 01:00-04:00\n" + "Sat 02:00-06:00\n" + "Tue 03:30-18:15\n" + "Tue 19:00-20:00\n" + "Wed 04:25-15:14\n" + "Wed 15:14-22:40\n" + "Thu 00:00-23:59\n" + "Mon 05:00-13:00\n" + "Mon 15:00-21:00"))
// 	// fmt.Println(exercise2.InvokeExercise3("Mon 01:00-23:00\n" + "Tue 01:00-23:00\n" + "Wed 01:00-23:00\n" + "Thu 01:00-23:00\n" + "Fri 01:00-23:00\n" + "Sat 01:00-23:00\n" + "Sun 01:00-21:00"))

// 	// fmt.Println(exercise2.InvokeExercise4([]string{"......", "......", "...X..", "......", "...X..", ".XO..."}))
// 	// fmt.Println(exercise2.InvokeExercise4([]string{"..X", ".X.", "O.."}))
// 	// fmt.Println(exercise2.InvokeExercise4([]string{"...", "..X", ".O."}))
// 	// fmt.Println(exercise2.InvokeExercise4([]string{"...", ".X.", ".O."}))
// 	// fmt.Println(exercise2.InvokeExercise4([]string{"O..", ".X.", "..."}))
// 	// fmt.Println(exercise2.InvokeExercise4([]string{"..X...", "......", "....X.", ".X....", "..X.X.", "...O.."}))
// 	// fmt.Println(exercise2.InvokeExercise4([]string{"X....", ".X...", "..O..", "...X.", "....."}))

// 	// fmt.Println(exercise2.InvokeExercise5("my.song.mp3 11b\n" + "greatSong.flac 1000b\n" + "not3.txt 5b\n" + "video.mp4 200b\n" + "game.exe 100b\n" + "mov!e.mkv 10000b"))

// 	// fmt.Println(exercise2.InvokePalindrome("ervervige"))
// 	// fmt.Println(exercise2.InvokePalindrome("aaabab"))
// 	// fmt.Println(exercise2.InvokePalindrome("x"))

// 	// fmt.Println(exercise2.InvokeNotification("And now here is my secret", 15))
// 	// fmt.Println(exercise2.InvokeNotification("There is an animal with four legs", 15))
// 	// fmt.Println(exercise2.InvokeNotification("super dog", 4))
// 	fmt.Println(exercise2.InvokeNotification("Hello world", 9))
// }

// func main() {
// 	sample.RunChannelSample()
// }

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", muxsample.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", muxsample.GetUser).Methods("GET")
	r.HandleFunc("/users", muxsample.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", muxsample.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", muxsample.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	muxsample.InitialMigration()
	initializeRouter()
}
