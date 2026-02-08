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