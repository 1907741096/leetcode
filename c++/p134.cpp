class Solution {
public:
	int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
		int res, i, j;
		for (i = 0; i < gas.size(); i++) {
			res = gas[i];
			for (j = i; j < i + gas.size(); j++) {
				res -= cost[j % cost.size()];
				if (res < 0) {
					i = j;
					break;
				}
				if ((j + 1) % gas.size() == i)  return i;
				res += gas[(j + 1) % gas.size()];
			}
		}
		return -1;
	}
};