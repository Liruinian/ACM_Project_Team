#include <bits/stdc++.h>
using namespace std;
long long n, l, r, a[1000000], b[1000000], ans;

int main()
{
    scanf("%lld", &n);
    for (int i = 1; i <= n; i++)
    {
        scanf("%lld", &a[i]);
        b[i] = a[i] - a[i - 1];
    }
    l = 1;
    r = n;
    while (l < r)
    {
        }
    printf("%lld", ans);
    return 0;
}