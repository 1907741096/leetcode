class Solution {
public:
	void moveZeroes(vector<int>& nums) {
		int temp, numIndex = 0, zeroIndex = 0;
		while (numIndex < nums.size() && zeroIndex < nums.size()) {
			if (nums[zeroIndex] == 0) {
				if (zeroIndex > numIndex) {
					numIndex = zeroIndex;
				}
				while (++numIndex < nums.size()) {
					if (nums[numIndex] != 0) {
						temp = nums[numIndex];
						nums[numIndex] = nums[zeroIndex];
						nums[zeroIndex] = temp;
						zeroIndex++;
						break;
					}
				}
			}
			else {
				zeroIndex++;
			}
		}
	}
};
