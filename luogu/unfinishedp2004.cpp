#include <bits/stdc++.h>

long long a[1010][1010], s[1010][1010];

int sum(int x1, int y1, int x2, int y2)
{
    printf("%d %d %d %d \n", x1, y1, x2, y2);
    return s[x2][y2] - s[x1 - 1][y2] - s[x2 - 1][y1] + s[x1 - 1][x2 - 1];
}
int main()
{
    int n, m, c;
    scanf("%d %d %d", &n, &m, &c);
    for (int i = 1; i <= n; i++)
    {
        for (int j = 1; j <= m; j++)
        {
            scanf("%lld", &a[i][j]);
            s[i][j] = s[i - 1][j] + s[i][j - 1] - s[i - 1][j - 1] + a[i][j];
        }
    }
    long long biggest = 0;
    int b_x, b_y;
    printf("\n");
    for (int i = c; i <= n; i++)
    {
        for (int j = c; j <= m; j++)
        {
            long long sum_l = sum(i - c, j - c, i, j);
            printf("%d %d %d \n", i, j, sum_l);
            if (sum_l > biggest)
            {
                biggest = sum_l;
                b_x = i - c + 1;
                b_y = j - c + 1;
            }
        }
    }
    printf("%d %d", b_x, b_y);
}