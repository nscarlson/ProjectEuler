#include <iostream>

using namespace std;

int main() {
    int sum;
    for (int i = 1; i < 1000; i++) {
        if (i % 3 == 0 || i % 5 == 0) {
            sum += i;
        }
    }
    
    cout << "The sum of all the multiples of 3 or 5 below 100 is: ";
    cout << sum << endl;
}