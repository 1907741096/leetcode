class Solution {
public:
	void nextPermutation(vector<int>& nums) {
		int temp, left, right;
		for (int i = nums.size() - 1; i > 0; i --) {
			if (nums[i] > nums[i - 1]) {
				for (int j = nums.size() - 1; j >= i; j--) {
					if (nums[j] > nums[i - 1]) {
						temp = nums[j];
						nums[j] = nums[i - 1];
						nums[i - 1] = temp;
						break;
					}
				}
				left = i, right = nums.size() - 1;
				while (left < right) {
					temp = nums[left];
					nums[left] = nums[right];
					nums[right] = temp;
					left++;
					right--;
				}
				break;
			}
			if (i == 1) {
				left = 0, right = nums.size() - 1;
				while (left < right) {
					temp = nums[left];
					nums[left] = nums[right];
					nums[right] = temp;
					left++;
					right--;
				}
			}
		}
	}
};