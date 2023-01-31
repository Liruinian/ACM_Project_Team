#include <stdio.h>
int main()
{
    char a;
    int c = 0;
    while (scanf("%c", &a) != EOF)
    {
        if (a == '(')
        {
            c++;
        }
        else if (a == ')')
        {
            if (c <= 0)
            {
                printf("NO");
                return 0;
            }
            else
                c--;
        }
    }
    if (c == 0)
        printf("YES");
    else
        printf("NO");
    return 0;
}