// Time: 77.13%
// Memory: 40.24%

class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        // Сreating matrix for storing which tiles are visited
        vector<vector<bool>> checked(grid.size(), vector<bool>(grid[0].size(), false));
        int islandCount = 0;

        // Checking every tile
        for (int i = 0; i < checked.size(); i++) {
            for (int j = 0; j < checked[i].size(); j++) {
                if (grid[i][j] == '1' && !checked[i][j]) { // if the tile is ground and it is not visited then it is new island
                    checkIsland(grid, checked, i, j); //  Marking every tile of the island as visited
                    islandCount++;
                }
            }
        }
        return islandCount;
    }
    void checkIsland(vector<vector<char>>& grid, vector<vector<bool>> &checked, int i, int j) {
        checked[i][j] = true;
        if (j > 0 && grid[i][j - 1] == '1' && !checked[i][j - 1])
            checkIsland(grid, checked, i, j - 1);
        if (j < grid[i].size() - 1 && grid[i][j + 1] == '1' && !checked[i][j + 1])
            checkIsland(grid, checked, i, j + 1);
        if (i < grid.size() - 1 && grid[i + 1][j] == '1' && !checked[i + 1][j])
            checkIsland(grid, checked, i + 1, j);
        if (i > 0 && grid[i - 1][j] == '1' && !checked[i - 1][j])
            checkIsland(grid, checked, i - 1, j);
    }
};