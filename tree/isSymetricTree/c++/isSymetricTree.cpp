#include <queue>
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
// class Solution {
// public:
//     bool isSymmetric(TreeNode* root) {
//         if (!root) return true;

//         return isSymmetric(root->left,root->right);
//     }

//     bool isSymmetric(TreeNode* left, TreeNode* right) {
//         if (!left && !right) return true;

//         if (!left || !right) return false;

//         return (left->val == right->val) && isSymmetric(left->left, right->right) && isSymmetric(left->right, right->left);
//     }
// };

class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        std::queue<TreeNode*> q;
        q.push(root);
        q.push(root);
        while (!q.empty()) {
            TreeNode* l = q.front();
            q.pop();
            TreeNode* r = q.front();
            q.pop();

            if (!l && !r) continue;

            if (!l || !r) return false;

            if (l->val != r->val) return false;

            q.push(l->left);
            q.push(r->right);
            q.push(l->right);
            q.push(r->left);
        }
        return true;
    }
};