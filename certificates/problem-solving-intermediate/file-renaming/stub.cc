#include <bits/stdc++.h>

using namespace std;



/*
 * Complete the 'renameFile' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING newName
 *  2. STRING oldName
 */

int renameFile(string newName, string oldName) {

}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string newName;
    getline(cin, newName);

    string oldName;
    getline(cin, oldName);

    int result = renameFile(newName, oldName);

    fout << result << "\n";

    fout.close();

    return 0;
}
