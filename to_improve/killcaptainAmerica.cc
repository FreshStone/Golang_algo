#include<iostream>
#include<stack>
#include<vector>
using namespace std;

class Solution{
 public: 
   int time=0;                             
   int min(int a, int b) {
      return (a<b)?a:b;
   }
                                  
   void findComponent(int u, int disc[], int low[], stack<int>&stk, bool stkItem[], vector<vector<int>> graph) {
      int neighbor;
      //static int time = 0;
      disc[u] = low[u] = ++time;
      stk.push(u);
      stkItem[u] = true;
      for(int v = 0; v < graph[u].size(); v++) {
         neighbor = graph[u][v];
         if (disc[neighbor] == 0) {
           findComponent(neighbor, disc, low, stk, stkItem,graph);
           low[u] = Solution::min(low[u], low[neighbor]);
         } else if(stkItem[neighbor]){
           low[u] = Solution::min(low[u], low[neighbor]);
         }
      }
      int poppedItem = 0;
      if(low[u] == disc[u]) {
         while(stk.top() != u) {
            poppedItem = stk.top();
            //cout << poppedItem << " ";
            low[poppedItem] = low[u];
            stkItem[poppedItem] = false;
            stk.pop();
         }
         poppedItem = stk.top();
         //cout << poppedItem << endl;
         stkItem[poppedItem] = false;
         stk.pop();
      }
   }
                                  
   int captainAmerica(int N, int M, vector<vector<int>> &V) {
      if (N == 1){
          return 1;
      }
      if (V.size() < N-1){
       return 0;
     }
      bool found = 0;
      int i, j, n;
      int disc[N+1], low[N+1];
      bool stkItem[N+1];
      stack<int> stk;
      vector<vector<int>> graph(N+1);

      for (i = 0; i < M; i++){
         graph[V[i][0]].push_back(V[i][1]);
      }
      //printGraph(graph);
      for(i = 1; i< N+1; i++) {    //initialize all elements
         disc[i] = low[i] = -1;
         stkItem[i] = false;
      }
      
      for(i = 1; i< N+1; i++){
         if(disc[i] == -1){
            Solution::findComponent(i, disc, low, stk, stkItem, graph);
         }
       }
       //for (i = 1; i < N; i++){cout << i<<"-"<< low[i] << endl;}
      for (i = 0; i < M; i++){
         if (low[V[i][0]] != low[V[i][1]]){
            stkItem[low[V[i][0]]] = 1; 
         }
       }
       //for (i = 1; i < N; i++){cout << i << "-"<< stkItem[i] << endl;}
       for (i = 1; i < N+1; i++){
         if (stkItem[low[i]] == 0){
            if (found  == 1){
               if (low[i] == n){
                  j += 1;
               }else{
                  return 0;
               }
            }else{
               found = 1;
               j = 1;
               n = low[i];
            }
         }
       }
       return j;
   }
};

int main() {
   int N[] = {5,6,8,7,3,3,3,3,13,13,8};
   int M[] = {5,8,8,8,3,3,2,2,16,15,9};
   int ans[] = {2,6,1,3,3,1,0,1,3,0,1};
   vector<vector<vector<int>>> tests{
    {{1,2},{2,3},{3,4},{4,3},{5,4}},
    {{1,2},{2,3},{3,4},{3,5},{5,6},{6,4},{4,2},{5,1}},
    {{1,2},{2,3},{3,4},{4,5},{5,6},{6,7},{7,4},{5,8}},
    {{1,2},{2,3},{3,4},{4,2},{3,5},{5,6},{6,7},{7,5}},
    {{1,2},{2,3},{3,1}},
    {{1,2},{1,3},{3,1}},
    {{1,2},{1,3}},
    {{1,2},{2,3}},
    {{1,2},{2,3},{3,4},{4,5},{5,3},{4,9},{9,6},{6,7},{7,8},{8,6},{9,10},{10,11},{11,12},{12,10},{11,13},{13,7}},
    {{1,2},{2,3},{3,4},{4,5},{5,3},{4,9},{9,10},{10,11},{11,12},{12,10},{11,13},{9,6},{6,7},{7,8},{8,6}},
    {{1,2},{2,3},{3,4},{3,5},{5,6},{5,7},{7,8},{8,6},{6,4}},
   };
   for (int i = 0; i < tests.size(); i++){
      Solution obj;
      obj.time = 0;
      if (ans[i] == obj.captainAmerica(N[i],M[i],tests[i])){
         cout << "correct answer" << endl;
      }else{
         cout << "wrong answer" << endl;
      }
   }   
}