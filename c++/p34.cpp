class Solution {
public:
	vector<int> searchRange(vector<int>& nums, int target) {
		vector<int> res;
		int mid, left = 0, right = nums.size() - 1;
		while (left < right) {
			mid = left + (right - left) / 2;
			if (nums[mid] == target) {
				if (mid == 0 || nums[mid - 1] < target) {
					break;
				}
				else {
					right = mid - 1;
				}
			}
			else if (nums[mid] < target) {
				left = mid + 1;
			}
			else {
				right = mid - 1;
			}
		}
		res.push_back(mid);
		left = 0;
		right = nums.size() - 1;
		while (left < right) {
			mid = left + (right - left) / 2;
			if (nums[mid] == target) {
				if (mid == 0 || nums[mid - 1] < target) {
					break;
				}
				else {
					right = mid - 1;
				}
			}
			else if (nums[mid] < target) {
				left = mid + 1;
			}
			else {
				right = mid - 1;
			}
		}
		res.push_back(mid);
		return res;
	}
};