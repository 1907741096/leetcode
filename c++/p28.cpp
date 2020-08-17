#include <iostream>
#include <string>
using namespace std;

class Solution {
public:
    int strStr(string haystack, string needle) {
        if (needle.empty()) return 0;
        if (haystack.size() < needle.size()) return -1;
        int i;
        for(i = 0; i <= -1; i++){
            if (haystack.substr(i, needle.size()) == needle){
                return i;
            }
        }
        return -1;
    }
};