// Accepted  9 ms  10.72 MB

struct Row {
    int idx;
    int power;
};

bool comp (Row a, Row b) { 
    if (a.power == b.power) {
        return a.idx < b.idx;
    }
    return a.power < b.power;
}

class Solution {
public:
    int value(vector<int> &r) {
        int mid;
        int left = 0;
        int right = r.size();

        while (left < right) {
            mid = left + (right - left) / 2;
            if (r[mid] == 0) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        return left - 1;
    }

    vector<int> kWeakestRows(vector<vector<int>>& mat, int k) {
        vector<Row> rows;

        for (int i = 0; i < mat.size(); i++) {
            Row r;
            r.idx = i;
            r.power = value(mat[i]);

            rows.push_back(r);
        }

        sort(rows.begin(), rows.end(), comp);

        vector<int> output;
        for (int i = 0; i < k; i++) {
            output.push_back(rows[i].idx);
        }

        return output;
    }
};