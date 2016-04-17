#include "face_recognition.hpp"
#include <opencv2/highgui/highgui.hpp>
#include <opencv2/face.hpp>
#include <opencv2/core/core.hpp>
#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <string>

using namespace cv;
using namespace std;
using namespace face;

std::string to_string(int val) {
	std::stringstream ss;
	ss << val;
	return ss.str();
}

Ptr<FaceRecognizer> model;
vector<Mat> images;
vector<int> labels;

void init_model() {
	model = createLBPHFaceRecognizer();
}

void update_model(int id, char* file) {
	std::string id_str(file);
    cout << id_str << endl;	
    // images for first person
	images.push_back(imread(id_str, CV_LOAD_IMAGE_GRAYSCALE)); labels.push_back(id);
	
	model->train(images, labels); 
}

int get_label(char* file) {
    cout << file << endl;
	int label;
	float confidence;
	
	model->predict(imread(file, CV_LOAD_IMAGE_GRAYSCALE), label, confidence);
	cout << "Label: " << label << " Confidence: " << confidence << endl;
	if(confidence < 0.5) {
		label = -1;
	}
	return label;
}
