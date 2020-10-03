class Solution {
public:
	int numJewelsInStones(string J, string S) {
		map<char, bool> setMap;
		int num = 0;
		for (int i = 0; i < J.length(); i++) {
			setMap[J[i]] = true;
		}
		
		for (int i = 0; i < S.length(); i++) {
			if (setMap[S[i]] == true) num++;
		}
		return num;
	} 
};