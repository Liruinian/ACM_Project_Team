#include <bits/stdc++.h>

char c[1000010];
long long s[1000010];

int main()
{
    scanf("%s", c + 1);
    // c + 1可以使数组从1开始记录
    int n = strlen(c + 1);
    for (int i = 1; i <= n; i++)
    {
        if (c[i] == 'L')
            s[i] = s[i - 1] + 1;
        else
            s[i] = s[i - 1];
    }
    long long ans = 0;
    for (int i = 2; i <= n - 1; i++)
    {
        if (c[i] == 'O')
            ans += s[i - 2] * (s[n] - s[i + 1]);
    }
    printf("%lld\n", ans);
    return 0;
}
