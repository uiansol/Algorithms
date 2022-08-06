#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);
int numberOfSpaces(string);
bool isHappy(string);
bool hasOnlyOne(string b);

string happyLadybugs(string b) {
    if (isHappy(b)) {
        return "YES";
    }
    
    if (hasOnlyOne(b)) {
        return "NO";
    }
    
    if (numberOfSpaces(b) == 0) {
        return "NO";
    }
    
    return "YES";
}

bool hasOnlyOne(string b) {
    std::map<char,int> letters;
    
    for (char &c: b) {
        letters[c]++;
    }
    
    for (map<char,int>::iterator it=letters.begin(); it != letters.end(); it++) {
        if (it->first != '_' && it->second == 1) {
            return true;
        }
    }
    
    return false;
}

bool isHappy(string b) {
    int n = b.length();
    int spaces = numberOfSpaces(b);
    
    if (spaces == n) {
        return true;
    }
    
    if (n == 1) {
        return false;
    }
    
    if (b[0] != '_' && b[0] != b[1]) {
        return false;
    }
    
    if (b[n - 1] != '_' && b[n - 1] != b[n - 2]) {
        return false;
    }
    
    for (int i = 1; i < n - 1; i++) {
        if (b[i] != '_') {
            if (b[i] != b[i - 1] && b[i] != b[i + 1]) {
                return false;
            }
        }
    }
    
    return true;
}

int numberOfSpaces(string b) { 
    int n = 0;
    
    for (char &c: b) {
        if (c == '_') {
            n++;
        }
    }
    
    return n;
}


int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string g_temp;
    getline(cin, g_temp);

    int g = stoi(ltrim(rtrim(g_temp)));

    for (int g_itr = 0; g_itr < g; g_itr++) {
        string n_temp;
        getline(cin, n_temp);

        stoi(ltrim(rtrim(n_temp)));

        string b;
        getline(cin, b);

        string result = happyLadybugs(b);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}

string ltrim(const string &str) {
    string s(str);

    s.erase(
        s.begin(),
        find_if(s.begin(), s.end(), not1(ptr_fun<int, int>(isspace)))
    );

    return s;
}

string rtrim(const string &str) {
    string s(str);

    s.erase(
        find_if(s.rbegin(), s.rend(), not1(ptr_fun<int, int>(isspace))).base(),
        s.end()
    );

    return s;
}
