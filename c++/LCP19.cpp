class Solution {
public:
	int minimumOperations(string leaves) {
		int n = leaves.size();
		int a = (leaves[0] == 'y');
		int b = INT_MAX;
		int c = INT_MAX;
		for (int i = 1; i < n; ++i) {
			int isRed = (leaves[i] == 'r');
			int isYellow = (leaves[i] == 'y');
			if (i >= 2) {
				c = min(b, c) + isYellow;
			}
			b = min(a, b) + isRed;
			a = a + isYellow;
		}
		return c;
	}
};