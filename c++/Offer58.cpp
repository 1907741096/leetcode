using namespace std;
class Solution {
public:
	string reverseWords(string s) {
		string str = "";
		int left = s.size() - 1, right = s.size() - 1;
		while (left >= 0) {
			if (s[left] == ' ') {
				if (left == right) {
					str += s.substr(left + 1, right - left);
					right = left;
				}
				else {
					left--;
					right--;
				}
			}
			else {
				left--;
			}
		}
		return str;
	}
};