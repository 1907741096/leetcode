class Solution {
public:
	int characterReplacement(string s, int k) {
		int max = 0, sum = 0, n = k, l;
		char c;
		for (int i = 0; i < s.size(); i++, n = k, sum = 0) {
			c = s[i];
			l = i;
			for (int j = l; j < s.size(); j++) {
				if (s[j] == c) {
					sum++;
					if (n == k) {
						i = j;
					}
				}
				else {
					if (n--) {
						sum++;
					}
					else {
						break;
					}
				}
			}
			if (n > 0) {
				for (int j = l - 1; j >= 0; j--) {
					if (s[j] == c) {
						sum++;
					}
					else {
						if (n--) {
							sum++;
						}
						else {
							break;
						}
					}
				}
			}
			if (sum > max) {
				max = sum;
			}
		}
		return max;
	}
};