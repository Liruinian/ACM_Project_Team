#include <bits/stdc++.h>

using namespace std;

stack<int> st;

char aa[6000];

int main()
{
    int _;
    scanf("%d", &_);
    int flag = 1;
    while (_--)
    {
        scanf("%s", aa);
        for (int i = 0; aa[i] != '\0'; i++)
        {
            if (aa[i] == '(')
            {
                st.push(1);
            }
            else if (aa[i] == ')')
            {
                if (st.empty())
                {
                    break;
                }
                else
                {
                    st.pop();
                }
            }
        }

        if (!st.empty())
        {
            flag = 0;
        }
        if (flag == 0)
        {
            printf("NO");
        }
        else
        {
            printf("YES");
        }
        return 0;
    }
}
