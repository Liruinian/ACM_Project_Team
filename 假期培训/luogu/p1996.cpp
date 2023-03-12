#include <cstdio>
#include <queue>
using namespace std;
int main()
{
    int n, m;
    queue<int> queue;
    scanf("%d%d", &n, &m);
    for (int i = 1; i <= n; i++)
    {
        queue.push(i);
    }
    for (int i = 1; i <= n; i++)
    {
        int f = queue.front();
        queue.pop();
        for (int j = 0; j < m - 1; j++)
        {
            queue.push(f);
            f = queue.front();
            queue.pop();
        }
        printf("%d ", f);
    }
    return 0;
}