#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);



/*
 * Complete the 'getMinCost' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY crew_id
 *  2. INTEGER_ARRAY job_id
 */

long getMinCost(vector<int> crew_id, vector<int> job_id) {

}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string crew_id_count_temp;
    getline(cin, crew_id_count_temp);

    int crew_id_count = stoi(ltrim(rtrim(crew_id_count_temp)));

    vector<int> crew_id(crew_id_count);

    for (int i = 0; i < crew_id_count; i++) {
        string crew_id_item_temp;
        getline(cin, crew_id_item_temp);

        int crew_id_item = stoi(ltrim(rtrim(crew_id_item_temp)));

        crew_id[i] = crew_id_item;
    }

    string job_id_count_temp;
    getline(cin, job_id_count_temp);

    int job_id_count = stoi(ltrim(rtrim(job_id_count_temp)));

    vector<int> job_id(job_id_count);

    for (int i = 0; i < job_id_count; i++) {
        string job_id_item_temp;
        getline(cin, job_id_item_temp);

        int job_id_item = stoi(ltrim(rtrim(job_id_item_temp)));

        job_id[i] = job_id_item;
    }

    long result = getMinCost(crew_id, job_id);

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
