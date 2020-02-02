#include <iostream>
#include "linkedList.h"

using namespace std;

// 数组快排
int partition(int *arr, int l, int r);
void QSort(int *arr, int l, int r);
void QuickSort(int* arr, int len);
void PrintArr(int arr[], int len);

// 列表快排
int partition2(LList *list, int l, int r);
void QSort2(LList *list, int l, int r);
void QuickSort2(LList* list, int len);
