#include <iostream>

using namespace std;

int main()
{
    string linea;
    cin >> linea;

    cout << linea[0];

    for (int i = 1; i < linea.size(); ++i) {
        if (linea[i] == '-' and i != linea.size()) {
            cout << linea[i+1];
        }
    }
    cout << "\n";

    return 0;
}
