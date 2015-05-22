#include <iostream>

#define NUMBER 100

using namespace std;

int sum_of_squares(int n) {
    return n * (n + 1) * (2 * n + 1) / 6;
}

int square_of_sum(int n) {
    return (n * (n + 1) / 2) * (n * (n + 1) / 2);
}

int main() {
    int difference = square_of_sum(NUMBER) - sum_of_squares(NUMBER);
    cout << "Difference between the sum of squares and the square of the sum: ";
    cout <<  difference << endl;
}