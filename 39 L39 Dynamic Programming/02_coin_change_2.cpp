// to_string --> string to convert into string values here;
// here you are able to find out overlapping subproblems that need to solve first
// So using hashmaps to solve this kinds of problems.

#include<bits/stdc++.h>
using namespace std; 

int main() {
    unordered_map<string , int> mp; 
    string key; 
    key =  to_string(111) + "sdfksjd"; 
    mp[key] = 33; 

    cout << mp[key] << endl; 

    return 0; 
}   



// Google, Amazon Interview Questions put 
// Problem Questions: Asks 
// 1. what is constaints of N, M < 1000 
// Do we need to change the row and col that we have set to zeroes ? 
/**
 *  Create an aux arrays and update the data in that matrix Initialize all 1's
 *  Traverse the orignal arrays. 
 *  And make the change in aux arrays 
 *  O(N*M*(N+M)), O(N*M)
 * row[i] == true or col[j] == true 
 * a[i][j] = 0 
 * Can we reduce the space for same ? 
 */

#include <bits/stdc++.h>
using namespace std;

// Function to count the number of ways to make change for a given target sum
long countWaysToMakeChangeUtil(vector<int>& arr, int ind, int T, vector<vector<long>>& dp) {
    // Base case: if we're at the first element
    if (ind == 0) {
        // Check if the target sum is divisible by the first element
        return (T % arr[0] == 0);
    }
    
    // If the result for this index and target sum is already calculated, return it
    if (dp[ind][T] != -1)
        return dp[ind][T];
        

    // Calculate the number of ways without taking the current element
    long notTaken = countWaysToMakeChangeUtil(arr, ind - 1, T, dp);
    
    // Calculate the number of ways by taking the current element
    long taken = 0;
    if (arr[ind] <= T)
        taken = countWaysToMakeChangeUtil(arr, ind, T - arr[ind], dp);
        
    // Store the sum of ways in the DP table and return it
    return dp[ind][T] = notTaken + taken;
}

// Function to count the number of ways to make change for the target sum
long countWaysToMakeChange(vector<int>& arr, int n, int T) {
    vector<vector<long>> dp(n, vector<long>(T + 1, -1)); // Create a DP table
    
    // Call the utility function to calculate the answer
    return countWaysToMakeChangeUtil(arr, n - 1, T, dp);
}

int main() {
    vector<int> arr = {1, 2, 3};
    int target = 4;
    int n = arr.size();
    
    cout << "The total number of ways is " << countWaysToMakeChange(arr, n, target) << endl;

    return 0; // Return 0 to indicate successful program execution
}

// N*T 