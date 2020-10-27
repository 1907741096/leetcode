class Solution {
public:
	bool isLongPressedName(string name, string typed) {
		int i = 0; 
		int j = 0;
		while (i < name.size() || j < typed.size()) {
			if (name[i] == typed[j])
			{
				if (i < name.size()) i++;
				j++;
			}
			else {
				if (i == 0) {
					return false;
				}
				else {
					if (name[i - 1] == typed[j]) {
						j++;
					}
					else {
						return false;
					}
				}
			}
		}
		return true;
	}
};