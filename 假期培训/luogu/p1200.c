#include <stdio.h>

int main()
{
    char h[10], d[10];
    int i;
    long long a = 1, b = 1;
    scanf("%s%s", &h, &d);
    for (i = 0; i < 7; i++)
    {
        if (h[i] - 'A' + 1 < 30 && h[i] - 'A' + 1 > 0)
        {
            a *= h[i] - 'A' + 1;
        }
    }

    for (i = 0; i < 7; i++)
    {

        if (d[i] - 'A' + 1 < 30 && d[i] - 'A' + 1 > 0)
        {
            b *= d[i] - 'A' + 1;
        }
    }
    if (b % 47 == a % 47)
    {
        printf("GO");
    }
    else
    {
        printf("STAY");
    }

    return 0;
}