#include <stdio.h>

int main()
{
    int n, x, i, b;
    long long ans = 0;
    int a[100009];
    scanf("%d%d", &n, &x);
    for (i = 1; i <= n; i++)
    {
        scanf("%d", &a[i]);
        if (a[i] + a[i - 1] > x)
        {
            b = a[i] + a[i - 1] - x;
            a[i] -= b;
            ans += b;
        }
    }

    printf("%lld\n", ans);
    return 0;
}