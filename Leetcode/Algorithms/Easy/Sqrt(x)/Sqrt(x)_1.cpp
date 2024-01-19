// Accepted  19 ms  6.32 MB

class Solution {
public:
    int mySqrt(int x) {
        for (long i = 0; i <= x; i++) {
            if ((i * i) == x) {
                return i;
            }

            if ((i * i) > x) {
                return i - 1;
            }
        }

        return 0;
    }
};