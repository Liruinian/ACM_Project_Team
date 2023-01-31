#include <bits/stdc++.h>

int i, j, sum, m;

int main()
{
    scanf("%d", &m);
    i = 1, j = 2;
    sum = 3;
    while (i <= m / 2)
    {
        if (sum < m)
        {
            j++;
            sum += j;
        }
        else if (sum > m)
        {
            sum -= i;
            i++;
        }
        else if (sum == m)
        {
            printf("%d %d\n", i, j);
            sum -= i;
            i++;
        }
    }
    return 0;
}