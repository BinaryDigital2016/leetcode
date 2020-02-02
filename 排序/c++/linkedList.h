#include<iostream>
using namespace std;

class LNode {
    public:
        LNode():m_iVal(-0xFF),m_pNext(NULL){}
        LNode(int val):m_iVal(val),m_pNext(NULL){}
        ~LNode(){
            if (m_pNext != NULL) {
                delete m_pNext;
                m_pNext = NULL;
            }
        }

    public:
        int m_iVal;
        LNode* m_pNext;
};

class LList {
    public:
        LList():m_iLen(0){
            LNode* head = new LNode();
            this->m_pHead = head;
        }

        ~LList(){
            if (this->m_pHead != NULL) {
                delete this->m_pHead;
                this->m_pHead = NULL;
            }
        }

        void Append(int val){
            LNode *newNode = new LNode(val);

            if (this->m_pHead == NULL){
                LNode* head = new LNode();
                this->m_pHead = head;
                this->m_iLen = 1;

                this->m_pHead->m_pNext = newNode;
            } else {
                LNode* cur = this->m_pHead;
                while (cur->m_pNext != NULL) {
                    cur = cur->m_pNext;
                }
                cur->m_pNext = newNode;
                ++this->m_iLen;
            }
        }

        int Get(int index) {
            if (this->m_pHead==NULL){
                return -0xff;
            }
            LNode* cur = this->m_pHead;
            while (cur->m_pNext && index > 0){
                cur = cur->m_pNext;
                --index;
            }
            if (index == 0){
                return cur->m_iVal;
            }
            return -0xff;
        }

        void Set(int index, int val){
            if (this->m_pHead==NULL){
                return;
            }
            LNode* cur = this->m_pHead;
            while (cur->m_pNext && index > 0){
                cur = cur->m_pNext;
                --index;
            }
            if (index == 0){
                cur->m_iVal = val;
            }
        }

        int Len(){
            return m_iLen;
        }

        void Swap(int i, int j){
            if (i == j){
                return;
            }
            int iVal = Get(i);
            int jVal = Get(j);
            Set(i,jVal);
            Set(j,iVal);
        }

        void Show(){
            if (this->m_pHead==NULL || this->m_pHead->m_pNext==NULL){
                cout << "empty" << endl;
            }
            LNode* cur = this->m_pHead->m_pNext;
            while (cur){
                cout << cur->m_iVal;
                cur = cur->m_pNext;
                if (cur) {
                    cout << "->";
                }
            }
            cout << endl;
        }

        void ShowReverse(){
            if (this->m_pHead==NULL || this->m_pHead->m_pNext==NULL){
                cout << "empty" << endl;
            }
            this->onShowReverse(this->m_pHead->m_pNext);
            cout << endl;
        }

    private:
        void onShowReverse(LNode* n){
            if (n->m_pNext != NULL){
                onShowReverse(n->m_pNext);
                cout << "<-";
            }
            cout << n->m_iVal;
        }

    private:
        int m_iLen;
        LNode* m_pHead;
};
