#include <stdio.h>
#include <iostream>
#include "quickSort.h"

//using namespace std;

// 数组快排
int partition(int *arr, int l, int r){
	int pivot = arr[l];
	int i = l, j = r;
	while (i<j){
		while (i<j && arr[j]>=pivot){
			--j;
		}
		arr[i]=arr[j];
		while(i<j&& arr[i]<= pivot){
			++i;
		}
		arr[j]=arr[i];
	}
	arr[i]=pivot;

	return i;
}

void QSort(int *arr, int l, int r){
	if (l < r) {
		int pivot = partition(arr,l,r);
		QSort(arr,l,pivot-1);
		QSort(arr,pivot+1,r);
	}
}


void QuickSort(int* arr, int len){
		QSort(arr, 0, len-1);
}

void PrintArr(int arr[], int len){
	for (int i=0;i<len;++i){
		std::cout<<arr[i];
		if (i != len-1){
			std::cout<< " ";
		} else {
			std::cout<<std::endl;
		}
	}
}

// 链表快排
int partition2(LList *list, int l, int r){
	int pivot = list->Get(l);
	int i = l, j = l+1;
	while (j <= r){
		while (j <= r && list->Get(j)>=pivot){
			++j;
		}
		if (j <= r){
			++i;
			list->Swap(i,j);
			++j;
		}
	}
	list->Swap(i,l);

	return i;
}

void QSort2(LList *list, int l, int r){
	if (l < r) {
		int pivot = partition2(list,l,r);
		QSort2(list,l,pivot-1);
		QSort2(list,pivot+1,r);
	}
}


void QuickSort2(LList* list, int len){
		QSort2(list, 1, len);
}
