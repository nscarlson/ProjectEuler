#include <iostream>
#include <cmath>

#define nthPrime 10001

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

    int numberOfPrimes = 1;
    int currentPrime = 0;

    for (int i = 3; numberOfPrimes < nthPrime; i += 2) {
        if (isPrime(i)) {
            numberOfPrimes++;
            currentPrime = i;
        }
    }

    cout << "The 10,001st prime number is: " << currentPrime << endl;
}