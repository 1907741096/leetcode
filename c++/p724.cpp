class Solution {
public:
	int pivotIndex(vector<int>& nums) {
		int total = accumulate(nums.begin(), nums.end(), 0);
		int sum = 0;
		for (int i = 0; i < nums.size(); ++i) {
			total -= nums[i];
			if (total == sum) {
				return i;
			}
			sum += nums[i];
		}
		return -1;
	}
};