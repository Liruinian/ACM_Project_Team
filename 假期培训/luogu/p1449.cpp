#include <iostream>
#include <stack>
using namespace std;

int main()
{
    stack<int> st;
    char c;
    int temp, a, b;
    while (c != '@')
    {
        c = getchar();
        switch (c)
        {
        case '+':
            a = st.top();
            st.pop();
            b = st.top();
            st.pop();
            st.push(a + b);
            break;
        case '-':
            a = st.top();
            st.pop();
            b = st.top();
            st.pop();
            st.push(b - a);
            break;
        case '*':
            a = st.top();
            st.pop();
            b = st.top();
            st.pop();
            st.push(a * b);
            break;
        case '/':
            a = st.top();
            st.pop();
            b = st.top();
            st.pop();
            st.push(b / a);
            break;
        case '.':
            st.push(temp);
            temp = 0;
            break;
        default:
            temp = temp * 10 + c - '0';
            break;
        }
    }
    cout << st.top();
    return 0;
}