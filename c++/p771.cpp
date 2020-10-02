class Solution {
public:
	int numJewelsInStones(string J, string S) {
		map<char, bool> setMap;
		int num = 0;
		for (int i = 0; i < J.sizeof(); i++) {
			setMap[J[i]] = true;
		}
		
		for (int i = 0; i < S.sizeof(); i++) {
			if (setMap[S[i]] == true) num++;
		}
		return num;
	}
};