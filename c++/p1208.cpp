class Solution {
public:
	int equalSubstring(string s, string t, int maxCost) {
		vector<int> arr(s.size());
		for (int i = 0; i < s.size(); i ++) {
			arr[i] = abs(s[i] - t[i]);
		}
		int maxNum = 0, cost = maxCost;
		for (int start = 0, end = 0; end < arr.size();) {
			if (maxCost < arr[end]) {
				end++;
				start = end;
				cost = maxCost;
			}
			else {
				if (cost >= arr[end]) {
					cost -= arr[end++];
					if (maxNum < end - start) {
						maxNum = end - start;
					}
				}
				else {
					cost += arr[start++];
				}
			}
		}
		return maxNum;
	}
};