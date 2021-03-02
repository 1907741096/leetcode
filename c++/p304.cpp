class NumMatrix {
public:
    vector<vector<int>> arr;

    NumMatrix(vector<vector<int>>& matrix) {
        if (matrix.size() == 0)
            return;
        arr = vector<vector<int>>(matrix.size() + 1);
        for (int i = 0; i <= matrix.size(); i++) {
            arr[i] = vector<int>(matrix[0].size() + 1, 0);
        }
        for (int i = matrix.size() - 1; i >= 0; i--) {
            for (int j = matrix[i].size() - 1; j >= 0; j--) {
                arr[i][j] = matrix[i][j] + arr[i + 1][j] + arr[i][j + 1] - arr[i + 1][j + 1];
            }
        }
    }

    int sumRegion(int row1, int col1, int row2, int col2) {
        return arr[row1][col1] - arr[row1][col2 + 1] - arr[row2 + 1][col1] + arr[row2 + 1][col2 + 1];
    }
};

/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix* obj = new NumMatrix(matrix);
 * int param_1 = obj->sumRegion(row1,col1,row2,col2);
 */