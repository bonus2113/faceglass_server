package main
// #cgo pkg-config: opencv
// #cgo CPPFLAGS: -IC:/libs/opencv/build/include
// #cgo LDFLAGS: -lstdc++  
// #include "face_recognition.h"
import "C"
func initModel() {
    C.init_model();
}

func updateModel(id int, file string) {
    C.update_model(C.int(id), C.CString(file));
}

func getLabel(lb int) int {
    return int(C.get_label(C.int(lb)))
}
