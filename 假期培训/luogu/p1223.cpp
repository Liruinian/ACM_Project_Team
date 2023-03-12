#include <bits/stdc++.h>
#include <algorithm>

using namespace std;
int n;
double avg;
struct table
{
    int num;
    long long time;
} a[1020];

bool sorted(table a, table b)
{
    return a.time < b.time;
}

int main()
{
    scanf("%d", &n);
    for (int i = 1; i <= n; i++)
    {
        scanf("%d", &a[i].time);
        a[i].num = i;
    }
    sort(a + 1, a + 1 + n, sorted);
    for (int i = 1; i <= n; i++)
    {
        printf("%d ", a[i].num);
        avg += a[i].time * (n - i);
    }
    avg /= n;
    printf("\n%.2f", avg);
    return 0;
}