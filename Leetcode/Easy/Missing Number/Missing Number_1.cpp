// Accepted  3 ms  18.29 MB

class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int sum = accumulate(nums.begin(), nums.end(), 0);
        int expected = (nums.size() * (nums.size() + 1)) / 2;

        return expected - sum;
    }
};