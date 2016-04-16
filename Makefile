.PHONY: clean

TARGET=face_recognition

$(TARGET): face_recognition.a
	go build -x .

face_recognition.a: face_recognition.o
	ar r $@ $^

%.o: %.cpp
	g++ -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
