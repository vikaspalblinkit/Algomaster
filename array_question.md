## Array Questions: 
1. Sorts 0,1,2 

    low=mid=0, high = n-1 
    mid <= high 

    1. mid == 0 increase the low,mid -> swap low, mid 
    2. mid == 2 swap(mid, high) decrease high 
    3. mid == 1 mid++ increase  

2. Search in 2D Matrix 
 - optimal approach was to convert the arrays the 2d array into 1D arrays 
    row = n  -> 5/4 = 1  10/4 = 2 
    col = m  -> 5%4 = 1  10%4 = 2 
    index = 5  
```cpp    
interesting questions to ask 
bool searchMatrix(vector<vector<int>>& matrix, int target) {
    int n = matrix.size();
    int m = matrix[0].size();

    //apply binary search:
    int low = 0, high = n * m - 1;
    while (low <= high) {
        int mid = (low + high) / 2;
        int row = mid / m, col = mid % m;
        if (matrix[row][col] == target) return true;
        else if (matrix[row][col] < target) low = mid + 1;
        else high = mid - 1;
    }
    return false;
}
``` 

3. Kadane's Algo : Max Subarray sum in a arrays  
```cpp
long long maxSubarraySum(int arr[], int n) {
    long long maxi = LONG_MIN; // maximum sum
    long long sum = 0;

    for (int i = 0; i < n; i++) {

        sum += arr[i];

        if (sum > maxi) {
            maxi = sum;
        }

        // If sum < 0: discard the sum calculated
        if (sum < 0) {
            sum = 0;
        }
    }

    // To consider the sum of the empty subarray
    // uncomment the following check:
    //if (maxi < 0) maxi = 0;
    return maxi;
} 
```

4. Buy and Sell stocks I Easy 
1. Maintain the current_min = array[0] 
2. Just if 1 to 1+1 if you can make the profits and then summization that additions. 


5. Find the majority elements by using the voting algorithms:
```cpp 
int majorityElement(vector<int> v) {

    //size of the given array:
    int n = v.size();
    int cnt = 0; // count
    int el; // Element

    //applying the algorithm:
    for (int i = 0; i < n; i++) {
        if (cnt == 0) {
            cnt = 1;
            el = v[i];
        }
        else if (el == v[i]) cnt++;
        else cnt--;
    }

    //checking if the stored element
    // is the majority element:
    int cnt1 = 0;
    for (int i = 0; i < n; i++) {
        if (v[i] == el) cnt1++;
    }

    if (cnt1 > (n / 2)) return el;
    return -1;
}
``` 
- Very Intitute tho nahi hai need to learn this. 
Use the teaching methods to learn anything you want to see. 



Right Side Views: 
```cpp
#include <iostream>
#include <vector>
#include <set>
#include <queue>
#include <map>

using namespace std;

// Node structure for the binary tree
struct Node {
    int data;
    Node* left;
    Node* right;
    // Constructor to initialize
    // the node with a value
    Node(int val) : data(val),
        left(nullptr), right(nullptr) {}
};

class Solution {
public:
    // Function to return the
    // Right view of the binary tree
    vector<int> rightsideView(Node* root) {
        // Vector to store
        // the result
        vector<int> res;

        // Get the level order
        // traversal of the tree
        vector<vector<int>> levelTraversal = levelOrder(root);

        // Iterate through each level and
        // add the last element to the result
        for (auto level : levelTraversal) {
            res.push_back(level.back());
        }

        return res;
    }

    // Function to return the
    // Left view of the binary tree
    vector<int> leftsideView(Node* root) {
        // Vector to store the result
        vector<int> res;

        // Get the level order
        // traversal of the tree
        vector<vector<int>> levelTraversal = levelOrder(root);

        // Iterate through each level and
        // add the first element to the result
        for (auto level : levelTraversal) {
            res.push_back(level.front());
        }

        return res;
    }

private:
    // Function that returns the
    // level order traversal of a Binary tree 
    vector<vector<int>> levelOrder(Node* root) {
        vector<vector<int>> ans;

        // Return an empty vector
        // if the tree is empty
        if (!root) {
            return ans;
        }

        // Use a queue to perform
        // level order traversal
        queue<Node*> q;
        q.push(root);

        while (!q.empty()) {
            int size = q.size();
            vector<int> level;

            // Process each node
            // in the current level
            for (int i = 0; i < size; i++) {
                Node* top = q.front();
                level.push_back(top->data);
                q.pop();

                // Enqueue the left
                // child if it exists
                if (top->left != NULL) {
                    q.push(top->left);
                }

                // Enqueue the right
                // child if it exists
                if (top->right != NULL) {
                    q.push(top->right);
                }
            }

            // Add the current
            // level to the result
            ans.push_back(level);
        }

        return ans;
    }
};




int main() {
    // Creating a sample binary tree
    Node* root = new Node(1);
    root->left = new Node(2);
    root->left->left = new Node(4);
    root->left->right = new Node(10);
    root->left->left->right = new Node(5);
    root->left->left->right->right = new Node(6);
    root->right = new Node(3);
    root->right->right = new Node(10);
    root->right->left = new Node(9);

    Solution solution;

        // Get the Right View traversal
    vector<int> rightView = solution.rightsideView(root);

    // Print the result for Right View
    cout << "Right View Traversal: ";
    for(auto node: rightView){
        cout << node << " ";
    }
    cout << endl;

    // Get the Left View traversal
    vector<int> leftView = solution.leftsideView(root);

    // Print the result for Left View
    cout << "Left View Traversal: ";
    for(auto node: leftView){
        cout << node << " ";
    }
    cout << endl;

    return 0;
} 
```

