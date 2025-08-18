class Solution {
public: 

    // O(N) + O(1)
    // prefix some question here we want to solve this 
    bool carPooling(vector<vector<int>>& trips, int capacity) {
        vector<int> a(1001, 0); 
        for (auto &tr : trips) {
            int pass = tr[0]; 
            int source = tr[1]; 
            int des = tr[2]; 

            a[source] += pass; 
            a[des] -= pass; 
        } 

        for (int i = 0; i < a.size(); i++) {
            capacity -= a[i]; 
            if (capacity < 0) {
                return false; 
            }
        }

        return true; 
    }
};

