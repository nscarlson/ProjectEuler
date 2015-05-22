#include <iostream>
#include <iomanip>
#include "./functions/functions.cpp"

#define GRID_SIZE 20
#define ll long long

using namespace std;
/*
 * There are essentially 2 ways of solving this problem.  First, the recursively-defined method,
 * which happens to be very inefficient, but for educational purposes I left it commented below.
 *
 * The second method is, in my opinion, a beautiful example of an analytical solution.  The problem 
 * can be shown as a combinatorics expression as simple as 2*GRID_SIZE choose GRID_SIZE.
 */
   

/*

ll S(ll i, ll j) {
	//cout << "Calculating: " << i << " " << j << endl;
	if (j == 0)
		return 1;
	if (0 < j && j < i)
		return S(i, j - 1) + S(i - 1, j);
	if (i == j)
		return 2 * S(i, i - 1);
}

*/

int main() {
	ll n = GRID_SIZE * 2;
	ll r = GRID_SIZE;

	cout << fixed << setprecision(0) << factorial(GRID_SIZE * 2) / (factorial(GRID_SIZE) * factorial(GRID_SIZE)) << endl;
}