### Check for Symmetrical Binary Tree
- Simply Problem hai 

### Check if the Binary Tree is Balanced Binary Tree 
- 

### Check for the maximum path sum 
1. Left 
2. right 
3. Root se hooga 
```cpp
#include <iostream>
#include <algorithm>
#include <climits>

using namespace std;

// Node structure for the binary tree
struct Node {
    int data;
    Node* left;
    Node* right;
    // Constructor to initialize
    // the node with a value
    Node(int val) : data(val), left(nullptr), right(nullptr) {}
};

class Solution {
public:
    // Recursive function to find the maximum path sum
    // or a given subtree rooted at 'root'
    // The variable 'maxi' is a reference parameter
    // updated to store the maximum path sum encountered
    int findMaxPathSum(Node* root, int &maxi) {
        // Base case: If the current node is null, return 0
        if (root == nullptr) {
            return 0;
        }

        // Calculate the maximum path sum
        // for the left and right subtrees
        int leftMaxPath = max(0, findMaxPathSum(root->left, maxi));
        int rightMaxPath = max(0, findMaxPathSum(root->right, maxi));

        // Update the overall maximum
        // path sum including the current node
        maxi = max(maxi, leftMaxPath + rightMaxPath + root->data);

        // Return the maximum sum considering
        // only one branch (either left or right)
        // along with the current node
        return max(leftMaxPath, rightMaxPath) + root->data;
    }

    // Function to find the maximum
    // path sum for the entire binary tree
    int maxPathSum(Node* root) {
        // Initialize maxi to the
        // minimum possible integer value
        int maxi = INT_MIN; 
         // Call the recursive function to
         // find the maximum path sum
        findMaxPathSum(root, maxi);
        // Return the final maximum path sum
        return maxi; 
    }
};


int main() {
    // Creating a sample binary tree
    Node* root = new Node(1);
    root->left = new Node(2);
    root->right = new Node(3);
    root->left->left = new Node(4);
    root->left->right = new Node(5);
    root->left->right->right = new Node(6);
    root->left->right->right->right = new Node(7);

    // Creating an instance of the Solution class
    Solution solution;

    // Finding and printing the maximum path sum
    int maxPathSum = solution.maxPathSum(root);
    cout << "Maximum Path Sum: " << maxPathSum << endl;


    return 0;
}
```

### ZigZag Traversal 
```cpp 
   vector<vector<int>> ZigZagLevelOrder(Node* root){
        // Vector to store the
        // result of zigzag traversal
        vector<vector<int>> result;
        
        // Check if the root is NULL,
        // return an empty result
        if(root == NULL){
            return result;
        }
        
        // Queue to perform
        // level order traversal
        queue<Node*> nodesQueue;
        nodesQueue.push(root);
        
        // Flag to determine the direction of
        // traversal (left to right or right to left)
        bool leftToRight = true;
        
        // Continue traversal until
        // the queue is empty
        while(!nodesQueue.empty()){
            // Get the number of nodes
            // at the current level
            int size = nodesQueue.size();
            
            // Vector to store the values
            // of nodes at the current level
            vector<int> row(size);
            
            // Traverse nodes at 
            // the current level
            for(int i = 0; i < size; i++){
                // Get the front node
                // from the queue
                Node* node = nodesQueue.front();
                nodesQueue.pop();
                
                // Determine the index to insert the node's
                // value based on the traversal direction
                int index = leftToRight ? i : (size - 1 - i);
                
                // Insert the node's value at
                // the determined index
                row[index] = node->data;
                
                // Enqueue the left and right
                // children if they exist
                if(node->left){
                    nodesQueue.push(node->left);
                }
                if(node->right){
                    nodesQueue.push(node->right);
                }
            }
            
            // Switch the traversal
            // direction for the next level
            leftToRight = !leftToRight;
            
            // Add the current level's
            // values to the result vector
            result.push_back(row);
        }
        
        // Return the final result of
        // zigzag level order traversal
        return result;
    }
```