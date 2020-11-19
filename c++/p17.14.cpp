class Solution {
public:
	vector<int> smallestK(vector<int>& arr, int k) {
		quickSort(arr, 0, arr.size() - 1, k);
		return vector<int>(arr.begin(), arr.begin() + k);
	}

	void quickSort(vector<int>& arr, int left, int right, int k) {
		if (left < right) {
			int leftNum = left;
			int rightNum = right;
			int midNum = arr[left];
			while (left < right) {
				while (left < right && arr[right] >= midNum) right--;
				arr[left] = arr[right];
				while (left < right && arr[left] <= midNum) left++;
				arr[right] = arr[left];
			}
			arr[left] = midNum;
			if (left == k) {
				return;
			}
			else if (left < k) {
				quickSort(arr, left + 1, rightNum, k);
			}
			else {
				quickSort(arr, leftNum, left - 1, k);
			}
		}
		return;
	}
};