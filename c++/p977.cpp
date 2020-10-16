class Solution {
public:
	vector<int> sortedSquares(vector<int>& A) {
		int left = 0;
		int right = A.size() - 1;
		vector<int> res(right + 1);
		int index = right;
		while (left <= right) {
			if (A[left] * A[left] > A[right] * A[right]) {
				res[index--] = A[left] * A[left];
				left++;
			}
			else {
				res[index--] = A[right] * A[right];
				right--;
			}
		}
		return res;
	}
};