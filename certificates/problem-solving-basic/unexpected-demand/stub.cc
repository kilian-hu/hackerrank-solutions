#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);



/*
 * Complete the 'filledOrders' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY order
 *  2. INTEGER k
 */

int filledOrders(vector<int> order, int k) {

}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string order_count_temp;
    getline(cin, order_count_temp);

    int order_count = stoi(ltrim(rtrim(order_count_temp)));

    vector<int> order(order_count);

    for (int i = 0; i < order_count; i++) {
        string order_item_temp;
        getline(cin, order_item_temp);

        int order_item = stoi(ltrim(rtrim(order_item_temp)));

        order[i] = order_item;
    }

    string k_temp;
    getline(cin, k_temp);

    int k = stoi(ltrim(rtrim(k_temp)));

    int result = filledOrders(order, k);

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
