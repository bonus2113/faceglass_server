package main
// #cgo pkg-config: opencv
// #cgo CPPFLAGS: -IC:/libs/opencv/build/include
// #cgo LDFLAGS: -lstdc++  
// #include "face_recognition.h"
import "C"

func getLabel(lb int) int {
    return int(C.get_label(C.int(lb)))
}
