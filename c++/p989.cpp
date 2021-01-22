class Solution {
public:
	vector<int> addToArrayForm(vector<int>& A, int K) {
		int length = A.size();
		A[length - 1] += K;
		while(--length) {
			if (A[length] < 10) {
				break;
			}
			A[length - 1] += A[length] / 10;
			A[length] %= 10;
		}
		while (A[0] >= 10) {
			A.push_back(0);
			for (int i = A.size() - 1; i > 1; i --) {
				A[i] = A[i - 1];
			}
			A[1] = A[0] % 10;
			A[0] = A[0] / 10;
		}
		return A;
	}
};