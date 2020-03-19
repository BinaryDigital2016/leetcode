/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int sumRootToLeaf(TreeNode* root) {
        int ret = 0;
        int sum = 0;
        sumRootToLeaf(root,sum,ret);
        return ret;
    }

    void sumRootToLeaf(TreeNode* root,int sum, int& ret){
        if (!root) {
            return;
        }
        sum = (sum<<1) + root->val;
        if (!root->left && !root->right) {
            ret += sum;
            return;
        }
        sumRootToLeaf(root->left, sum, ret);
        sumRootToLeaf(root->right,sum,ret);
    }
};