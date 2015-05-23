#include <iostream>

using namespace std;

int fibonacci(int n) {
    if (n == 1) {
        return 1;
    }
    
    if (n == 2) {
        return 2;
    }

    return fibonacci(n - 2) + fibonacci(n - 1);
}

int main() {
    int sum;
    int current_fibonacci;
    
    for (int i = 1; fibonacci(i) <= 4000000; i++) {
        current_fibonacci = fibonacci(i);
        if (current_fibonacci % 2 == 0) {
            sum += current_fibonacci;
        }
    }
    cout << "The sum of all even fibonacci #s less than 4000000 is: " << sum << endl;
}