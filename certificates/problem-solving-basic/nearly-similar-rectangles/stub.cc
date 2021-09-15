#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);
vector<string> split(const string &);



/*
 * Complete the 'nearlySimilarRectangles' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts 2D_LONG_INTEGER_ARRAY sides as parameter.
 */

long nearlySimilarRectangles(vector<vector<long>> sides) {

}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string sides_rows_temp;
    getline(cin, sides_rows_temp);

    int sides_rows = stoi(ltrim(rtrim(sides_rows_temp)));

    string sides_columns_temp;
    getline(cin, sides_columns_temp);

    int sides_columns = stoi(ltrim(rtrim(sides_columns_temp)));

    vector<vector<long>> sides(sides_rows);

    for (int i = 0; i < sides_rows; i++) {
        sides[i].resize(sides_columns);

        string sides_row_temp_temp;
        getline(cin, sides_row_temp_temp);

        vector<string> sides_row_temp = split(rtrim(sides_row_temp_temp));

        for (int j = 0; j < sides_columns; j++) {
            long sides_row_item = stol(sides_row_temp[j]);

            sides[i][j] = sides_row_item;
        }
    }

    long result = nearlySimilarRectangles(sides);

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

vector<string> split(const string &str) {
    vector<string> tokens;

    string::size_type start = 0;
    string::size_type end = 0;

    while ((end = str.find(" ", start)) != string::npos) {
        tokens.push_back(str.substr(start, end - start));

        start = end + 1;
    }

    tokens.push_back(str.substr(start));

    return tokens;
}
