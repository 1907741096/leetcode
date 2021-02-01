#include <utility>
using namespace std;

class Solution {
public:
	vector<int> fairCandySwap(vector<int>& A, vector<int>& B) {
		int sumA = accumulate(A.begin(), A.end(), 0), sumB = accumulate(B.begin(), B.end(), 0);
		map<int, int> mapA;
		for (int i = 0; i < A.size(); i ++) {
			sumA += A[i];
			mapA.insert(make_pair(A[i], i));
		}
		for (int &i : B) {
			sumB += i;
		}
		int x;
		for (int& i : B) {
			if ((sumA - sumB + 2 * i) % 2 == 0) {
				x = (sumA - sumB + 2 * i) / 2;
				if (mapA.find(x) != mapA.end()) {
					return vector<int>{x, i};
				}
			}
		}
		return vector<int>{};
	}
};