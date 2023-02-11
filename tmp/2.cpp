
#include "stdafx.h"
#include<iostream>
#include<cstdio>
#include<algorithm>
#include<cmath>
using namespace std;
struct Node
{
    double l;//最左可被覆盖的坐标
    double r;//最右可被覆盖的坐标
};
Node a[1010];
int i,j,k,n,m;
double d,x,y;
int ans;//结果
const double ZERO=1e-8;//控制精度
 
 
bool cmp(Node a,Node b)//最左可被覆盖的坐标排序，从左到右排序
{
    if (a.l<b.l) return 1;
    return 0;
}
 
void Greedy()
{
    double now=a[0].r;
    ans++;
    for (int i=1;i<n;i++)
    {
        if (a[i].l>now+ZERO)//下个点的最左被覆盖的坐标大于当前最右可被覆盖坐标
        {
            ans++;
            now=a[i].r;
        }else if (a[i].r<now+ZERO)//下个点的最左被覆盖的坐标小于当前最右可被覆盖坐标
        {
            now=a[i].r;
        }
 
    }
}
 
int main()
{
    printf("输入覆盖半径d:\n");
    cin>>d;
    printf("输入船的数量n:\n");
    cin>>n;
    printf("输入n条船的坐标:\n");
    for(int i=0;i<n;i++)
    {
        cin>>x>>y;
        double len=sqrt((double)(d*d-y*y));//勾股定理
        a[i].l=x-len;//计算最左可被覆盖的坐标
        a[i].r=x+len;//计算最右可被覆盖的坐标
    }
    sort(a,a+n,cmp);
    ans=0;
    Greedy();
    printf("在x轴上至少要安放%d个点才可以将所有点覆盖\n",ans);
    return 0;
}