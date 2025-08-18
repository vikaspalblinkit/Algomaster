#include<bits/stdc++.h>
using namespace std;


int findMajorityElements(int arr[], int n) {
    // step 1 : Find the frequency of each elements 
    map<int,int> frequency; 

    for (int i = 0; i < n; i++) {
        frequency[arr[i]]++; 
    }

    for (auto p : frequency) {
        if (p.second > n / 2) {
            return p.first; 
        }
    }

    return 0; 
}


// Voting Algorithm to find the majority elements
// Cool Algorithm to find the majority elements in the codes interview
int findCandidate(int arr[], int size) {
    int maj_index = 0, count = 1; 
    for (int i = 1; i < size; i++) {
        if (arr[maj_index] == arr[i]) {
            count++; 
        } else {
            count--; 
        }
        if (count == 0) {
            maj_index = i; 
            count = 1; 
        }
    } 
    return arr[maj_index];
}


// N/3 extension of the above problems we have solved. 

int main() {

    int arr[] = {1,3,3,1,2}; 
    int n = sizeof(arr)/sizeof(arr[0]);

    int ans = findCandidate(arr, n); 
    cout << "Majority Element " << ans; 


    return 0; 
}

/*
    Best Time Complexity : Voting Algorithm to select the One leaders 
    O(n), O(1) 



*/