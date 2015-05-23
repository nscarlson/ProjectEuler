#include <iostream>

#define LOWER_LIM 100
#define UPPER_LIM 999

using namespace std;

bool isPalindrome(int n) {
    int num = n;
    int reverse = 0;
    
    while (num > 0) {
        
        int digit = num % 10;
        reverse = reverse * 10 + digit;
        
        num = num / 10;
        
    }
    
    return reverse == n ? true : false;
}

int main() {
    
    int largest = 0;
    
    for (int i = LOWER_LIM; i <= UPPER_LIM; i++) {
        for (int j = i; j <= UPPER_LIM; j++) {
            int product = i * j;
            
            if (isPalindrome(product) && product > largest) {
                largest = product;
            }
        }
    }
    
    cout << "The largest palindrome made from the product of two 3-digit numbers is: ";
    cout << largest << endl;
}