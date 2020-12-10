class Solution {
public:
	bool lemonadeChange(vector<int>& bills) {
		int n5 = 0, n10 = 0, n20 = 0, res;
		for (int n : bills) {
			if (n == 5) {
				n5++;
			}
			else if (n == 10) {
				n10++;
			}
			else
			{
				n20++;
			}
			n -= 5;
			while (n > 0)
			{
				if (n - 20 >= 0 && n20 > 0) {
					n -= 20;
					n20--;
				}
				else if (n - 10 >= 0 && n10 > 0) {
					n -= 10;
					n10--;
				}
				else if (n - 5 >= 0 && n5 > 0) {
					n -= 5;
					n5--;
				}
				else {
					return false;
				}
			}
		}
		return true;
	}
};