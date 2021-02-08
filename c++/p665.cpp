class Solution {
public:
	bool checkPossibility(vector<int>& nums) {
		bool flag = true;
		for (int i = 1; i < nums.size() - 1; i++) {
			if (nums[i] < nums[i - 1] && nums[i] > nums[i + 1]) {
				return false;
			}
			else if (nums[i] < nums[i - 1] || nums[i] > nums[i + 1]) {
				if (flag) {
					flag = false;
					if (nums[i] > nums[i - 1] && nums[i] > nums[i + 1]) {
						if (nums[i - 1] <= nums[i + 1]) {
							nums[i] = nums[i + 1];
						}
						else {
							nums[i + 1] = nums[i];
						}
					}
				}
				else {
					return false;
				}
			}
			else {

			}
		}
		return true;
	}
};