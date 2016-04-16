#include "face_recognition.hpp"
#include <opencv2/highgui/highgui.hpp>
#include "opencv2/contrib/contrib.hpp"
#include <opencv2/core/core.hpp>
#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <string>

using namespace cv;
using namespace std;

std::string to_string(int val) {
	std::stringstream ss;
	ss << val;
	return ss.str();
}

static void read_csv(const string& filename, vector<Mat>& images, vector<int>& labels, char separator = ';') {
	std::ifstream file(filename.c_str(), ifstream::in);
	if (!file) {
		string error_message = "No valid input file was given, please check the given filename.";
		CV_Error(CV_StsBadArg, error_message);
	}
	string line, path, classlabel;
	while (getline(file, line)) {
		stringstream liness(line);
		getline(liness, path, separator);
		getline(liness, classlabel);
		if(!path.empty() && !classlabel.empty()) {
			images.push_back(imread(path, 0));
			labels.push_back(atoi(classlabel.c_str()));
		}
	}
}

int get_label(int id) {
	std::string id_str = to_string(id);
	
	// holds images and labels
	vector<Mat> images;
	vector<int> labels;
	// images for first person
	images.push_back(imread("./asset/"+id_str+"/user_farss.png", CV_LOAD_IMAGE_GRAYSCALE)); labels.push_back(0);
	
	Ptr<FaceRecognizer> model =  createFisherFaceRecognizer();
	model->train(images, labels);
}
