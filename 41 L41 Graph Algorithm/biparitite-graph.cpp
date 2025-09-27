class Solution {
public: 

    bool bfs(vector<vector<int>>& graph, int node, int cur, vector<int>& color) { 
        queue<pair<int,int>> q; 
        q.push({node, cur}); 

        while (!q.empty()) {
            auto [node, cur] = q.front(); 
            q.pop(); 
            color[node] = cur; 
            for (auto it : graph[node]) {
                if (color[it] == -1) {
                    q.push({it, 1-cur}); 
                } else if (color[it] == cur) {
                    return false; 
                } 
            }
        }
        return true; 
    }
 
    bool isBipartite(vector<vector<int>>& graph) {
        int n = graph.size();   
        vector<int> color(n, -1); 
        for (int i = 0; i < n; i++) {
            if (color[i] == -1) {
                bool ans = bfs(graph, i, 0, color); 
                if (ans == false) {
                    return false; 
                }
            }
        }
        return true; 
    }
};