class Solution {
public:
	int numEquivDominoPairs(vector<vector<int>>& dominoes) {
		vector<int> nums(100);
		int num, res = 0;
		for (int i = 0; i < dominoes.size(); i ++) {
			num = dominoes[i][0] > dominoes[i][1] ? dominoes[i][0] * 10 + dominoes[i][1] : dominoes[i][1] * 10 + dominoes[i][0];
			nums[num] ++;
			res += (nums[num] - 1);
		}
		return res;
	}
};