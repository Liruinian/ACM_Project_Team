#include <bits/stdc++.h>

long long n, a[200005], b[200005], ans = 0;
int main()
{
    scanf("%lld", &n);
    for (int i = 1; i <= n; i++)
    {
        scanf("%lld", &a[i]);
        if (i == 1)
        {
            b[i] = a[i];
        }
        else
        {
            b[i] = b[i - 1] + a[i];
        }
    }
    for (int i = 1; i < n; i++)
    {
        ans += a[i] * (b[n] - b[i]);
    }
    printf("%lld\n", ans);
    return 0;
}