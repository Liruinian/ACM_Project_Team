#include <bits/stdc++.h>
using namespace std;
int n, p, a[50000010], s[50000010], x, y, z;

int main()
{
    scanf("%d %d", &n, &p);
    for (int i = 1; i <= n; i++)
    {
        scanf("%d", &a[i]);
        s[i] = a[i] - a[i - 1];
    }
    while (p--)
    {
        scanf("%d %d %d", &x, &y, &z);
        s[x] += z;
        s[y + 1] -= z;
    }
    int ans = 1000000000;
    for (int i = 1; i <= n; i++)
    {
        a[i] = a[i - 1] + s[i];
        if (a[i] < ans)
        {
            ans = a[i];
        }
    }
    printf("%d", ans);
    return 0;
}