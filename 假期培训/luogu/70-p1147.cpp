#include <bits/stdc++.h>

int n;
long long c, a[200010], ans = 0;
int main()
{
    scanf("%d %lld", &n, &c);
    for (int i = 1; i <= n; i++)
    {
        scanf("%lld", &a[i]);
        for (int j = 1; j <= n; j++)
        {
            if (a[n] + c == a[j])
            {
                while (a[j])
                {
                    printf("%lld - %lld = c\n", a[n], a[j]);
                    ans++;
                    a[j]--;
                }
            }
        }
    }
    printf("%lld\n", ans);
    return 0;
}