#include <iostream>  
#include <stack>
#include <string>

using namespace std;

class Solution {
public:
    bool isValid(string s) {
        stack<char> stk;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] == '{' || s[i] == '[' || s[i] == '(') {
                stk.push(s[i]);
            }
            else {
                if (stk.size() == 0) {
                    return false;
                }
                if (s[i] == '}') {
                    if (stk.top() != '{') {
                        return false;
                    }
                    stk.pop();
                }
                else if (s[i] == ']') {
                    if (stk.top() != '[') {
                        return false;
                    }
                    stk.pop();
                }
                else {
                    if (stk.top() != '(') {
                        return false;
                    }
                    stk.pop();
                }
            }
        }
        if (stk.size() == 0) {
            return true;
        }
        return false;
    }
};