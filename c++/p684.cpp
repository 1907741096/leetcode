class Solution {
public:
	vector<int> findRedundantConnection(vector<vector<int>>& edges) {
		int length = edges.size();
		vector<int> parent(length + 1);
		for (int i = 1; i <= length; i ++) {
			parent[i] = i;
		}
		for (int i = 0; i < length; i++) {
			if (Find(parent, edges[i][0]) != Find(parent, edges[i][1])) {
				Union(parent, edges[i][0], edges[i][1]);
			}
			else {
				return edges[i];
			}
		}
		return vector<int>{};
	}

	int Find(vector<int>& parent, int index) {
		if (parent[index] != index) {
			parent[index] = Find(parent, parent[index]);
		}
		return parent[index];
	}

	void Union(vector<int>& parent, int index1, int index2) {
		parent[Find(parent, index1)] = Find(parent, index2);
	}
};