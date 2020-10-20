#include <stack>
class Solution {
public:

	bool backspaceCompare(string S, string T) {
		std::stack<char> s;
		std::stack<char> t;
		for (int i = 0; i < S.size(); i++) {
			if (S[i] == '#') {
				if (s.size() != 0) {
					s.pop();
				}
			}
			else {
				s.push(S[i]);
			}
		}
		for (int i = 0; i < T.size(); i++) {
			if (T[i] == '#') {
				if (t.size() != 0) {
					t.pop();
				}
			}
			else {
				t.push(T[i]);
			}
		}
		if (s.size() == t.size()) {
			int n = s.size();
			while(n--) {
				if (s.top() != t.top()) {
					return false;
				}
				s.pop();
				t.pop();
			}
			return true;
		}
		return false;
	}
};