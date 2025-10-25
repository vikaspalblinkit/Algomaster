/*
  Methods Vectors that most commonly used
  1. push_back()--> Double the memory each tines
  2. pop_back()
  3. insert(d.begin()+3)
  4. erase(d.begin()+3)
  5. resize()
  6. clear() clear the elements of the arrays means the vectors is empty is
  7. empty()
  8. front() return the first elements of the vectors
  9. back() last elements of the vectors


** Reserve --> avoid the capacity of memory that is increasing in push_back()
  notes :
  1. Shrink the elements

*/

/*
 Solving a Car problems which based on sorting of the location which is near cab at the origin of me
 Two methods : 1. vector pair
         2. custom user defined class

*/
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

// custom comparator for Sorting the pairs
bool compare(pair <int, int> v1, pair <int, int> v2) // Accepting the pair which is used for the comparision
{
  // Calculate the Distance from the origin
  int d1 = v1.first * v1.first + v1.second * v1.second;
  int d2 = v2.first * v2.first + v2.second * v2.second;
  if (d1 == d2) { // if both distances is equal then return which less corrdinates is less
    return v1.first < v2.first;
  }
  return d1 < d2;
}


int main() {
  // five x,y corrdinates of that car
  // vector<pair<int, int>> P1;
  // int n;
  // cin >> n;
  // for (int i = 0; i < n; i++) {
  //   int x, y;
  //   cin >> x >> y;
  //   P1.push_back(make_pair(x, y));
  // }
  // sort(P1.begin(), P1.end(), compare);
  // //for each loopb
  // for (auto x : P1) {
  //   cout << x.first << " " << x.second << endl;
  // } 

  vector<int> v {1, 1, 2, 4, 4, 5, 5, 6, 6}; 
  // sort(v.begin(), v.end()); 
  int ans = 0; 
  // this is the first elements we are trying to incomplete once
  for(int i = 0; i < v.size(); i++) {
      ans = ans ^ v[i]; 
  }  

  // Important solution using the binary search approach 

  cout << ans; 
  return 0;
}

void removeChar(char* s, char c)
{
 
    int j, n = strlen(s);
    for (int i = j = 0; i < n; i++)
        if (s[i] != c)
            s[j++] = s[i];
 
    s[j] = '\0'; // end of the string representation of the string here. 
} 



/**
 1. Leadership Principle  for amazon and LLD Patterns 
 2. 
 * 
 */
 
/* 
 1. Binary Search on Koko Eating Banana 
 2. Answers on the binary searchs 
 */