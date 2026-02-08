#include<bits/stdc++.h>
using namespace std;
/**
 * TC : O(N'2)
 * SC : O(N'2)
 */



int dp[100][100]; 

int maxProfit(int prices[], int start, int end, int year) {

    if (start > end) return 0; 
    if (start == end) return prices[start]*year; 

    if (dp[start][end] != -1) return dp[start][end];
    
    int left = prices[start] * year + maxProfit(prices, start+1, end, year+1);
    int right = prices[end] * year + maxProfit(prices, start, end-1, year+1);

    dp[start][end] = max(left, right);
    return dp[start][end];
}

int main() {
    memset(dp, -1, sizeof(dp));
    int prices[] = {2, 4, 6, 2, 5}; 
    cout << maxProfit(prices, 0, 4, 1) << endl; 
    return 0;
} 

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std; 

int main() {
    vector<vector<int>> intervals{{0,30}, {5,10}, {15,20}}; 
    sort(intervals.begin(), intervals.end());  

    bool ans = true; 

    for (int i = 1; i < intervals.size(); i++) {
        cout << intervals[i][0]  << " " <<  intervals[i-1][1] << endl; 
        if (intervals[i][0] < intervals[i-1][1]) {
            ans = false; 
            break; 
        }
    }    
    if (!ans) {
        cout << "NO" << endl; 
    } else {
        cout << "YES" << endl; 
    }

    return 0; 
}