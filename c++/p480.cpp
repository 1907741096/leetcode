class Solution {
public:
	vector<double> medianSlidingWindow(vector<int>& nums, int k) {
		vector<double> res;
		for (int i = 1; i < k - 1; i ++) {
			insertSort(nums, nums[i], 0, i);
		}
		for (int i = k - 1; i < nums.size(); i++) {
			insertSort(nums, nums[i], i - k + 1, i);
			if (k % 2 == 1) {
				res.push_back(nums[(2 * i - k + 1) / 2]);
			}
			else {
				res.push_back((nums[(2 * i - k + 1) / 2] + nums[(2 * i - k + 1) / 2 + 1]) / 2);
			}
		}
		return res;
	}
	
	void insertSort(vector<int>& nums, int k, int start, int end) {
		for (int i = start; i < end; ++ i) {
			if (nums[i] > k) {
				for (int j = end; j > i; -- j) {
					nums[j] = nums[j - 1];
				}
				nums[i] = k;
				break;
			}
		}
		for (int i = 0; i < nums.size(); i ++) {
			cout << nums[i] << " ";
		}
		cout << endl;

	}
};