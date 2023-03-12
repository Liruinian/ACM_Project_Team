#include <stdio.h>
int main()
{
    int a[100], i, j;
    for (i = 0; i < 100; i++)
    {
        scanf("%d", &a[i]);
        if (a[i] == 0)
        {
            break;
        }
    }

    for (j = i - 1; j >= 0; j--)
    {
        printf("%d ", a[j]);
    }
    return 0;
}