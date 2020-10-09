class Solution {
public:
	vector<int> getLeastNumbers(vector<int>& arr, int k) {
		if (k == 0) return {};
		vector<int> heapList;
		for (int i = 0; i < k; i++) {
			heapList.push_back(arr[i]);
		}
		initHeap(heapList);
		for (int i = k; i < arr.size(); i++) {
			if (arr[i] < heapList[0]) {
				heapList[0] = arr[i];
				setHeap(heapList);
			}
		}
		return heapList;
	}

	void initHeap(vector<int>& heapList) {
		int swap;
		for (int i = heapList.size() / 2 - 1; i >= 0; i--) {
			if (2 * i + 2 < heapList.size() && heapList[2 * i + 1] < heapList[2 * i + 2]) {
				swap = heapList[2 * i + 1];
				heapList[2 * i + 1] = heapList[2 * i + 2];
				heapList[2 * i + 2] = swap;
			}
			if (2 * i + 1 < heapList.size() && heapList[i] < heapList[2 * i + 1]) {
				swap = heapList[i];
				heapList[i] = heapList[2 * i + 1];
				heapList[2 * i + 1] = swap;
			}
		}
	}

	void setHeap(vector<int>& heapList) {
		int swap;
		for (int i = 0; i < heapList.size() / 2;) {
			if (2 * i + 2 < heapList.size() && heapList[i] < heapList[2 * i + 2]) {
				if (heapList[2 * i + 2] < heapList[2 * i + 1]) {
					swap = heapList[i];
					heapList[i] = heapList[2 * i + 1];
					heapList[2 * i + 1] = swap;
					i = 2 * i + 1;
				}
				else {
					swap = heapList[i];
					heapList[i] = heapList[2 * i + 2];
					heapList[2 * i + 2] = swap;
					i = 2 * i + 2;
				}
			}
			else if (2 * i + 1 < heapList.size() && heapList[i] < heapList[2 * i + 1]) {
				swap = heapList[i];
				heapList[i] = heapList[2 * i + 1];
				heapList[2 * i + 1] = swap;
				i = 2 * i + 1;
			}
			else {
				break;
			}
		}
	}
};