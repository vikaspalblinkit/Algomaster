/*
	Find min steps to reach one:
		a.
		b. 


		Post Save methods 
		Pre Save  methods are
	
		dispatch module we have the receiver modules that we can calls // decators for the sames.

Time Complexity: O(n), as there will be n unique calls.
Space Complexity: O(n)
*/

#include<iostream>
using namespace std;

// Recursive approach :
int minStep(int n, int dp[]) {

	// Base Case 
	if (n == 1) return 0;

	// check if exists in cache and if that exists in the dp arrays 
	if (dp[n] > 0) {
		return dp[n];
	}

	// option 1 
	int op1 = minStep(n - 1, dp); // n-1 
	int min_step = op1;

	// option 2
	if (n % 2 == 0) {
		int op2 = minStep( n / 2, dp);
		if (op2 < min_step) {
			min_step = op2;
		}
	}

	// option 3 
	if (n % 3 == 0) {
		int op3 = minStep( n / 3, dp);
		if (op3 < min_step) {
			min_step = op3;
		}
	}

	// then add 1 to its  1 + current steps 
	dp[n] = 1 + min_step;
	return dp[n];
}


int minStepToOneRec(int n) {
    // Base case 
    if (n == 1) return 0; 

    // Recursive Case 
    int op1 = minStepToOneRec(n - 1); 
    int ans = op1; 

    if (n % 2 == 0) {
        int op2 = minStepToOneRec(n / 2);
        if (op2 < ans) {
            ans = op2; 
        }
    }

    if (n % 3 == 0) {
        int op3 = minStepToOneRec(n / 3);
        if (op3 < ans) {
            ans = op3; 
        }
    }
    return ans + 1; 
}

// this is questions user wants to solved
// Overlapping problem is currently have existing
int main() {
	int n = 3;
	int dp[n + 1];

	// Ensure that doesn't contains the garbage values
	for (int i = 0; i <= n; i++) {
		dp[i] = 0;
	}
	cout << minStep(n, dp);
	return 0;
}