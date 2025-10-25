March 8 
1. Longest Increasing subsequences 
2. Max Rectangle in Binary Matrix 
3. Distinct Subsequences. 
4. Unique path in a Grid 
5. Max Product Subarrays 



1. longest bitonic Subsequences Algo 
2. Climbing Stairs 
3. Smart Robber
4. Number of paths 
5. Longest Valid Parenthesis 
6. Maximum rectangles of 1's 
7. Buys and sells stocks 
 

 decodes ways 
 word breaks 


```cpp
 int minStep(int n, int dp[]) {

	if (n == 1) return 0;

	// check if exists in cache
	if (dp[n] > 0) {
		return dp[n];
	}

	int op1 = minStep(n - 1, dp);
	int min_step = op1;

	if (n % 2 == 0) {
		int op2 = minStep( n / 2, dp);
		if (op2 < min_step) {
			min_step = op2;
		}
	}

	if (n % 3 == 0) {
		int op3 = minStep( n / 3, dp);
		if (op3 < min_step) {
			min_step = op3;
		}
	}

	dp[n] = 1 + min_step;
	return dp[n];
}
``` 

/**

	10-12-30 

	unordered_map<int,pair<list<int>:: iterator, int>>  um; 
	


 */