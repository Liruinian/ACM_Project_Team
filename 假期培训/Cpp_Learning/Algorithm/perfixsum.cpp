#include <bits/stdc++.h>

long long a[200010];
long long s[200010];

int main()
{
    int n, q;
    scanf("%d%d", &n, &q);
    for (int i = 1; i <= n; i++)
    {
        scanf("%lld", &a[i]);
        s[i] = s[i - 1] + a[i];
    }
    while (q--)
    {
        int l, r;
        scanf("%d%d", &l, &r);
        printf("%lld\n", s[r] - s[l - 1]);
    }
}

// O(n+q)