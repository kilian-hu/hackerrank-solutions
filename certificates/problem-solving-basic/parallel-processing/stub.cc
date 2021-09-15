#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);



/*
 * Complete the 'minTime' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY files
 *  2. INTEGER numCores
 *  3. INTEGER limit
 */

long minTime(vector<int> files, int numCores, int limit) {

}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string files_count_temp;
    getline(cin, files_count_temp);

    int files_count = stoi(ltrim(rtrim(files_count_temp)));

    vector<int> files(files_count);

    for (int i = 0; i < files_count; i++) {
        string files_item_temp;
        getline(cin, files_item_temp);

        int files_item = stoi(ltrim(rtrim(files_item_temp)));

        files[i] = files_item;
    }

    string numCores_temp;
    getline(cin, numCores_temp);

    int numCores = stoi(ltrim(rtrim(numCores_temp)));

    string limit_temp;
    getline(cin, limit_temp);

    int limit = stoi(ltrim(rtrim(limit_temp)));

    long result = minTime(files, numCores, limit);

    fout << result << "\n";

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
