#include <bits/stdc++.h>

using namespace std;

int marsExploration(string s) {
    int n = s.length();
    int result = 0;
    
    for (int i = 0; i < n; i++) {
        if ((i - 1) % 3 == 0) {
            if (s[i] != 'O') {
                result++;
            }
        } else {
            if (s[i] != 'S') {
                result++;
            }
        }
    }
    
    return result;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    int result = marsExploration(s);

    fout << result << "\n";

    fout.close();

    return 0;
}
