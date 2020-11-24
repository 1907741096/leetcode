class Solution {
public:
	vector<int> smallestK(vector<int>& arr, int k) {
		return vector<int>(arr.rbegin(), arr.rbegin() + k);
	}

	void BubbleSort(vector<int>& arr, int k) {
		int temp;
		bool flag;
		for (int i = 0; i < k; i ++) {
			flag = true;
			for (int j = arr.size() - 1; j > i; j --) {
				if (arr[j] < arr[j - 1]) {
					temp = arr[j];
					arr[j] = arr[j - 1];
					arr[j - 1] = temp;
					flag = false;
				}
			}
			if (flag) return;
		}
	}

	void SelectSort(vector<int>& arr, int k) {
		int temp, minIndex;
		for (int i = 0; i < k; i++) {
			minIndex = i;
			for (int j = arr.size() - 1; j > i; j--) {
				if (arr[j] > arr[minIndex]) {
					minIndex = j;
				}
			}
			temp = arr[i];
			arr[i] = arr[minIndex];
			arr[minIndex] = temp;
		}
	}

	void InsertSort(vector<int>& arr, int k) {
		int insertNum, i, j;
		for (i = 1; i < arr.size(); i++) {
			insertNum = arr[i];
			for (j = i - 1; j >= 0; j--) {
				if (arr[j] > insertNum) {
					arr[j + 1] = arr[j];
				}
				else {
					break;
				}
			}
			arr[j + 1] = insertNum;
		}
	}

	void ShellSort(vector<int>& arr) {
		int n = arr.size(), insertNum, d, i, j;
		for (d = n/2; d >= 1; d /= 2) {
			for (i = d; i < n; i++) {
				insertNum = arr[i];
				for (j = i - d; j >= 0; j-=d) {
					if (j >= 0 && arr[j] > insertNum) {
						arr[j + d] = arr[j];
					}
					else {
						break;
					}
				}
				arr[j + d] = insertNum;
			}
		}
	}

	void HeapSort1(vector<int>& arr, int k) {
		int length = arr.size(), temp;
		BuildMinHeap(arr, length);
		temp = arr[0];
		arr[0] = arr[length - 1];
		arr[length - 1] = temp;
		for (int i = 1; i < k; i++) {
			AdjustMinHeap(arr, 0, length - i);
			temp = arr[0];
			arr[0] = arr[length - i - 1];
			arr[length - i - 1] = temp;
		}
	}
	void BuildMinHeap(vector<int>& arr, int length) {
		for (int i = (length - 1) / 2; i >= 0; i--) {
			AdjustMinHeap(arr, i, length);
		}
	}
	void AdjustMinHeap(vector<int>& arr, int i, int length) {
		int index, nextIndex, temp = arr[i];
		for (index = i; index * 2 + 1 < length;) {
			if (index * 2 + 2 < length && arr[index * 2 + 1] > arr[index * 2 + 2]) {
				nextIndex = index * 2 + 2;
			}
			else {
				nextIndex = index * 2 + 1;
			}
			if (temp < arr[nextIndex]) {
				break;
			}
			else {
				arr[index] = arr[nextIndex];
				index = nextIndex;
			}
		}
		arr[index] = temp;
	}

	vector<int> HeapSort2(vector<int>& arr, int k) {
		int length = arr.size(), temp;
		if (k == length) {
			return arr;
		}
		BuildMaxHeap(arr, k);
		if (arr[0] > arr[k]) {
			temp = arr[0];
			arr[0] = arr[k];
			arr[k] = temp;
		}
		for (int i = k; i < length; i++) {
			AdjustMaxHeap(arr, 0, k);
			if (arr[0] > arr[i]) {
				temp = arr[0];
				arr[0] = arr[i];
				arr[i] = temp;
			}
		}
		return vector<int>(arr.begin(), arr.begin() + k);
	}
	void BuildMaxHeap(vector<int>& arr, int length) {
		for (int i = (length - 1) / 2; i >= 0; i--) {
			AdjustMaxHeap(arr, i, length);
		}
	}
	void AdjustMaxHeap(vector<int>& arr, int i, int length) {
		int index, nextIndex, temp = arr[i];
		for (index = i; index * 2 + 1 < length;) {
			if (index * 2 + 2 < length && arr[index * 2 + 1] < arr[index * 2 + 2]) {
				nextIndex = index * 2 + 2;
			}
			else {
				nextIndex = index * 2 + 1;
			}
			if (temp > arr[nextIndex]) {
				break;
			}
			else {
				arr[index] = arr[nextIndex];
				index = nextIndex;
			}
		}
		arr[index] = temp;
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
	vector<int> arr2;
	void MergeSort(vector<int>& arr, int left, int right) {
		if (left < right) {
			int mid = left + (right - left) / 2, i, j, k;
			MergeSort(arr, left, mid);
			MergeSort(arr, mid + 1, right);
			for (i = left; i <= right; i++) {
				arr2[i] = arr[i];
			}
			i = left; j = mid + i; k = i;
			while (i <= mid && j <= right) {
				if (arr2[i] <= arr2[j]) {
					arr[k++] = arr[i++];
				}
				else {
					arr[k++] = arr[j++];
				}
			}
			while (i <= mid) arr[k++] = arr[i++];
			while (j <= right) arr[k++] = arr[j++];
		}
		return;
	}
};