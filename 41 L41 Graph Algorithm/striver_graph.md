## Connected components: 
```cpp 
class Solution {
public:
    // create the adjacent list of the arrays  
    void dfs(int node, vector<vector<int>>& adjList, vector<bool> &visited) {
        visited[node] = true;
        for (int neighbor : adjList[node]) {
            if (!visited[neighbor]) {
                dfs(neighbor, adjList, visited);
            }
        }
    }

    int findCircleNum(vector<vector<int>>& isConnected) {
        int n = isConnected.size();
        vector<vector<int>> adjList(n);
        
        // Build the adjacency list
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (isConnected[i][j] == 1 && i != j) {
                    adjList[i].push_back(j);
                }
            }
        }
        
        vector<bool> visited(n, false);
        int provinces = 0;

        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                provinces++;
                dfs(i, adjList, visited);
            }
        }
        return provinces; 
    }
};
``` 

## Biparitite Graph 
```cpp 
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
                if (color[it] == -1) q.push({it, 1-cur}); 
                else if (color[it] == cur) return false; 
            }
        }
        return true; 
    }
 
    bool isBipartite(vector<vector<int>>& graph) {
        int n = graph.size();   
        vector<int> color(n, -1); 
        for (int i = 0; i < n; i++) {
            if (color[i] == -1) {
                if(!bfs(graph, i, 0, color)) return false; 
            }
        }
        return true; 
    }
};
``` 

## 

```cpp
class Solution {
public:
    void dfs(int row,int col, vector<vector<int>> &ans,int ini,vector<vector<int>>& image,int color){
        ans[row][col] = color;
        int n = image.size();
        int m = image[0].size();

        int dx[] = {-1,0,1,0};
        int dy[] = {0,1,0,-1};

        for(int k = 0;k < 4;k++){
            int nrow = row + dx[k];
            int ncol = col + dy[k];
            if(nrow >= 0 && nrow < n && ncol >= 0 && ncol < m && image[nrow][ncol] == ini && ans[nrow][ncol] != color){
                dfs(nrow,ncol,ans,ini,image,color);
            }
        }
    }

    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int color) {
        int intial = image[sr][sc];
        vector<vector<int>> ans = image;
        dfs(sr,sc,ans,intial,image,color);
        return ans;
    }
}; 
```

```cpp 
class Solution {
public:
	void solve(vector<vector<char>>& board) {

		int m = board.size(); // no of rows
		int n = board[0].size(); // no of cols

		// first row
		for (int j = 0; j < n; j++) {
			if (board[0][j] == 'O') {
				dfs(0, j, m, n, board);
			}
		}

		// last colums
		for (int i = 0; i < m; i++) {
			if (board[i][n - 1] == 'O') {
				dfs(i, n - 1, m, n, board);
			}
		}

		// first column
		for (int i = 0; i < m; i++) {
			if (board[i][0] == 'O') {
				dfs(i, 0, m, n, board);
			}
		}

		// last row
		for (int j = 0; j < n; j++) {
			if (board[m - 1][j] == 'O') {
				dfs(m - 1, j, m, n, board);
			}
		}

		for (int i = 0; i < m; i++) {
			for (int j = 0; j < n; j++) {
				if (board[i][j] == 'O') {
					board[i][j] = 'X';
				}
				else if (board[i][j] == '#') {
					board[i][j] = 'O';
				}
			}
		}

		return;
	}

	void dfs(int i, int j, int m, int n, vector<vector<char>> &board) {
		if (i < 0 || i == m || j < 0 || j == n || board[i][j] == 'X' || board[i][j] == '#') {
			return;
		}

		board[i][j] = '#';
		dfs(i - 1, j, m, n, board);
		dfs(i + 1, j, m, n, board);
		dfs(i, j - 1, m, n, board);
		dfs(i, j + 1, m, n, board);

		return;
	}
};
```


## Word ladder problems 

```cpp 
class Solution {
public:
	int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
		unordered_set<string> dict(wordList.begin(), wordList.end());
		queue<string> todo;
		todo.push(beginWord);
		int ladder = 1;
		while (!todo.empty()) {
			int n = todo.size();
			for (int i = 0; i < n; i++) {
				string word = todo.front();
				todo.pop();
				if (word == endWord) {
					return ladder;
				}
				dict.erase(word);
				for (int j = 0; j < word.size(); j++) {
					char c = word[j];
					for (int k = 0; k < 26; k++) {
						word[j] = 'a' + k;
						if (dict.find(word) != dict.end()) {
						`	todo.push(word);
						}
					}
					word[j] = c;
				}
			}
			ladder++;
		}
		return 0;
	}
};
```