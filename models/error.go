package models

import "log"

func ErrorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}