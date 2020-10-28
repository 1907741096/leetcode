class Solution {
public:
	bool uniqueOccurrences(vector<int>& arr) {
		map<int, int> occur;
		for (const auto& x : arr) {
			occur[x]++;
		}
		set<int> times;
		for (const auto& x : occur) {
			times.insert(x.second);
		}
		return times.size() == occur.size();
	}
};