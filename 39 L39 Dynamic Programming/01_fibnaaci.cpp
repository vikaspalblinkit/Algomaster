// Compute the Nth fibonaaci Number 
#include<iostream>
using namespace std; 


// Recursive Implementations  time : 0(2^n) space: 0(N)
int fib(int n) {

    // Base Case 
    if(n == 0 and n == 1) {
        return n; 
    }

    int ans; 
    ans = fib(n-1) + fib(n-2);
    return ans;
}


// Top Down approaches time : 0(N) , space : O(N)
int fib(int n, int dp[]) {
   
   // Base case 
   if(n == 0 and n == 1) {
       return n; 
   }

    // check the answer in the dp arrays 
    if(dp[n] != 0) {
        return dp[n];
    }
    int ans; 
    ans = fib(n-1, dp) + fib(n-2, dp);
    return dp[n]=ans; 
}



// Bottom Up DP time : O(N) , space : O(N) 
// climibing stairs will have the same problems 
// dp(i-1), dp(i-2) additions is the state of the ansewers 
int fibBU(int n) {
    
    int dp[100] = {0};
    dp[0] = 0;
    dp[1] = 1;

    for(int i=2; i<=n; i++) {
        dp[i] = dp[i-1] + dp[i-2];
    }

    return dp[n];
}



// Bottom up with space optimization 
int fibBUSpace(int n) {
    
    long long int  a = 0;
    int b = 1; 
    int c = 0; 

    for(int i=0; i<n; i++) {
        c = a + b; 
        b = c; 
        a = b; 
    }
 
    return c; 
}



int main() {

    int n; cin >> n; 
    int dp[100] = {0};

    cout << fib(n) << endl;

    cout << fib(n, dp);

    return 0;
}