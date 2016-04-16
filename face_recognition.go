package main
// #cgo CPPFLAGS: -IC:/libs/opencv/build/include
// #cgo LDFLAGS:-lstdc++
// #include "face_recognition.hpp"
import "C"

func getLabel(lb int) int {
    return int(C.get_label(C.int(lb)))
}