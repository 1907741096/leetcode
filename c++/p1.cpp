class Solution {
public:
	vector<int> twoSum(vector<int>& nums, int target) {
		map<int, int> hashMap;
		vector<int> res;
		for (int i = 0; i < nums.size(); i++) {
			int num = target - nums[i];
			auto it = hashMap.find(target - nums[i]);
			if (it != hashMap.end()) {
				return { it->second, i };
			}
			hashMap.insert(make_pair(nums[i], i));
		}
		return res;
	}
};