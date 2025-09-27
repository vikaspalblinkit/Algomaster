// Given  a set of wines your task it to calculate the maximum profits you can make by selling on the wines 
// 1. You can sell only the first and last wines and each year the wines prices is going to increase 

#include<bits/stdc++.h>
using namespace std; 

int dp[100][100]; 


// recursive solutions 
int maxProfit(int price[], int start, int end, int y) {

    // Base Case 
    if(start > end) {
        return 0;
    }

    // recursive case 
    int opt1 = price[start]*y + maxProfit(price, start+1, end, y+1);   // include the start as first options and call on the smaller subproblems  
    int opt2 = price[end]*y + maxProfit(price, start, end-1, y+1);   // include the last bottle and call on the smaller subproblems 

    return max(opt1, opt2); 
}



int main() {

    memset(dp, -1, sizeof(dp))
    int price[] = {2,3,5,1,4};
    int n = sizeof(price)/sizeof(int); 

    cout << maxProfit(price, 0, n-1, 1);

    return 0;
}