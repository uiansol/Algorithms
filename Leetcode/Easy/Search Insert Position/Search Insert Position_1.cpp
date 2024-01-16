// Accepted  8 ms  9.95 MB

class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size() - 1;
        int mid;

        while (left < right) {
            mid = left + (right - left) / 2;

            if (nums[mid] == target) {
                return mid;
            }

            if (nums[mid] > target) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        if (target > nums[left]) {
            return left + 1;
        }
        return left;
    }
};