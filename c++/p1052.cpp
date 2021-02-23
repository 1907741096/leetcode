class Solution {
public:
	int maxSatisfied(vector<int>& customers, vector<int>& grumpy, int X) {
		int max = 0, sum = 0;
		for (int i = 0; i < customers.size(); i++) {
			if (grumpy[i] == 0) {
				sum += customers[i];
			}
		}
		for (int i = 0; i < X; i++) {
			if (grumpy[i] == 1) {
				sum += customers[i];
			}
		}
		max = sum;
		for (int i = 0; i < customers.size() - X; i++) {
			if (grumpy[i] == 1) {
				sum -= customers[i];
			}
			if (grumpy[i + X] == 1) {
				sum += customers[i + X];
			}
			if (sum > max) {
				max = sum;
			}
		}
		return max;
	}
};