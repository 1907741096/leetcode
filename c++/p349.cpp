class Solution {
public:
	vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
		vector<int> res;
		sort(nums1.begin(), nums1.end());
		sort(nums2.begin(), nums2.end());
		// 获取交集
		set_intersection(nums1.begin(), nums1.end(), nums2.begin(), nums2.end(), back_inserter(res));
		vector<int>::iterator iter = unique(res.begin(), res.end());
		res.erase(iter, res.end());
		return res;
	}
};