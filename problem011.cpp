#include <iostream>
#include <sstream>
#include <fstream>
#include <string>
#include <ctime>

#define LINE_SIZE 4
#define GRID_SIZE 20

using namespace std;

int main() {
	const clock_t begin_time = clock();
	int grid[GRID_SIZE][GRID_SIZE];

	ifstream file("problem011grid.txt");
	int largestProduct = 0;
	string line;
	for (int i = 0; getline(file, line); i++)
	{
	  	istringstream iss(line);
		
		int value;
	  	for (int j = 0; iss >> value; j++)
	  	{
	  		if (value < 10) {
	  			cout << "0";
	  		}
	  		cout << value << " ";

	  		grid[j][i] = value;
	  	}
	  	cout << endl;
	}

	int product;

	for (int j = 0; j < GRID_SIZE; j++) {
		for (int i = 0; i < GRID_SIZE; i++) {
			
			// Check straight left
			if (i >= LINE_SIZE - 1) {
				product = 1;
				//cout << "Checking straight left: ";
				for (int c = 0; c < LINE_SIZE; c++) {
					//cout << grid[i - c][j] << endl;
					product *= grid[i - c][j];

					if (product > largestProduct)
						largestProduct = product;
				}
			}

			// Check straight right
			if (i <= GRID_SIZE - LINE_SIZE) {
				product = 1;
				for (int c = 0; c < LINE_SIZE; c++) {
					product *= grid[i + c][j];

					if (product > largestProduct)
						largestProduct = product;
				}
			}
	

			// Check straight down
			if (j <= GRID_SIZE - LINE_SIZE) {
				product = 1;
				for (int c = 0; c < LINE_SIZE; c++) {
					product *= grid[i][j + c];
						
					if (product > largestProduct)
						largestProduct = product;
				}
			}	

			// Check left diagonal
				if (i >= LINE_SIZE - 1 && j <= GRID_SIZE - LINE_SIZE) {
					product = 1;
					for (int c = 0; c < LINE_SIZE; c++) {
						product *= grid[i - c][j + c];

						if (product > largestProduct)
							largestProduct = product;
					}
				}

			// Check right diagonal
			if (i <= GRID_SIZE - LINE_SIZE && j <= GRID_SIZE - LINE_SIZE) {
				product = 1;
				for (int c = 0; c < LINE_SIZE; c++) {
					product *= grid[i + c][j + c];

					if (product > largestProduct)
						largestProduct = product;
				}
			}
		}
	}
	cout << largestProduct << endl;
	cout << float(clock() - begin_time) / CLOCKS_PER_SEC;
}