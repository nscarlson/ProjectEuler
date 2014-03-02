#include <iostream>
#include <cmath>

#define NUMBER 600851475143

using namespace std;

bool isPrime(int n)
{
    int limit = (int) sqrt(n);
    if (n % 2 == 0) {
        return false;
    }
    
    for (int i = 3; i <= limit; i += 2) {
        if (n % i == 0) {
            return false;
        }
    }
    return true;
}

int main() {

    int largest = 0;
    int limit = (int) sqrt(NUMBER);
    
    for (int i = 1; i < limit; i++) {
        if (NUMBER % i == 0 && isPrime(i) && i > largest) {
            largest = i;
        }
    }
    
    cout << "The largest prime factor of 600851475143 is: " << largest << endl;
}