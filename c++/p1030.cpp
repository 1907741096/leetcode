class Solution {
public:
	vector<vector<int>> allCellsDistOrder(int R, int C, int r0, int c0) {
		vector<vector<int>> res;
		res.resize(R * C, vector<int>(2, 0));
		for (int i = 0; i < R; i++) {
			for (int j = 0; j < C; j++) {
				res[i * C + j][0] = i;
				res[i * C + j][1] = j;
			}
		}
		sort(res.begin(), res.end(), [&](vector<int>& a, vector<int>& b) {
			return abs(a[0] - r0) + abs(a[1] - c0) < abs(b[0] - r0) + abs(b[1] - c0);
			});
		return res;
	}
};