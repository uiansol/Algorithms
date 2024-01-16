// Accepted  11 ms  11.56 MB

class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        int l1 = nums1.size();
        int l2 = nums2.size();
        map<int, bool> map1;
        map<int, bool> map2;
        vector<int> output;

        for (int i = 0; i < l1; i++) {
            map1[nums1[i]] = true;
        }

        for (int i = 0; i < l2; i++) {
            map2[nums2[i]] = true;
        }

        for (map<int, bool>::iterator it = map1.begin(); it != map1.end(); it++) {
            if (map2[it->first]) {
                output.push_back(it->first);
            }
        }

        return output;
    }
};