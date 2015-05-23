#include <iostream>
#include <cmath>

#define LIMIT 2000000

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

	long long sum = 2;

	for (int i = 3; i < LIMIT; i += 2) {
		if (isPrime(i)) {
			sum += i;
		}
	}

	cout << "The sum of all primes below 2,000,000 is: " << sum << endl;
}