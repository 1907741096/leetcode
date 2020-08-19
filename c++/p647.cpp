#include <iostream>
using namespace std;

class Solution {
public:
	int countSubstrings(string s) {
		int num = 0;
		int l = s.length();
		for (int n = 0; n < l; n++) {
			int i = n, j = n;
			while (i >= 0 && j < l && s[i] == s[j]) {
				i--;
				j++;
				num++;
			}
			i = n;
			int k = n + 1;
			while (i >= 0 && k < l && s[i] == s[k]) {
				i--;
				k++;
				num++;
			}
		}
		return num;
	}
};