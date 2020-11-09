class Solution {
public:
	vector<vector<int>> kClosest(vector<vector<int>>& points, int K) {
		sort(points.begin(), points.end(), [](vector<int>& x, vector<int>& y) {
			return x[0] * x[0] + x[1] * x[1] < y[0] * y[0] + y[1] * y[1];
			});
		return { points.begin(), points.begin() + K };
	}
};