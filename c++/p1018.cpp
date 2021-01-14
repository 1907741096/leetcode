class Solution {
public:
	vector<bool> prefixesDivBy5(vector<int>& A) {
		vector<bool> res;
		for (int i = 0; i < A.size(); i++) {
			res.push_back(((prefix << 1) + A[i]) % 5 == 0);
		}
		return res;
	}
};