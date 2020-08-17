package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var loc *time.Location

func setTimezone(tz string) error {
	location, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}
	loc = location
	return nil
}

func getTime(t time.Time) time.Time {
	return t.In(loc)
}

func hello(w http.ResponseWriter, req *http.Request) {
	timezone := os.Getenv("TIMEZONE")
	if err := setTimezone(timezone); err != nil {
		log.Fatal(err) // most likely timezone not loaded in Docker OS
	}
	t := getTime(time.Now())
	fmt.Fprintln(w, t.String())
	fmt.Fprintln(w, "TIMEZONE: "+timezone)
}

func unix(w http.ResponseWriter, req *http.Request) {
	timezone := os.Getenv("TIMEZONE")
	if err := setTimezone(timezone); err != nil {
		log.Fatal(err)
	}
	t := getTime(time.Now()).Format(time.UnixDate)
	fmt.Fprintln(w, t)
	fmt.Fprintln(w, "TIMEZONE: "+timezone)
}

func kitchen(w http.ResponseWriter, req *http.Request) {
	timezone := os.Getenv("TIMEZONE")
	if err := setTimezone(timezone); err != nil {
		log.Fatal(err)
	}
	t := getTime(time.Now()).Format(time.Kitchen)
	fmt.Fprintln(w, t)
	fmt.Fprintln(w, "TIMEZONE: "+timezone)
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/kitchen", kitchen)
	http.HandleFunc("/unix", unix)

	log.Fatal(http.ListenAndServe(":8083", nil))
}
