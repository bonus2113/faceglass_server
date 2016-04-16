package main
// #cgo CPPFLAGS: -IC:/libs/opencv/build/include
// #cgo LDFLAGS: -L${SRCDIR} -lstdc++ -lface_recognition
// #include "face_recognition.hpp"
import "C"

func getLabel(lb int) int {
    return int(C.get_label(C.int(lb)))
}