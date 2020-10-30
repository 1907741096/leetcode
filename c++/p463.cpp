#include <queue>
struct point {
	int x, y;
};
class Solution {
public:
	int dir[4][2] = { {1, 0} ,{-1, 0},{0, 1},{0, -1} };
	int islandPerimeter(std::vector<std::vector<int>>& grid) {
		std::queue<point> q;
		int num = 0;
		vector<vector<bool>> flagList(grid.size());
		for (int i = 0; i < grid.size(); i++) {
			flagList[i] = vector<bool>(grid[i].size());
		}
		for (int i = 0; i < grid.size(); i++) {
			for (int j = 0; j < grid[i].size(); j++) {
				if (grid[i][j] == 1) {
					flagList[i][j] = true;
					q.push(point{ i, j });
					break;
				}
			}
			if (!q.empty()) break;
		}

		while (!q.empty()) {
			point p = q.front();
			q.pop();
			for (int i = 0; i < 4; i ++) {
				int x = p.x + dir[i][0], y = p.y + dir[i][1];
				if (x >= 0 && x < grid.size() && y >= 0 && y < grid[0].size() && grid[x][y] == 1) {
					if (!flagList[x][y]) {
						q.push(point{ x, y });
						flagList[x][y] = true;
					}
				}
				else {
					num++;
				}
			}
		}

		return num;
	}
};