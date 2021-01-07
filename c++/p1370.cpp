class Solution {
public:
	string sortString(string s) {
		vector<int> arr(26);
		for (char& c : s) {
			arr[c - 'a'] ++;
		}
		string res;
		bool isZero, flag = true;
		while (true) {
			isZero = true;
			if (flag) {
				for (int i = 0; i < 26; i ++) {
					if (arr[i] > 0) {
						res += ('a' + i);
						arr[i]--;
						isZero = false;
					}
				}
			}
			else {
				for (int i = 25; i >= 0; i--) {
					if (arr[i] > 0) {
						res += ('a' + i);
						arr[i]--;
						isZero = false;
					}
				}
			}
			if (isZero) break;
			flag = !flag;
		}
		return res;
	}
};