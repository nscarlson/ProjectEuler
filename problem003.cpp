#include <iostream>
#include <cmath>

#define NUMBER 600851475143

using namespace std;

bool isPrime(int n)
{
    int limit = (int) ceil(sqrt(n));
    
    for (int i = 2; i <= limit; i++) {
        
        if (n % i == 0) {
            return false;
        }
    }
    cout << n << " is prime!" << endl;
    return true;
}

int main() {

    int largest = 0;
    int limit = (int) ceil(sqrt(NUMBER));
    
    for (int i = 1; i < limit; i++) {
        if (NUMBER % i == 0 && isPrime(i) && i > largest) {
            largest = i;
        }
    }
    
    cout << "The largest prime factor of 600851475143 is: " << largest << endl;
}